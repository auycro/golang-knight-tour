package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/auycro/golang-knight-tour/src/algorithm"
	"github.com/auycro/golang-knight-tour/src/board"
)

func main() {
	run_algo := false
	random := false
	argsWithProg := os.Args
	for _, v := range argsWithProg {
		if strings.ToLower(v) == "warnsdoff" {
			run_algo = true
		}

		if strings.ToLower(v) == "random" {
			random = true
		}
	}

	if run_algo {
		runWarnsdoff(random)
	} else {
		gameStart()
	}
}

func gameStart() {
	current_board := board.InitKnightBoard(false)
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

func runWarnsdoff(random bool) {
	current_board := board.InitKnightBoard(random)
	fmt.Printf("Board Init: %s\n", board.BoardToFen(current_board))
	fmt.Print(board.BoardToAscii(current_board))

	turn := 0
	for !board.GAME_END && (turn < (board.ROW*board.COLUMN)-1) {
		time.Sleep(1 * time.Second)
		moves := board.GetValidKnightMoves(current_board)
		next_move := algorithm.WarnsdorffSelectMove(moves, current_board)
		current_board = board.KnightMove(next_move, current_board)
		fmt.Printf("move: %s\n", board.ConvertMoveToAlgebric(next_move))
		//fmt.Printf(" ↪️ Board Status: %s\n", board.BoardToFen(current_board))
		fmt.Print(board.BoardToAscii(current_board))
		turn++
	}
}
