package main

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	ans := make([][]int, 0)

	idx := 0
	for idx < n && intervals[idx][1] < newInterval[0] {
		ans = append(ans, intervals[idx])
		idx++
	}
	for idx < n && intervals[idx][0] <= newInterval[1] {
		newInterval = []int{
			min(intervals[idx][0], newInterval[0]),
			max(intervals[idx][1], newInterval[1]),
		}
		idx++
	}

	ans = append(ans, newInterval)

	for idx < n {
		ans = append(ans, intervals[idx])
		idx++
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
