package serve

import (
	"context"
	"fmt"
	"net/http"
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
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	router.GET("/api/pixels/get-chunk", controller.GetPixelHandler)
}

func ListenAndServe(queryClient wasmTypes.QueryClient) {
	// websocket
	client, err := rpchttp.New("tcp://0.0.0.0:36657", "/websocket")

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
	_, envErr := client.Subscribe(ctx, "test-client", query)
	if envErr != nil {
		errorHandler()
	}
	// todo add save new board state func
	go func() {
		fmt.Print("new block")
		// controller.GetNewBlockHandler()
	}()

	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	// recover from panic, return 500 err instead
	router.Use(gin.Recovery())

	// setup routes
	setupRoute(router)

	//server listen on port 8080
	router.Run(":3000")
}
