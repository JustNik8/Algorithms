package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type student struct {
	pos   int
	index int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var studentCnt, talkDist int

	fmt.Fscan(reader, &studentCnt, &talkDist)

	students := make([]student, studentCnt)
	for i := 0; i < studentCnt; i++ {
		var pos int
		fmt.Fscan(reader, &pos)
		students[i] = student{
			pos:   pos,
			index: i,
		}
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i].pos < students[j].pos
	})

	currentType := 1
	maxTypes := 1
	types := make([]int, studentCnt)

	firstStudent := students[0]
	types[firstStudent.index] = currentType

	for i := 1; i < studentCnt; i++ {
		if students[i].pos-students[i-1].pos <= talkDist {
			currentType++
		} else {
			maxTypes = int(math.Max(float64(currentType), float64(maxTypes)))
			currentType = 1
		}
		index := students[i].index
		types[index] = currentType
	}

	maxTypes = int(math.Max(float64(currentType), float64(maxTypes)))
	fmt.Println(maxTypes)

	for _, t := range types {
		fmt.Printf("%d ", t)
	}
}
