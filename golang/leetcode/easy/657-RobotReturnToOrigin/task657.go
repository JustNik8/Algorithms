package main

func judgeCircle(moves string) bool {
	x, y := 0, 0

	for i := 0; i < len(moves); i++ {
		move := moves[i]
		if move == 'R' {
			x++
		} else if move == 'L' {
			x--
		} else if move == 'U' {
			y++
		} else if move == 'D' {
			y--
		}
	}

	return x == 0 && y == 0
}
