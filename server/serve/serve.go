package serve

import (
	"context"
	"fmt"
	"time"

	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controller "github.com/notional-labs/pixel/server/serve/controller"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func errorHandler() {
	fmt.Printf("error")
}

func setupRoute(router *gin.Engine) {
	// routes
	router.GET("/api/pixels", controller.GetPixelHandler)
}

func ListenAndServe(queryClient wasmTypes.QueryClient, port string) {
	// websocket
	client, err := rpchttp.New("http://95.217.121.243:2081", "/websocket")

	if err != nil {
		fmt.Println(err)
		errorHandler()
	}
	err = client.Start()
	if err != nil {
		fmt.Println(err)
		errorHandler()
	}
	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	query := "tm.event = 'NewBlock'"
	txs, envErr := client.Subscribe(ctx, "client", query)
	if envErr != nil {
		errorHandler()
	}
	controller.GetNewBlockHandler()

	go func() {
		for range txs {
			controller.GetNewBlockHandler()
		}
	}()

	router := gin.Default()

	router.Use(cors.Default())

	// recover from panic, return 500 err instead
	router.Use(gin.Recovery())

	// setup routes
	setupRoute(router)

	//server listen on port
	router.Run(":" + port)
}
