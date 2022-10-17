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
	addressContract := "juno13jr4jmz2pu5vu7avejy20s583rvpmp5ctmhzjpjhygqnxg5h5rzsk70cng"
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
	var jsonByte []byte
	if res != nil {
		jsonByte, _ = res.Data.MarshalJSON()
	}

	var JSONChunkData JsonChunkData
	err := json.Unmarshal(jsonByte, &JSONChunkData)
	if err != nil {
		panic(err)
	}

	for i := range JSONChunkData.Grid {
		for j := range JSONChunkData.Grid[i] {
			if JSONChunkData.Grid[i][j].Color == 0 {
				continue
			}
			pixel := Pixel{
				X:     j + chunkX*lengthChunk,
				Y:     i + chunkY*lengthChunk,
				Color: JSONChunkData.Grid[i][j].Color,
			}
			ans = append(ans, pixel)
		}
	}
	return ans
}

func ParsePixelArray(pixelArray []Pixel) map[int]([]map[int]uint8) {
	ans := make(map[int]([]map[int]uint8))

	for _, pixel := range pixelArray {
		newPixel := make(map[int]uint8)
		newPixel[pixel.Y] = pixel.Color
		ans[pixel.X] = append(ans[pixel.X], newPixel)
	}

	return ans
}
