package main

import "sort"

func main() {

}

type car struct {
	position int
	speed    int
	time     float64
}

func carFleet(target int, position []int, speed []int) int {
	size := len(position)
	cars := make([]car, size)

	for i := 0; i < size; i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars[i] = car{position: position[i], speed: speed[i], time: time}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	carFleet := 1
	currMax := cars[0].time

	for i := 1; i < size; i++ {
		if cars[i].time > currMax {
			currMax = cars[i].time
			carFleet++
		}
	}

	return carFleet
}
