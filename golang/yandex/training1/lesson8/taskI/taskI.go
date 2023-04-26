package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	value    string
	children []*node
}

var root *node = nil

// https://contest.yandex.ru/contest/28069/problems/I/
// Не завешил
func main() {
	reader := bufio.NewReader(os.Stdin)

	var cnt int
	fmt.Fscan(reader, &cnt)

	var child, parent string
	fmt.Fscan(reader, &child, &parent)

	root = initNode(parent, child)

	for i := 0; i < cnt-2; i++ {
		fmt.Fscan(reader, &child, &parent)
	}
}

func initNode(value string, children ...string) *node {
	var ch []*node = nil
	if len(children) > 0 {
		ch = make([]*node, len(children))
		for i, child := range children {
			ch[i] = &node{value: child}
		}
	}

	root := &node{value: value, children: ch}
	return root
}
