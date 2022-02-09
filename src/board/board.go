package board

import (
	"fmt"
	"strconv"
	"strings"
)

const DEFAULT_POSITION string = "8/8/8/8/8/8/8/6N1"
const ROW int = 8

var COLUMN = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

var History = []string{}

func InitKnightBoard() string {
	History = []string{}
	return DEFAULT_POSITION
}

func ToAscii(board string) string {
	fen_row := strings.Split(board, "/")
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

//TODO: Restructure
func Move(move string, board string) string {
	result := ""
	History = append(History, move)
	board = strings.Replace(board, "N", "X", -1)
	old_board := strings.Split(board, "/")

	fen_move := strings.Split(move, "")
	new_x := COLUMN[fen_move[1]]
	new_y, _ := strconv.Atoi(fen_move[2])

	for i := 0; i < ROW; i++ {
		if new_y == (ROW - i) {
			blank := 0
			for j := 0; j < len(COLUMN); j++ {
				if new_x == j {
					result += strconv.Itoa(blank)
					result += "N"
					blank = 0
				} else {
					blank++
				}
			}
			if blank != 0 {
				result += strconv.Itoa(blank)
			}
		} else {
			result += old_board[i]
		}

		if i < ROW-1 {
			result += "/"
		}
	}

	return result
}
