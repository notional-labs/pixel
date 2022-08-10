package server

type Pixel struct {
	x    int
	y    int
	Info PixelInfo `json:"info"`
}

type chunkData struct {
	data []Pixel
}

type PixelInfo struct {
	Color   uint8  `json:"color"`
	Painter string `json:"painter"`
}
type JsonChunkData struct {
	Grid [][]PixelInfo `json:"grid"`
}
