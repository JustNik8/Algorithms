package main

import "sort"

type Interval struct {
	Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */

type event struct {
	time      int
	eventType int
}

func MinMeetingRooms(intervals []*Interval) int {
	events := make([]event, 0, len(intervals)*2)
	for i := 0; i < len(intervals); i++ {
		events = append(events, event{time: intervals[i].Start, eventType: 1})
		events = append(events, event{time: intervals[i].End, eventType: -1})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].eventType < events[j].eventType
		}
		return events[i].time < events[j].time
	})

	cnt := 0
	minCnt := 0
	for i := 0; i < len(events); i++ {
		cnt += events[i].eventType
		minCnt = max(minCnt, cnt)
	}

	return minCnt
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
