package blockexplore

import (
	"context"

	"code.nkcmr.net/async"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"

	rpcclient "github.com/tendermint/tendermint/rpc/client"

	tmtypes "github.com/tendermint/tendermint/types"
)

func AsyncGetBlockData(node rpcclient.Client, startHeight int, endHeight int) async.Promise[tmtypes.Txs] {
	return async.NewPromise(func() (tmtypes.Txs, error) {
		var ans tmtypes.Txs
		// header -> BlockchainInfo
		// header, tx -> Block
		// results -> BlockResults
		for height := startHeight; height < endHeight; height++ {

			tmp := int64(height)
			h := &tmp

			res, err := node.Block(context.Background(), h)
			if err != nil {
				return nil, err
			}
			ans = append(ans, res.Block.Txs...)
		}

		return ans, nil
	})
}

func getBlocks(node rpcclient.Client, startHeight int, endHeight int, numberProcess int) (tmtypes.Txs, error) {
	ctx := context.Background()

	var txsPromisr = make([]async.Promise[tmtypes.Txs], numberProcess)
	var ans tmtypes.Txs

	temp := (endHeight - startHeight) / numberProcess

	for i := 0; i < numberProcess; i++ {
		if i == numberProcess-1 {
			txsPromisr[i] = AsyncGetBlockData(node, startHeight+i*temp, endHeight)
		}
		txsPromisr[i] = AsyncGetBlockData(node, startHeight+i*temp, startHeight+(i+1)*temp)
	}

	for i := 0; i < numberProcess; i++ {
		txs, err := txsPromisr[i].Await(ctx)
		if err != nil {
			return nil, err
		}
		ans = append(ans, txs...)
	}
	return ans, nil
}

func queryContract(queryClient wasmTypes.QueryClient, node rpcclient.Client, chunkX int64, chunkY int64) wasmTypes.QuerySmartContractStateResponse {
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

func parseQueryData(chunkX, chunkY int64) []byte {
	return []byte("")
}
