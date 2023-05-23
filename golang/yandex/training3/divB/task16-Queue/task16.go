package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type queue struct {
	elems []int
}

// https://contest.yandex.ru/contest/45468/problems/16/
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	q := queue{}

	var cmd string
	for {
		sc.Scan()
		cmd = sc.Text()
		switch cmd {
		case "push":
			sc.Scan()
			num, _ := strconv.Atoi(sc.Text())
			fmt.Println(q.push(num))
		case "pop":
			fmt.Println(q.pop())
		case "front":
			fmt.Println(q.front())
		case "size":
			fmt.Println(q.size())
		case "clear":
			fmt.Println(q.clear())
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}

func (q *queue) push(n int) string {
	q.elems = append(q.elems, n)
	return "ok"
}

func (q *queue) pop() string {
	if len(q.elems) == 0 {
		return "error"
	}
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return strconv.Itoa(elem)
}

func (q *queue) front() string {
	if len(q.elems) == 0 {
		return "error"
	}
	elem := q.elems[0]
	return strconv.Itoa(elem)
}

func (q *queue) size() int {
	return len(q.elems)
}

func (q *queue) clear() string {
	q.elems = q.elems[:0]
	return "ok"
}
