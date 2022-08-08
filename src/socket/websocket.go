package socket

import (
	"context"
	"fmt"
	"time"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
)

func errorHandler() {
	fmt.Printf("error")
}

func GetWebsocketClient() {
	client, err := rpchttp.New("http://95.217.121.243:2071") // add juno testnet rpc

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
			fmt.Println("got ", e.Data.(types.EventDataTx))
		}
	}()
}
