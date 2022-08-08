package socket

import (
	"context"
	"fmt"
	"time"

	rpchttp "github.com/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
)

func errorHandler() {
	fmt.Printf("error")
}

func GetWebsocketCLient() {
	client, err := rpchttp.New("tcp://0.0.0.0:26657", "/websocket") // add juno testnet rpc

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
	txs, err := client.WaitForOneEvent(ctx, "test-client", query)
	if err != nil {
		errorHandler()
	}

	go func() {
		for e := range txs {
			fmt.Println("got ", e.Data.(types.EventDataTx))
		}
	}()
}
