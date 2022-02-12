package board

import (
	"log"
	"strconv"
	"strings"
)

type Position struct {
	row    int
	column int
}

var KNIGHT_OFFSET = []Position{
	{-2, -1}, {-2, 1},
	{-1, -2}, {-1, 2},
	{1, -2}, {1, 2},
	{2, -1}, {2, 1},
}

func GetCurrentPosition(piece string, board Board) Position {
	for i := ROW; i > 0; i-- {
		for j := 1; j < COLUMN+1; j++ {
			if board[Position{i, j}] == piece {
				return Position{i, j}
			}
		}
	}
	return Position{}
}

func FenMoveToPosition(fen_move string) Position {
	move := strings.Split(fen_move, "")
	if !IsValidMove(fen_move) {
		log.Fatal("invalid move")
	}
	column := COLUMN_Atoi[move[1]]
	row, _ := strconv.Atoi(move[2])
	return Position{row, column}
}
