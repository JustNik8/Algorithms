package main

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	maxLen := max(len(num1), len(num2))
	if len(num2) > len(num1) {
		num2, num1 = num1, num2
	}

	p1 := len(num1) - 1
	p2 := len(num2) - 1
	ansP := maxLen
	ans := make([]int, maxLen+1)

	for p2 >= 0 {
		d1 := int(num1[p1] - '0')
		d2 := int(num2[p2] - '0')

		ans[ansP] += d1 + d2
		ans[ansP-1] = ans[ansP] / 10
		ans[ansP] = ans[ansP] % 10

		p1--
		p2--
		ansP--
	}

	for p1 >= 0 {
		ans[ansP] += int(num1[p1] - '0')
		ans[ansP-1] = ans[ansP] / 10
		ans[ansP] = ans[ansP] % 10
		p1--
		ansP--
	}

	var sb strings.Builder
	i := 0
	if ans[0] == 0 {
		i++
	}

	for i < len(ans) {
		sb.WriteString(strconv.Itoa(ans[i]))
		i++
	}

	return sb.String()
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
