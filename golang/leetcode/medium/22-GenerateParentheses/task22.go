package main

import "strings"

func main() {

}

type Stack []string

func generateParenthesis(n int) []string {
	stack := Stack{}
	ans := make([]string, 0)

	var backtrack func(openCnt, closedCnt int)
	backtrack = func(openCnt, closedCnt int) {
		if openCnt == n && closedCnt == n {
			ans = append(ans, strings.Join(stack, ""))
			return
		}

		if openCnt < n {
			stack.push("(")
			backtrack(openCnt+1, closedCnt)
			stack.pop()
		}

		if closedCnt < openCnt {
			stack.push(")")
			backtrack(openCnt, closedCnt+1)
			stack.pop()
		}
	}

	backtrack(0, 0)
	return ans
}

func (s *Stack) pop() {
	size := len(*s)
	*s = (*s)[:size-1]
}

func (s *Stack) push(elem string) {
	*s = append(*s, elem)
}