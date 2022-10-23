package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/notional-labs/pixel/global"
	"github.com/notional-labs/pixel/server"

	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func GetPixelHandler(c *gin.Context) {
	board := global.GetBoard()
	ans := server.ParsePixelArray(board)
	c.JSON(http.StatusOK, gin.H{
		"data": ans,
	})
}

func GetNewBlockHandler() {
	node, err := client.NewClientFromNode("http://95.217.121.243:2081")

	if err != nil {
		return
	}

	clientCtx := client.Context{}
	clientCtx = clientCtx.WithClient(node).WithNodeURI("http://95.217.121.243:2081")

	queryClient := wasmTypes.NewQueryClient(clientCtx)

	data, err := server.GetData(queryClient, 11, 11)
	global.SetBoard(data)
}
