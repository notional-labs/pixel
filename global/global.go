package global

import "github.com/notional-labs/pixel/server"

var board server.ChunkData

func SetBoard(slice []server.Pixel) {
	copy(board.Data, slice)
}

func GetBoard() []server.Pixel {
	return board.Data
}
