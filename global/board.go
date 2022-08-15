package global

import (
	"github.com/notional-labs/pixel/server"
)

type Board struct {
	Data []server.Pixel `json:"data"`
}

var board = Board{}

func (b *Board) copySlice(slice []server.Pixel) {
	b.Data = make([]server.Pixel, len(slice))
	copy(b.Data, slice)
}

func SetBoard(slice []server.Pixel) {
	// prevent concurrent update
	board.copySlice(slice)
}

func GetBoard() []server.Pixel {
	return board.Data
}
