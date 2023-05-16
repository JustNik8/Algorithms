package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(reader, &n, &k)
	var str string
	fmt.Fscan(reader, &str)

	var words = make(map[uint8]int)
	var left = 0
	var right = 0
	var maxLeft = 0
	var maxRight = 0

	for right < n {
		char := str[right]
		cnt, ok := words[char]
		if !ok {
			words[char] = 1
		} else {
			words[char] = cnt + 1
		}
		if cnt+1 > k {
			if maxRight-maxLeft < right-left-1 {
				maxLeft = left
				maxRight = right - 1
			}
			left++
			words[char] = cnt - 1
		} else {
			right++
		}
	}
	if maxRight-maxLeft < right-left-1 {
		maxLeft = left
		maxRight = right - 1
	}

	fmt.Printf("%d %d", maxRight-maxLeft+1, maxLeft+1)
}
