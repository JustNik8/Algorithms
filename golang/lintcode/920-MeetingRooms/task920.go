package main

import "sort"

type Interval struct {
	Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: if a person could attend all meetings
 */
func CanAttendMeetings(intervals []*Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}

	return true
}
