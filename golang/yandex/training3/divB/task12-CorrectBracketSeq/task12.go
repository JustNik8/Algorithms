package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/45468/problems/12/
func main() {
	reader := bufio.NewReader(os.Stdin)

	mappings := map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}

	var seq string
	fmt.Fscan(reader, &seq)

	stack := make([]uint8, 0)

	for i := 0; i < len(seq); i++ {
		bracket := seq[i]
		if bracket == '(' || bracket == '[' || bracket == '{' {
			stack = append(stack, bracket)
		} else {
			if len(stack) <= 0 {
				fmt.Println("no")
				return
			}
			openBracket := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if openBracket != mappings[bracket] {
				fmt.Println("no")
				return
			}
		}
	}

	if len(stack) == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
