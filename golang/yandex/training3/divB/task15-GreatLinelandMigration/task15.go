package main

import (
	"bufio"
	"fmt"
	"os"
)

type city struct {
	index int
	cost  int
}

type stack []city

// https://contest.yandex.ru/contest/45468/problems/15/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	s := stack{}
	for i := 0; i < n; i++ {
		var cost int
		fmt.Fscan(reader, &cost)
		for len(s) > 0 && cost < s.top().cost {
			c := s.pop()
			ans[c.index] = i
		}
		s = append(s, city{index: i, cost: cost})
	}

	for _, an := range ans {
		fmt.Printf("%d ", an)
	}
}

func (s *stack) pop() city {
	num := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return num
}

func (s *stack) top() city {
	return (*s)[len(*s)-1]
}
