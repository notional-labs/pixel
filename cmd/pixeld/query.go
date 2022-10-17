package main

import (
	"fmt"
	"time"

	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/notional-labs/pixel/server"
	"github.com/notional-labs/pixel/server/serve"

	"github.com/spf13/cobra"
)

func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query ",
		Short: "query",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			node, _ := client.NewClientFromNode("http://95.217.121.243:2071")
			clientCtx := client.Context{}
			clientCtx = clientCtx.WithClient(node).WithNodeURI("http://95.217.121.243:2071")
			now := time.Now()

			queryClient := wasmTypes.NewQueryClient(clientCtx)
			data, _ := server.GetData(queryClient, 11, 11)
			fmt.Println(server.ParsePixelArray(data))
			fmt.Println(time.Now().Sub(now))
			return nil
		},
	}

	return cmd
}

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run ",
		Short: "query",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			node, _ := client.NewClientFromNode("http://95.217.121.243:2071")
			clientCtx := client.Context{}
			clientCtx = clientCtx.WithClient(node).WithNodeURI("http://95.217.121.243:2071")

			queryClient := wasmTypes.NewQueryClient(clientCtx)
			serve.ListenAndServe(queryClient, args[0])

			return nil
		},
	}

	return cmd
}
