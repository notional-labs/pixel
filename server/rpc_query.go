package blockexplore

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"code.nkcmr.net/async"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

const lengthChunk = 32

func AsyncGetChuckData(queryClient wasmTypes.QueryClient, chunkX int, chunkY int) async.Promise[chunkData] {
	return async.NewPromise(func() (chunkData, error) {
		res := queryContract(queryClient, chunkX, chunkY)
		// implement here
		return parseDataFromRes(res, chunkX, chunkY), nil
	})
}

func getBlocks(queryClient wasmTypes.QueryClient, maxChunkX int, maxChunkY int) ([]chunkData, error) {
	ctx := context.Background()
	var ans []chunkData
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

func parseDataFromRes(res wasmTypes.QuerySmartContractStateResponse, chunkX, chunkY int) (ans chunkData) {
	jsonByte, _ := res.Data.MarshalJSON()
	var JSONChunkData JsonChunkData
	err := json.Unmarshal(jsonByte, &JSONChunkData)
	if err != nil {
		panic(err)
	}
	for i := range JSONChunkData.grid {
		for j := range JSONChunkData.grid[i] {
			pixel := Pixel{
				x:    i + chunkX*lengthChunk,
				y:    j + chunkY*lengthChunk,
				Info: JSONChunkData.grid[i][j],
			}
			ans.data = append(ans.data, pixel)
		}
	}
	return ans
}
