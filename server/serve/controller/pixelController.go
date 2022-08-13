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

	for _, item := range global.GetBoard() {
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
	node, err := client.NewClientFromNode("https://rpc.uni.junonetwork.io:443")

	if err != nil {
		return
	}

	clientCtx := client.Context{}
	clientCtx = clientCtx.WithClient(node).WithNodeURI("https://rpc.uni.junonetwork.io:443")

	queryClient := wasmTypes.NewQueryClient(clientCtx)

	data, err := server.GetData(queryClient, 11, 11)

	global.SetBoard(data)

	fmt.Printf("%v", global.GetBoard())
}
