package main

import (
	"flag"
	"fmt"
	"github.com/alexk307/warnsdorff/board"
	"math"
)

func main() {
	size := flag.Int("size", 10, "Size of board")
	startX := flag.Int("startX", 1, "Starting X position on board. Must be less than the size and greater than 0")
	startY := flag.Int("startY", 1, "Starting Y position on board. Must be less than the size and greater than 0")

	flag.Parse()
	startingX := *startX
	startingY := *startY
	visitCount := 1
	boardSize := *size
	b := board.NewBoard(boardSize, startingX, startingY)

	for {
		candidates := b.CurrentCandidates()
		least := math.MaxInt8
		leastTile := board.Tile{}
		for _, candidate := range candidates {
			nextMoves := b.Candidates(candidate.X, candidate.Y)
			if len(nextMoves) < least {
				least = len(nextMoves)
				leastTile = candidate
			}
		}
		visitCount += 1
		b.Tiles[leastTile.X][leastTile.Y].Visit = visitCount
		b.MoveToTile(leastTile)

		if b.PositionX == startingX && b.PositionY == startingY {
			fmt.Println("Found a closed path")
			break
		}

		if visitCount == boardSize*boardSize {
			fmt.Println("Found an open path")
			break
		}

	}

	for i := 0; i < boardSize; i++ {
		fmt.Println()
		for j := 0; j < boardSize; j++ {
			fmt.Printf(" %3d", b.Tiles[i][j].Visit)
		}
	}
	fmt.Println()

}
