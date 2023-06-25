package main

func main() {

}

type empty struct{}
type set map[byte]empty

type squareCoord struct {
	colCoord int
	rowCoord int
}

func isValidSudoku(board [][]byte) bool {
	cols := make(map[int]set)
	rows := make(map[int]set)
	squares := make(map[squareCoord]set)

	size := 9

	for i := 0; i < size; i++ {
		cols[i] = make(set)
		rows[i] = make(set)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sc := squareCoord{
				colCoord: i,
				rowCoord: j,
			}
			squares[sc] = make(set)
		}
	}

	for col := 0; col < size; col++ {
		for row := 0; row < size; row++ {
			char := board[col][row]
			if char == '.' {
				continue
			}

			sc := squareCoord{
				colCoord: col / 3,
				rowCoord: row / 3,
			}

			_, isFoundInCol := cols[col][char]
			_, isFoundInRow := rows[row][char]
			_, isFoundInSquare := squares[sc][char]

			if isFoundInCol || isFoundInRow || isFoundInSquare {
				return false
			}

			cols[col][char] = empty{}
			rows[row][char] = empty{}
			squares[sc][char] = empty{}
		}
	}

	return true
}
