package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := strings.TrimSpace(scanner.Text())

	bracketCnt := 0
	for i := 0; i < len(str); i++ {
		bracket := str[i]

		if bracket == '(' {
			bracketCnt++
		} else {
			bracketCnt--
		}
		if bracketCnt < 0 {
			break
		}
	}

	if bracketCnt == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
