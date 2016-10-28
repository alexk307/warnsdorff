package board

type Tile struct {
	Visit int
	X     int
	Y     int
}

type Board struct {
	Tiles     [][]Tile
	PositionX int
	PositionY int
}

func NewBoard(size, startX, startY int) Board {
	b := Board{PositionX: startX, PositionY: startY}
	for i := 0; i < size; i++ {
		b.Tiles = append(b.Tiles, []Tile{})
		for j := 0; j < size; j++ {
			b.Tiles[i] = append(b.Tiles[i], Tile{Visit: 0, X: i, Y: j})
		}
	}
	b.Tiles[startX][startY].Visit = 1
	return b
}

func (b *Board) CurrentCandidates() []Tile {
	return b.Candidates(b.PositionX, b.PositionY)
}

func (b *Board) Candidates(x, y int) []Tile {

	potentialMoves := []Tile{}

	if x >= len(b.Tiles) || y >= len(b.Tiles) {
		return potentialMoves
	}

	// Try up potentialMoves
	if x+1 <= len(b.Tiles)-1 && y+2 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x+1][y+2])
	}
	if x-1 >= 0 && y+2 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x-1][y+2])
	}

	// Try down potentialMoves
	if x+1 <= len(b.Tiles)-1 && y-2 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x+1][y-2])
	}
	if x-1 >= 0 && y-2 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x-1][y-2])
	}

	// Try right potentialMoves
	if x+2 <= len(b.Tiles)-1 && y+1 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x+2][y+1])
	}
	if x+2 <= len(b.Tiles)-1 && y-1 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x+2][y-1])
	}

	// Try left potentialMoves
	if x-2 >= 0 && y+1 <= len(b.Tiles)-1 {
		potentialMoves = append(potentialMoves, b.Tiles[x-2][y+1])
	}
	if x-2 >= 0 && y-1 >= 0 {
		potentialMoves = append(potentialMoves, b.Tiles[x-2][y-1])
	}

	moves := []Tile{}
	//Remove all potential moves that have been visited before
	for _, move := range potentialMoves {
		if move.Visit == 0 {
			moves = append(moves, move)
		}
	}
	return moves
}

func (b *Board) MoveToTile(tile Tile) {
	tile.Visit = 1
	b.PositionX = tile.X
	b.PositionY = tile.Y
}
