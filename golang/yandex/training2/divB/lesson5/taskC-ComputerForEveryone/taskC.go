package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type data struct {
	cnt int
	pos int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int

	fmt.Fscan(reader, &n, &m)

	groups := make([]data, n)
	classes := make([]data, m)

	for i := 0; i < n; i++ {
		var groupCnt int
		fmt.Fscan(reader, &groupCnt)
		groups[i] = data{cnt: groupCnt, pos: i}
	}

	for i := 0; i < m; i++ {
		var classCnt int
		fmt.Fscan(reader, &classCnt)
		classes[i] = data{cnt: classCnt, pos: i}
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].cnt < groups[j].cnt
	})
	sort.Slice(classes, func(i, j int) bool {
		return classes[i].cnt < classes[j].cnt
	})

	groupPointer := 0
	classPointer := 0

	ans := make([]int, n)
	cnt := 0
	for groupPointer < n && classPointer < m {
		group := groups[groupPointer]
		class := classes[classPointer]

		if group.cnt+1 <= class.cnt {
			ans[group.pos] = class.pos + 1
			groupPointer++
			classPointer++
			cnt++
		} else {
			classPointer++
		}
	}

	fmt.Println(cnt)
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}
