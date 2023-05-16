package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28069/problems/F/
func main() {
	reader := bufio.NewReader(os.Stdin)

	tree := make([]interface{}, 0)
	var num int
	fmt.Fscan(reader, &num)
	initTree(&tree, num)

	for {
		fmt.Fscan(reader, &num)
		if num == 0 {
			break
		}
		add(&tree, num)
	}

	leftSearch(&tree)
}

func leftSearch(tree *[]interface{}) {
	root := *tree

	left, leftOk := root[1].([]interface{})
	right, rightOk := root[2].([]interface{})

	if leftOk {
		leftSearch(&left)
	}

	if leftOk && rightOk {
		fmt.Println(root[0])
	}

	if rightOk {
		leftSearch(&right)
	}

}

func add(tree *[]interface{}, x int) {
	root := (*tree)[0].(int)
	if x == root {
		return
	}
	if x < root {
		if (*tree)[1] == nil {
			subTree := make([]interface{}, 0)
			initTree(&subTree, x)
			(*tree)[1] = subTree
		} else {
			slice := (*tree)[1].([]interface{})
			add(&slice, x)
		}
	} else {
		if (*tree)[2] == nil {
			subTree := make([]interface{}, 0)
			initTree(&subTree, x)
			(*tree)[2] = subTree
		} else {
			slice := (*tree)[2].([]interface{})
			add(&slice, x)
		}
	}

}

func initTree(tree *[]interface{}, root int) {
	*tree = append(*tree, root)
	*tree = append(*tree, nil)
	*tree = append(*tree, nil)
}
