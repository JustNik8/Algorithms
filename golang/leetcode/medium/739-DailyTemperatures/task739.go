package main

func main() {

}

type Node struct {
	data  int
	index int
}

type Stack []Node

func dailyTemperatures(temperatures []int) []int {
	stack := Stack{}
	ans := make([]int, len(temperatures))

	for index, temp := range temperatures {
		for len(stack) > 0 && stack.top().data < temp {
			node := stack.pop()

			ans[node.index] = index - node.index
		}

		node := Node{data: temp, index: index}
		stack.push(node)
	}

	for len(stack) > 0 {
		node := stack.pop()
		ans[node.index] = 0
	}

	return ans
}

func (s *Stack) push(data Node) {
	*s = append(*s, data)
}

func (s *Stack) pop() Node {
	size := len(*s)
	elem := (*s)[size-1]

	*s = (*s)[:size-1]
	return elem
}

func (s *Stack) top() Node {
	size := len(*s)
	elem := (*s)[size-1]

	return elem
}
