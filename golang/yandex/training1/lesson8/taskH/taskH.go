package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var isAvl = true

// https://contest.yandex.ru/contest/28069/problems/H/
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

	findHeight(&tree)

	if isAvl {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func findHeight(tree *[]interface{}) int {
	if (*tree)[1] == nil && (*tree)[2] == nil {
		return 1
	}
	left, leftOk := (*tree)[1].([]interface{})
	right, rightOk := (*tree)[2].([]interface{})

	var leftHeight, rightHeight int

	if leftOk {
		leftHeight = findHeight(&left)
	}
	if rightOk {
		rightHeight = findHeight(&right)
	}

	if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
		isAvl = false
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
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
