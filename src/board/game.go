package board

import (
	"math/rand"
	"time"
)

var GAME_END bool = false
var HISTORY = []Position{}
var STARTING_POSITION = Position{1, 7}

func InitKnightBoard(random bool) Board {
	HISTORY = []Position{}
	CurrentBoard := Board{}
	for i := ROW; i > 0; i-- {
		for j := 1; j < COLUMN+1; j++ {
			CurrentBoard[Position{i, j}] = ""
		}
	}
	if random {
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(ROW-1) + 1
		column := rand.Intn(COLUMN-1) + 1
		STARTING_POSITION = Position{row, column}
	}
	CurrentBoard[STARTING_POSITION] = "N"
	return CurrentBoard
}
