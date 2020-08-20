package main

import (
	"fmt"
	"log"

	"github.com/Lexus123/go-deribit"
	"github.com/Lexus123/go-deribit/client/public"
)

func main() {
	errs := make(chan error)
	stop := make(chan bool)
	e, err := deribit.NewExchange(true, errs, stop)

	if err != nil {
		log.Fatalf("Error creating connection: %s", err)
	}
	if err := e.Connect(); err != nil {
		log.Fatalf("Error connecting to exchange: %s", err)
	}
	go func() {
		err := <-errs
		stop <- true
		log.Fatalf("RPC error: %s", err)
	}()
	client := e.Client()

	// Hit the test RPC endpoint
	res, err := client.Public.GetPublicTest(&public.GetPublicTestParams{})
	if err != nil {
		log.Fatalf("Error testing connection: %s", err)
	}
	log.Printf("Connected to Deribit API v%s", *res.Payload.Result.Version)

	// Sub to tradekind
	trade, err := e.SubscribeTradesKind("future", "BTC", "raw")
	if err != nil {
		log.Fatalf("Error subscribing to the book: %s", err)
	}
	for b := range trade {
		fmt.Printf("Amount: %f, Instument: %s\n", *b.Amount, b.InstrumentName)
	}

	e.Close()
}

func strPointer(str string) *string {
	return &str
}
