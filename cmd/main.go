package main

import (
	"fmt"

	board_class "github.com/auycro/golang-knight-travel/src"
)

func main() {
	board := board_class.InitKnightBoard()
	//fmt.Printf("%s\n", board)
	ascii := board_class.ToAscii(board)
	fmt.Print(ascii)
}
