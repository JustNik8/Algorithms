package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	size := len(hand)
	if size%groupSize != 0 {
		return false
	}

	countMap := make(map[int]int)
	for _, card := range hand {
		countMap[card] += 1
	}

	keys := make([]int, 0, len(countMap))
	for k, _ := range countMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	groups := size / groupSize
	p := 0
	for i := 0; i < groups; i++ {
		for countMap[keys[p]] == 0 {
			p++
		}

		num := keys[p]

		for j := num; j < num+groupSize; j++ {
			if countMap[j] == 0 {
				return false
			}
			countMap[j] -= 1
		}
	}

	return true
}
