package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const segmentStart = 1
const findPoint = 2
const segmentEnd = 3

type point struct {
	coord     int64
	pointType int64
	ansIndex  int64
}

// https://contest.yandex.ru/contest/27883/problems/B/
func main() {

	reader := bufio.NewReader(os.Stdin)
	var segmentCnt, pointCnt int64
	fmt.Fscan(reader, &segmentCnt, &pointCnt)

	points := make([]point, 2*segmentCnt+pointCnt)
	for i := int64(0); i < 2*segmentCnt; i += 2 {
		var start, end int64
		fmt.Fscan(reader, &start, &end)
		points[i] = point{coord: start, pointType: segmentStart}
		points[i+1] = point{coord: end, pointType: segmentEnd}
	}

	ansIndex := int64(0)
	for i := 2 * segmentCnt; i < 2*segmentCnt+pointCnt; i++ {
		var pointCoord int64
		fmt.Fscan(reader, &pointCoord)
		p := point{coord: pointCoord, pointType: findPoint, ansIndex: ansIndex}
		points[i] = p
		ansIndex++
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].coord == points[j].coord {
			return points[i].pointType < points[j].pointType
		} else {
			return points[i].coord < points[j].coord
		}
	})

	nowSegments := int64(0)
	ans := make([]int64, pointCnt)
	for _, p := range points {
		if p.pointType == findPoint {
			index := p.ansIndex
			ans[index] = nowSegments
		} else if p.pointType == segmentStart {
			nowSegments++
		} else if p.pointType == segmentEnd {
			nowSegments--
		}
	}

	for i := int64(0); i < pointCnt-1; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Printf("%d", ans[pointCnt-1])
}
