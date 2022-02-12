package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/auycro/golang-knight-travel/src/board"
)

func main() {
	gameStart()
}

func gameStart() {
	current_board := board.InitKnightBoard()
	fmt.Printf("Board Init: %s\n", board.BoardToFen(current_board))
	fmt.Print(board.BoardToAscii(current_board))

	scanner := bufio.NewScanner(os.Stdin)
	for !board.GAME_END {

		moves := board.GetValidKnightMoves(current_board)

		if len(moves) == 0 {
			break
		}

		for i := 0; i < len(moves); i++ {
			fmt.Printf("{%d:N%s} ", i+1, board.ConvertMoveToAlgebric(moves[i]))
		}
		fmt.Println()

		fmt.Printf("%d) Enter your move: ", len(board.HISTORY)+1)
		scanner.Scan()
		text := scanner.Text()
		choice, err := strconv.Atoi(text)
		if err == nil && choice <= len(moves) && choice > 0 {
			current_board = board.KnightMove(moves[choice-1], current_board)
			fmt.Printf(" ↪️ Board Status: %s\n", board.BoardToFen(current_board))
			fmt.Print(board.BoardToAscii(current_board))
		} else {
			log.Fatal("invalid move")
			break
		}
	}

	game_status := "Lose"
	if board.GAME_END {
		game_status = "Win"
	}
	fmt.Printf("-----End-----: %s\n", game_status)
}
