package blockexplore

type Pixel struct {
	x    int
	y    int
	Info PixelInfo
}

type chunkData struct {
	chunkX int
	chunkY int
	data   []Pixel
}

type PixelInfo struct {
	color   uint8  `json:"color"`
	painter string `json:"painter"`
}
type JsonChunkData struct {
	grid [][]PixelInfo `json:"grid"`
}
