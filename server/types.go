package server

type Pixel struct {
	X     int
	Y     int
	Color uint8 `json:"color"`
}

type Result struct {
	X     uint8 `json:"x"`
	Y     uint8 `json:"y"`
	Color uint8 `json:"color"`
}

type PixelInfo struct {
	Color   uint8  `json:"color"`
	Painter string `json:"painter"`
}
type JsonChunkData struct {
	Grid [][]PixelInfo `json:"grid"`
}
