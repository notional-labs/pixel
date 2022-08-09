package server

import (
	"context"
	"encoding/json"
	"fmt"

	"code.nkcmr.net/async"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

const lengthChunk = 32

func AsyncGetChuckData(queryClient wasmTypes.QueryClient, chunkX int, chunkY int) async.Promise[[]Pixel] {
	return async.NewPromise(func() ([]Pixel, error) {
		res := QueryContract(queryClient, chunkX, chunkY)
		return ParseDataFromRes(res, chunkX, chunkY), nil
	})
}

func GetData(queryClient wasmTypes.QueryClient, maxChunkX int, maxChunkY int) ([]Pixel, error) {
	ctx := context.Background()

	var ans []Pixel
	var chunkPromisr = make([]async.Promise[[]Pixel], maxChunkX*maxChunkY)

	for y := 0; y < maxChunkY; y++ {
		for x := 0; x < maxChunkX; x++ {
			chunkPromisr[x+y*maxChunkX] = AsyncGetChuckData(queryClient, x, y)
		}
	}

	for i := 0; i < maxChunkY*maxChunkY; i++ {
		data, err := chunkPromisr[i].Await(ctx)
		if err != nil {
			return nil, err
		}
		if len(data) == 0 {
			continue
		}
		ans = append(ans, data...)
	}
	return ans, nil
}

func QueryContract(queryClient wasmTypes.QueryClient, chunkX, chunkY int) *wasmTypes.QuerySmartContractStateResponse {
	addressContract := "juno1w7xyscaxkwruma9g4m530syjgv58e2s50rt2tr2u3e4dwnqe80lqyyhaye"
	res, _ := queryClient.SmartContractState(
		context.Background(),
		&wasmTypes.QuerySmartContractStateRequest{
			Address:   addressContract,
			QueryData: parseQueryData(chunkX, chunkY),
		},
	)
	// handler logic
	return res
}

func parseQueryData(chunkX, chunkY int) []byte {
	rawStr := fmt.Sprintf(`{"get_chunk":{"x": %d,"y": %d}}`, chunkX, chunkY)
	decoder := newArgDecoder(asciiDecodeString)
	queryData, err := decoder.DecodeString(rawStr)
	if err != nil {
		panic(err)
	}
	return queryData
}

func ParseDataFromRes(res *wasmTypes.QuerySmartContractStateResponse, chunkX, chunkY int) (ans []Pixel) {
	jsonByte, err := res.Data.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var JSONChunkData JsonChunkData
	err = json.Unmarshal(jsonByte, &JSONChunkData)
	if err != nil {
		panic(err)
	}

	for i := range JSONChunkData.Grid {
		for j := range JSONChunkData.Grid[i] {
			if JSONChunkData.Grid[i][j].Color == 0 {
				continue
			}
			pixel := Pixel{
				x:    i + chunkX*lengthChunk,
				y:    j + chunkY*lengthChunk,
				Info: JSONChunkData.Grid[i][j],
			}
			ans = append(ans, pixel)
		}
	}
	return ans
}
