package main

import "strconv"

func main() {

}

type Stack []int

func evalRPN(tokens []string) int {
	stack := Stack{}

	for _, token := range tokens {
		if isOperator(token) {
			num1 := stack.pop()
			num2 := stack.pop()

			if token == "+" {
				stack.push(num2 + num1)
			} else if token == "-" {
				stack.push(num2 - num1)
			} else if token == "*" {
				stack.push(num2 * num1)
			} else if token == "/" {
				stack.push(num2 / num1)
			}
		} else {
			num, _ := strconv.Atoi(token)
			stack.push(num)
		}
	}

	return stack[0]
}

func (s *Stack) push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) pop() int {
	size := len(*s)
	num := (*s)[size-1]

	*s = (*s)[:size-1]
	return num
}

func (s *Stack) top() int {
	size := len(*s)
	return (*s)[size-1]
}

func isOperator(token string) bool {
	_, err := strconv.Atoi(token)
	return err != nil
}
