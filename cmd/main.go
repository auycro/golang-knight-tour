package main

import (
	"fmt"

	"github.com/auycro/golang-knight-travel/src/board"
)

func main() {
	current_board := board.InitKnightBoard()
	current_board = board.Move("Nh3", current_board)
	current_board = board.Move("Ng5", current_board)
	ascii := board.ToAscii(current_board)
	fmt.Print(ascii)
}
