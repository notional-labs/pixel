package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/notional-labs/pixel/global"
	"github.com/notional-labs/pixel/server"

	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func GetPixelHandler(c *gin.Context) {
	// print board
	var buffer bytes.Buffer
	var b []byte
	var err error
	board := global.GetBoard()

	for _, item := range board {
		b, err = json.Marshal(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		buffer.WriteString(string(b) + ",")
	}

	s := strings.TrimSpace(buffer.String())
	// trim last comma

	c.JSON(http.StatusOK, gin.H{
		"data": s,
	})
}

func GetNewBlockHandler() {
	node, err := client.NewClientFromNode("http://95.217.121.243:2071")

	if err != nil {
		return
	}

	clientCtx := client.Context{}
	clientCtx = clientCtx.WithClient(node).WithNodeURI("http://95.217.121.243:2071")

	queryClient := wasmTypes.NewQueryClient(clientCtx)

	data, err := server.GetData(queryClient, 11, 11)

	global.SetBoard(data)

	fmt.Printf("%v", global.GetBoard())
}
