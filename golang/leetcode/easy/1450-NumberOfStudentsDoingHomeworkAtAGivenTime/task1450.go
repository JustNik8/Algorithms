package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt := 0

	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			cnt++
		}
	}

	return cnt
}
