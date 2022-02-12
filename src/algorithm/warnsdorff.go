package algorithm

import (
	"github.com/auycro/golang-knight-tour/src/board"
)

func WarnsdorffSelectMove(moves []board.Position, current_board board.Board) board.Position {
	result := board.Position{}
	min := 8
	for i := 0; i < len(moves); i++ {
		temp_board := CopyBoard(current_board)
		//fmt.Printf("BOARD:%s\n", board.BoardToFen(temp_board))
		next_board := board.KnightMove(moves[i], temp_board)
		next_move := board.GetValidKnightMoves(next_board)
		if min > len(next_move) {
			min = len(next_move)
			result = moves[i]
		}
	}
	return result
}

func CopyBoard(current_board board.Board) board.Board {
	result := board.Board{}
	for k, v := range current_board {
		result[k] = v
	}
	return result
}
