package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ans := make([]int, len(num1)+len(num2))

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			idx1 := len(num1) - i - 1
			idx2 := len(num2) - j - 1

			d1 := int(num1[idx1] - '0')
			d2 := int(num2[idx2] - '0')

			ans[i+j] += d1 * d2
			ans[i+j+1] += ans[i+j] / 10
			ans[i+j] = ans[i+j] % 10
		}
	}

	fmt.Println(ans)
	pointer := len(ans) - 1
	for ans[pointer] == 0 {
		pointer--
	}

	ansStr := make([]string, pointer+1)
	for pointer >= 0 {
		ansStr[len(ansStr)-pointer-1] = strconv.Itoa(ans[pointer])
		pointer--
	}

	return strings.Join(ansStr, "")
}
