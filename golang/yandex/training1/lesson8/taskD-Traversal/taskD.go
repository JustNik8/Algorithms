package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28069/problems/D/
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
	if tree == nil {
		return
	}

	left, ok := (*tree)[1].([]interface{})
	if ok {
		leftSearch(&left)
	}

	fmt.Println((*tree)[0])

	right, ok := (*tree)[2].([]interface{})
	if ok {
		leftSearch(&right)
	}
}

func add(tree *[]interface{}, x int) {
	root := (*tree)[0].(int)
	if x == root {
		return
	} else if x < root {
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
