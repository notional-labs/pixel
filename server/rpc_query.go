package blockexplore

import (
	"context"
	"encoding/hex"
	"fmt"

	"code.nkcmr.net/async"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

func AsyncGetChuckData(queryClient wasmTypes.QueryClient, chunkX int, chunkY int) async.Promise[chunkData] {
	return async.NewPromise(func() (chunkData, error) {
		res := queryContract(queryClient, chunkX, chunkY)
		// implement here
		return parseDataFromRes(res), nil
	})
}

func getBlocks(node rpcclient.Client, maxChunkX int, maxChunkY int) ([]chunkData, error) {
	ctx := context.Background()
	var ans []chunkData
	var queryClient wasmTypes.QueryClient
	var chunkPromisr = make([]async.Promise[chunkData], maxChunkX*maxChunkY)

	for y := 0; y <= maxChunkY; y++ {
		for x := 0; x <= maxChunkX; x++ {
			chunkPromisr[x+y*maxChunkX] = AsyncGetChuckData(queryClient, x, y)
		}
	}

	for i := 0; i < maxChunkY*maxChunkY; i++ {
		chunkData, err := chunkPromisr[i].Await(ctx)
		if err != nil {
			return nil, err
		}
		ans = append(ans, chunkData)
	}
	return ans, nil
}

func queryContract(queryClient wasmTypes.QueryClient, chunkX, chunkY int) wasmTypes.QuerySmartContractStateResponse {
	// TODO: fill query data
	addressContract := ""
	res, _ := queryClient.SmartContractState(
		context.Background(),
		&wasmTypes.QuerySmartContractStateRequest{
			Address:   addressContract,
			QueryData: parseQueryData(chunkX, chunkY),
		},
	)

	// handler logic
	return *res
}

func parseQueryData(chunkX, chunkY int) []byte {
	rawStr := fmt.Sprintf(`{"get_chunk": {"x": %d,"y": %d}}`, chunkX, chunkY)
	decoder := newArgDecoder(hex.DecodeString)
	queryData, err := decoder.DecodeString(rawStr)
	if err != nil {
		panic(err)
	}
	return queryData
}

func parseDataFromRes(wasmTypes.QuerySmartContractStateResponse) chunkData {

	return chunkData{}
}
