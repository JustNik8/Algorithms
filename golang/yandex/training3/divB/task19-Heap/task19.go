package main

import (
	"bufio"
	"fmt"
	"os"
)

type heap struct {
	elems []int
}

// https://contest.yandex.ru/contest/45468/problems/19/
func main() {
	reader := bufio.NewReader(os.Stdin)

	h := heap{}

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var cmd int
		fmt.Fscan(reader, &cmd)
		if cmd == 0 {
			var num int
			fmt.Fscan(reader, &num)
			h.insert(num)
		} else {
			fmt.Println(h.extract())
		}
	}
}

func (h *heap) insert(n int) {
	h.elems = append(h.elems, n)
	pos := len(h.elems) - 1

	parentPos := (pos - 1) / 2
	for pos > 0 && h.elems[pos] > h.elems[parentPos] {
		h.elems[pos], h.elems[parentPos] = h.elems[parentPos], h.elems[pos]
		pos = (pos - 1) / 2
		parentPos = (pos - 1) / 2
	}
}

func (h *heap) extract() int {
	ans := h.elems[0]
	h.elems[0] = h.elems[len(h.elems)-1]
	pos := 0

	for pos*2+2 < len(h.elems) {
		maxSonIndex := pos*2 + 1
		if h.elems[pos*2+2] > h.elems[maxSonIndex] {
			maxSonIndex = pos*2 + 2
		}

		if h.elems[pos] < h.elems[maxSonIndex] {
			h.elems[pos], h.elems[maxSonIndex] = h.elems[maxSonIndex], h.elems[pos]
			pos = maxSonIndex
		} else {
			break
		}
	}

	h.elems = h.elems[:len(h.elems)-1]
	return ans
}
