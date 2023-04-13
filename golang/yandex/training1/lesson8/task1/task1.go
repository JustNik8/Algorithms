package main

import (
	"bufio"
	"fmt"
	"os"
)

var maxHeight = 0

func main() {
	reader := bufio.NewReader(os.Stdin)

	tree := make([]interface{}, 0)

	var num int
	fmt.Fscan(reader, &num)

	if num == 0 {
		fmt.Println(0)
		return
	}

	initTree(&tree, num)
	maxHeight = 1

	for {
		fmt.Fscan(reader, &num)
		if num == 0 {
			break
		}
		add(&tree, num, 1)
	}

	fmt.Println(maxHeight)
}

func add(tree *[]interface{}, x int, currentHeight int) {
	root := (*tree)[0].(int)
	if x == root {
		return
	} else if x < root {
		if (*tree)[1] == nil {
			subTree := make([]interface{}, 0)
			initTree(&subTree, x)
			(*tree)[1] = subTree
			checkHeight(currentHeight + 1)
		} else {
			slice := (*tree)[1].([]interface{})
			add(&slice, x, currentHeight+1)
		}
	} else {
		if (*tree)[2] == nil {
			subTree := make([]interface{}, 0)
			initTree(&subTree, x)
			(*tree)[2] = subTree
			checkHeight(currentHeight + 1)
		} else {
			slice := (*tree)[2].([]interface{})
			add(&slice, x, currentHeight+1)
		}
	}
}

func initTree(tree *[]interface{}, root int) {
	*tree = append(*tree, root)
	*tree = append(*tree, nil)
	*tree = append(*tree, nil)
}

func checkHeight(currentHeight int) {
	if currentHeight > maxHeight {
		maxHeight = currentHeight
	}
}
