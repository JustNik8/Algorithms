package main

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	ans := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newInterval := intervals[0]
	idx := 0

	for idx < n {
		for idx < n && newInterval[1] >= intervals[idx][0] {
			newInterval = []int{
				min(newInterval[0], intervals[idx][0]),
				max(newInterval[1], intervals[idx][1]),
			}
			idx++
		}

		ans = append(ans, newInterval)

		if idx >= n {
			break
		}

		newInterval = intervals[idx]
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
