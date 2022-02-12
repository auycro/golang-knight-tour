package board

import (
	"fmt"
	"regexp"
	"strconv"
)

func IsValidMove(fen_move string) bool {
	r, _ := regexp.Compile("N[a-h][1-8]")
	return r.MatchString(fen_move)
}

func GetValidKnightMoves(board Board) []Position {
	result := []Position{}
	current_position := GetCurrentPosition("N", board)
	//fmt.Printf("Current position: %s\n", ConvertMoveToAlgebric(current_position))

	for i := 0; i < len(KNIGHT_OFFSET); i++ {
		possible_row := current_position.row + KNIGHT_OFFSET[i].row
		possible_column := current_position.column + KNIGHT_OFFSET[i].column
		if (possible_row > 0) && (possible_row < 9) {
			if (possible_column > 0) && (possible_column < 9) {
				possible := Position{possible_row, possible_column}
				// No re-visit
				if board[possible] == "" {
					//fmt.Printf("Possible: %s\n", ConvertMoveToAlgebric(possible))
					result = append(result, possible)
				}
			}
		}
	}
	//fmt.Printf("Possible %d moves\n", len(result))
	return result
}

func ConvertMoveToAlgebric(pos Position) string {
	//column := utility.GetFirstKeyByValue(COLUMN_Atoi, pos.column)
	return fmt.Sprintf("%s%s", COLUMN_Itoa[pos.column], strconv.Itoa(pos.row))
}

func KnightMove(move Position, board Board) Board {
	HISTORY = append(HISTORY, move)

	current_knight_position := GetCurrentPosition("N", board)
	board[current_knight_position] = "X"
	board[move] = "N"
	return board
}
