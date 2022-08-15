package global

import (
	"sync"

	"github.com/notional-labs/pixel/server"
)

type Board struct {
	Data []server.Pixel `json:"data"`
}

var lock = &sync.Mutex{}

var board = Board{}

func (b *Board) copySlice(slice []server.Pixel) {
	b.Data = make([]server.Pixel, len(slice))
	copy(b.Data, slice)
}

func SetBoard(slice []server.Pixel) {
	// prevent concurrent update
	lock.Lock()
	board.copySlice(slice)
	lock.Unlock()
}

func GetBoard() []server.Pixel {
	return board.Data
}
