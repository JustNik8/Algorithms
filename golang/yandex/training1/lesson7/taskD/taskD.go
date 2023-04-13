package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const arrive int = 0
const leave = 1

type event struct {
	time      int
	eventType int
}

type maxTime struct {
	time int
	cnt  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buyerCnt int

	fmt.Fscan(reader, &buyerCnt)

	events := make([]event, buyerCnt*2)
	for i := 0; i < buyerCnt*2; i += 2 {
		var timeArrive, timeLeave int
		fmt.Fscan(reader, &timeArrive, &timeLeave)
		events[i] = event{time: timeArrive, eventType: arrive}
		events[i+1] = event{time: timeLeave, eventType: leave}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].eventType < events[j].eventType
		} else {
			return events[i].time < events[j].time
		}
	})

	fmt.Println(events)

}

