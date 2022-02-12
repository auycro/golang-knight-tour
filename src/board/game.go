package board

var GAME_END bool = false

var HISTORY = []Position{}

const DEFAULT_POSITION string = "8/8/8/8/8/8/8/6N1"

var STARTING_POSITION = Position{1, 7}

func InitKnightBoard() Board {
	HISTORY = []Position{}
	CurrentBoard := Board{}
	for i := ROW; i > 0; i-- {
		for j := 1; j < COLUMN+1; j++ {
			CurrentBoard[Position{i, j}] = ""
		}
	}
	CurrentBoard[STARTING_POSITION] = "N"
	return CurrentBoard
}
