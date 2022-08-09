package blockexplore

type pixel struct {
	x        int
	y        int
	addresss string
	color    int
}

type chunkData struct {
	chunkX int
	chunkY int
	data   []pixel
}
