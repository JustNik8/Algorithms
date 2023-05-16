package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const start = 1
const end = 0

type event struct {
	position  int
	typeEvent int
}

// https://contest.yandex.ru/contest/27883/problems/A/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	events := make([]event, 2*m)
	for i := 0; i < 2*m; i += 2 {
		var startPos, endPos int
		fmt.Fscan(reader, &startPos, &endPos)
		start := event{position: startPos, typeEvent: start}
		end := event{position: endPos, typeEvent: end}
		events[i] = start
		events[i+1] = end
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].position == events[j].position {
			return events[i].typeEvent > events[j].typeEvent
		} else {
			return events[i].position < events[j].position
		}
	})

	canCheat := events[0].position
	nowWatch := 1
	for i := 1; i < len(events); i++ {
		if nowWatch == 0 {
			canCheat += events[i].position - events[i-1].position - 1
		}
		if events[i].typeEvent == start {
			nowWatch++
		} else if events[i].typeEvent == end {
			nowWatch--
		}
	}

	canCheat += n - events[2*m-1].position - 1
	fmt.Print(canCheat)
}
