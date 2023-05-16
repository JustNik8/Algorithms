package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/27794/problems/E/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)

	trees := make([]int, n)
	for i := 0; i < n; i++ {
		var tree int
		fmt.Fscan(reader, &tree)
		trees[i] = tree
	}

	if n == 1 {
		fmt.Println("1 1")
		return
	}

	left := 0
	right := 0
	minLeft := -1
	minRight := -1
	sorts := make(map[int]int)
	sorts[trees[0]] = 1

	for right < n {
		if len(sorts) < k {
			right++
			if right >= n {
				continue
			}
			sortRight := trees[right]
			_, ok := sorts[sortRight]
			if !ok {
				sorts[sortRight] = 0
			}
			sorts[sortRight] += 1
		} else {
			sortLeft := trees[left]
			if right-left < minRight-minLeft || minLeft == -1 {
				minLeft = left
				minRight = right
			}
			cnt, _ := sorts[sortLeft]
			if cnt == 1 {
				delete(sorts, sortLeft)
			} else {
				sorts[sortLeft] -= 1
			}
			left++
		}
	}

	fmt.Printf("%d %d", minLeft+1, minRight+1)
}
