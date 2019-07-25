package deribit

import (
	"fmt"
	"strings"
	"time"

	"github.com/adampointer/go-deribit/v3/models"

	"github.com/gorilla/websocket"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RPCCore actually sends and receives messages
type RPCCore struct {
	calls        *callManager
	subs         *subManager
	connMgr      *connManager
	onDisconnect func(*RPCCore)
	errors       chan error
	stop         chan bool
	auth         *models.PublicAuthResponse
}

// Submit satisfies the runtime.ClientTransport interface
func (r *RPCCore) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	// Strip leading slashes
	method := strings.TrimPrefix(operation.PathPattern, "/")
	req := NewRPCRequest(method)
	if err := operation.Params.WriteToRequest(req, strfmt.Default); err != nil {
		return nil, err
	}
	req.AddAuth(r.auth)
	res, err := r.rpcRequest(req)
	if err != nil {
		return nil, err
	}
	return operation.Reader.ReadResponse(res, runtime.JSONConsumer())
}

func (r *RPCCore) close() error {
	r.connMgr.close()
	if err := r.connMgr.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return r.connMgr.conn.Close()
}

func (r *RPCCore) rpcRequest(req *RPCRequest) (*RPCResponse, error) {
	call := NewRPCCall(req)
	r.calls.addCall(call)
	// Send
	if err := r.connMgr.conn.WriteJSON(&req); err != nil {
		r.calls.deleteCall(call)
		return nil, err
	}

	// Wait for response or timeout
	select {
	case <-call.Done:
	case <-time.After(3 * time.Second):
		call.Error = fmt.Errorf("request %d timed out", call.Req.ID)
	}
	if call.Error != nil {
		return nil, call.Error
	}
	if call.Res.Error != nil {
		return nil, fmt.Errorf("request failed with code (%d): %s", call.Res.Error.Code, call.Res.Error.Message)
	}
	return call.Res, nil
}

// read takes messages off the websocket and deals with them accordingly
func (r *RPCCore) read() {
	var resErr error
Loop:
	for {
		select {
		case <-r.stop:
			break Loop
		default:
			var raw composite
			if err := r.connMgr.conn.ReadJSON(&raw); err != nil {
				// fix for `use of closed network connection`
				if r.connMgr.closed() {
					break Loop
				}
				r.onDisconnect(r)
				break Loop
			}

			if raw.ID != 0 || raw.Error != nil {
				resErr = r.handleResponses(raw.toResponse())
			} else if raw.Method == "subscription" {
				r.handleSubscriptions(raw.toNotification())
			}
		}
	}
	if resErr != nil {
		r.calls.closeAllWithError(resErr)
	}
}

func (r *RPCCore) handleResponses(res *RPCResponse) error {
	call := r.calls.getCall(res.ID)

	if strings.Contains(call.Req.Method, "subscribe") {
		if len(res.Result) <= 2 && res.Error == nil {
			res.Error = &RPCError{Code: 10001, Message: "empty result"}
		}
	}

	if res.Error != nil && res.Error.Code != 0 {
		resErr := fmt.Errorf("request failed with code (%d): %s", res.Error.Code, res.Error.Message)
		call.CloseError(resErr)
	} else {
		if call == nil {
			return fmt.Errorf("no pending request found for response ID %d", res.ID)
		}
		call.Res = res
		call.CloseOK()
		r.calls.deleteCall(call)
	}
	return nil
}

func (r *RPCCore) handleSubscriptions(res *RPCNotification) {
	sub := r.subs.getSubscription(res.Params.Channel)
	if sub == nil {
		// Send error to main error channel
		r.errors <- fmt.Errorf("no subscription found for %s", res.Params.Channel)
		return
	}
	// Send the notification to the right channel
	sub.Data <- res
}
