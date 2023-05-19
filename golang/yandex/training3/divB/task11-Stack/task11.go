package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stack struct {
	size  int
	elems []int
}

// https://contest.yandex.ru/contest/45468/problems/11/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	s := stack{}

	for scanner.Scan() {
		cmd := scanner.Text()
		switch cmd {
		case "push":
			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			fmt.Println(s.push(num))
		case "pop":
			fmt.Println(s.pop())
		case "back":
			fmt.Println(s.back())
		case "size":
			fmt.Println(s.size)
		case "clear":
			fmt.Println(s.clear())
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}

func (s *stack) push(n int) string {
	s.elems = append(s.elems, n)
	s.size++
	return "ok"
}

func (s *stack) pop() string {
	if s.size <= 0 {
		return "error"
	}
	elem := s.elems[s.size-1]
	s.elems = s.elems[:s.size-1]
	s.size--
	return strconv.Itoa(elem)
}

func (s *stack) back() string {
	if s.size <= 0 {
		return "error"
	}
	elem := s.elems[s.size-1]
	return strconv.Itoa(elem)
}

func (s *stack) clear() string {
	s.elems = s.elems[:0]
	s.size = 0
	return "ok"
}
