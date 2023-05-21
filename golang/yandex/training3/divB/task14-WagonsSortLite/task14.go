package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []int

// https://contest.yandex.ru/contest/45468/problems/14/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	s := stack{}
	cnt := 0

	for i := 0; i < n; i++ {
		var wagon int
		fmt.Fscan(reader, &wagon)
		if wagon-1 == cnt {
			cnt++
			canTake := true
			for len(s) > 0 && canTake {
				num := s.last()
				if num-1 == cnt {
					cnt++
					s = s[:len(s)-1]
				} else {
					canTake = false
				}
			}
		} else {
			s = append(s, wagon)
		}
	}

	size := len(s)
	for i := 0; i < size; i++ {
		num := s.last()
		if num-1 != cnt {
			fmt.Println("NO")
			return
		}
		s = s[:len(s)-1]
	}
	fmt.Println("YES")
}

func (s stack) last() int {
	return s[len(s)-1]
}
