package board

import (
	"fmt"
	"strconv"
	"strings"
)

var DEFAULT_POSITION = "8/8/8/8/8/8/8/6N1"

var history = []string{}

func InitKnightBoard() string {
	return DEFAULT_POSITION
}

func ToAscii(board string) string {
	row_num := [8]int{8, 7, 6, 5, 4, 3, 2, 1}
	fen_row := strings.Split(board, "/")
	s := "   +------------------------+\n"
	for i := 0; i < len(row_num); i++ {
		// Add Column Header
		s += fmt.Sprintf(" %d |", row_num[i])

		fen_value := strings.Split(fen_row[i], "")
		for _, v := range fen_value {
			blanks, err := strconv.Atoi(v)
			if err == nil {
				// print blank
				for j := 0; j < blanks; j++ {
					s += " . "
				}

			} else {
				// print piece (we have only N)
				s += fmt.Sprintf(" %s ", "N")
			}
		}

		// Add EOL
		s += "|\n"
	}

	s += "   +------------------------+\n"
	s += "     a  b  c  d  e  f  g  h\n"
	return s
}
