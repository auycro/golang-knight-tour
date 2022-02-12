package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Board = map[Position]string

func BoardToAscii(board Board) string {
	s := "   +------------------------+\n"
	for i := ROW; i > 0; i-- {
		// Add ROW Header
		s += fmt.Sprintf(" %d |", i)
		for j := 1; j < COLUMN+1; j++ {
			if board[Position{i, j}] == "" {
				s += " . "
			} else {
				s += fmt.Sprintf(" %s ", board[Position{i, j}])
			}
		}
		// Add EOL
		s += "|\n"
	}
	s += "   +------------------------+\n"
	s += "     a  b  c  d  e  f  g  h\n"
	return s
}

func BoardToFen(board Board) string {
	result := ""
	for i := ROW; i > 0; i-- {
		blank := 0
		for j := 1; j < COLUMN+1; j++ {
			if board[Position{i, j}] == "" {
				blank++
			} else {
				if blank > 0 {
					result += strconv.Itoa(blank)
				}
				result += board[Position{i, j}]
				blank = 0
			}
		}
		if blank > 0 {
			result += strconv.Itoa(blank)
		}
		if i-1 > 0 {
			result += "/"
		}
	}
	return result
}

func FenToAscii(board_fen string) string {
	fen_row := strings.Split(board_fen, "/")
	s := "   +------------------------+\n"
	for i := 0; i < ROW; i++ {
		// Add Column Header
		s += fmt.Sprintf(" %d |", ROW-i)

		fen_value := strings.Split(fen_row[i], "")
		for _, v := range fen_value {
			blanks, err := strconv.Atoi(v)
			if err == nil {
				// print blank
				for j := 0; j < blanks; j++ {
					s += " . "
				}

			} else {
				// print symbol
				s += fmt.Sprintf(" %s ", v)
			}
		}

		// Add EOL
		s += "|\n"
	}

	s += "   +------------------------+\n"
	s += "     a  b  c  d  e  f  g  h\n"
	return s
}
