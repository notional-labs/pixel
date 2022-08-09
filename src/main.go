package main

import (
	"fmt"
	"log"
	"net/http"

	"context"
	"time"

	controller "github.com/notionals-lab/pixel/src/controller"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
)

func errorHandler() {
	fmt.Printf("error")
}

func setupRoute() {
	// routes

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/api/pixels", controller.GetPixelHandler)
}

func main() {

	// websocket
	client, err := rpchttp.New("http://95.217.121.243:2071", "/websocket")

	if err != nil {
		errorHandler()
	}

	err = client.Start()
	if err != nil {
		errorHandler()
	}
	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	query := "tm.event = 'NewBlock'"
	txs, err := client.Subscribe(ctx, "test-client", query)
	if err != nil {
		errorHandler()
	}

	go func() {
		for e := range txs {
			fmt.Println("got ", e.Data.(types.EventDataNewBlock))
		}
	}()

	// setup routes
	setupRoute()

	//server listen on port 8080
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
