package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://contest.yandex.ru/contest/45468/problems/13/
func main() {
	reader := bufio.NewReader(os.Stdin)
	stack := make([]string, 0)

	var elem string
	for {
		_, err := fmt.Fscan(reader, &elem)
		if err != nil {
			break
		}

		if elem == "*" || elem == "+" || elem == "-" {
			makeOp(&stack, elem)
		} else {
			stack = append(stack, elem)
		}
		//fmt.Println(stack)
	}

	fmt.Println(stack[0])
}

func makeOp(s *[]string, oper string) {
	first, second := getOperands(s)
	var res int

	switch oper {
	case "*":
		res = second * first
	case "+":
		res = second + first
	case "-":
		res = second - first
	}

	*s = append(*s, strconv.Itoa(res))
}

func getOperands(s *[]string) (int, int) {
	size := len(*s)
	first, _ := strconv.Atoi((*s)[size-1])
	second, _ := strconv.Atoi((*s)[size-2])

	*s = (*s)[:size-2]
	return first, second
}

// test 11
// 1 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
