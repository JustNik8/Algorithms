package main

import "math"

func main() {

}

func canEat(piles []int, speed int, timeLimit int) bool {
	needTime := 0
	for _, pile := range piles {
		needTime += int(math.Ceil(float64(pile) / float64(speed)))
	}

	return needTime <= timeLimit
}

func minEatingSpeed(piles []int, h int) int {
	minSpeed := -1
	l, r := 1, 1_000_000_000

	for l <= r {
		speed := (l + r) / 2

		if canEat(piles, speed, h) {
			r = speed - 1
			minSpeed = speed
		} else {
			l = speed + 1
		}
	}

	return minSpeed
}
