package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	controller "github.com/notional-labs/pixel/src/controller"
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
	_, envErr := client.Subscribe(ctx, "test-client", query)
	if envErr != nil {
		errorHandler()
	}

	// todo add save new board state func
	go func() {
		// queryClient.AsyncGetChuckData()
	}()

	router := gin.Default()

	// setup routes
	setupRoute(router)

	//server listen on port 8080
	router.Run()
}
