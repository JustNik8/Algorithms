package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type deque struct {
	elems []int
}

// https://contest.yandex.ru/contest/45468/problems/18/
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	d := deque{}

	var cmd string
	for sc.Scan() {
		cmd = sc.Text()
		switch cmd {
		case "push_front":
			sc.Scan()
			num, _ := strconv.Atoi(sc.Text())
			fmt.Println(d.pushFront(num))
		case "push_back":
			sc.Scan()
			num, _ := strconv.Atoi(sc.Text())
			fmt.Println(d.pushBack(num))
		case "pop_front":
			fmt.Println(d.popFront())
		case "pop_back":
			fmt.Println(d.popBack())
		case "front":
			fmt.Println(d.front())
		case "back":
			fmt.Println(d.back())
		case "size":
			fmt.Println(d.size())
		case "clear":
			fmt.Println(d.clear())
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}

func (d *deque) pushFront(n int) string {
	d.elems = append(d.elems, 0)
	copy(d.elems[1:], d.elems)
	d.elems[0] = n
	return "ok"
}

func (d *deque) pushBack(n int) string {
	d.elems = append(d.elems, n)
	return "ok"
}

func (d *deque) popFront() string {
	if len(d.elems) == 0 {
		return "error"
	}
	elem := (d.elems)[0]
	d.elems = (d.elems)[1:]
	return strconv.Itoa(elem)
}

func (d *deque) popBack() string {
	size := len(d.elems)
	if size == 0 {
		return "error"
	}
	elem := d.elems[size-1]
	d.elems = d.elems[:size-1]

	return strconv.Itoa(elem)
}

func (d *deque) front() string {
	if len(d.elems) == 0 {
		return "error"
	}
	elem := d.elems[0]
	return strconv.Itoa(elem)
}

func (d *deque) back() string {
	size := len(d.elems)
	if size == 0 {
		return "error"
	}
	elem := d.elems[size-1]
	return strconv.Itoa(elem)
}

func (d *deque) size() int {
	return len(d.elems)
}

func (d *deque) clear() string {
	d.elems = d.elems[:0]
	return "ok"
}
