package main

import (
	"bufio"
	"fmt"
	"os"
)

type section struct {
	start  int
	end    int
	isWork bool
}

// https://contest.yandex.ru/contest/45468/problems/6/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var sectionCnt int
	fmt.Fscan(reader, &sectionCnt)
	fmt.Fscan(reader, &sectionCnt)

	sections := make([]section, 0, sectionCnt)
	cnt := 0
	for i := 0; i < sectionCnt; i++ {
		var start, end int
		fmt.Fscan(reader, &start, &end)
		for j := 0; j < i; j++ {
			s := &sections[j]
			if start <= s.end && s.start <= end && s.isWork {
				s.isWork = false
				cnt--
			}
		}
		sections = append(sections, section{start: start, end: end, isWork: true})
		cnt++
	}

	fmt.Println(cnt)
}
