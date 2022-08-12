package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notional-labs/pixel/server"

	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func GetPixelHandler(c *gin.Context) {
	node, err := client.NewClientFromNode("https://rpc.uni.junonetwork.io:443")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot connect to chain",
		})
	}

	clientCtx := client.Context{}
	clientCtx = clientCtx.WithClient(node).WithNodeURI("https://rpc.uni.junonetwork.io:443")

	queryClient := wasmTypes.NewQueryClient(clientCtx)

	data, err := server.GetData(queryClient, 11, 11)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot fetch data",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
