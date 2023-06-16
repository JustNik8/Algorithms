package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := compress([]byte{'a', 'b', 'c'})
	fmt.Println(ans)
}

func compress(chars []byte) int {
	pointer := 0

	size := len(chars)

	cnt := 1
	for i := 1; i < size; i++ {
		if chars[i] == chars[i-1] {
			cnt++
		} else {
			chars[pointer] = chars[i-1]
			pointer++

			if cnt != 1 {
				s := strconv.Itoa(cnt)
				for j := 0; j < len(s); j++ {
					chars[pointer] = s[j]
					pointer++
				}
			}

			cnt = 1
		}
	}

	chars[pointer] = chars[len(chars)-1]
	pointer++
	if cnt != 1 {
		s := strconv.Itoa(cnt)
		for j := 0; j < len(s); j++ {
			chars[pointer] = s[j]
			pointer++
		}
	}

	newChars := chars[:pointer]

	return len(newChars)
}
