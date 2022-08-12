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
	c.JSON(http.StatusOK, gin.H{
		"data": global.Board,
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

	if err != nil {
		fmt.Printf("")
	}

	var buffer bytes.Buffer
	var b []byte

	for _, item := range data {
		b, err = json.Marshal(item)
		if err != nil {
			fmt.Printf("")
		}

		buffer.WriteString(string(b) + ",")
	}

	s := strings.TrimSpace(buffer.String())
	// trim last comma
	s = s[:len(s)-1]

	fmt.Printf(s)

	copy(data, global.Board)
}
