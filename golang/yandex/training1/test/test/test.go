package main

import "fmt"

func main() {
	m := 31623000000000 - 1
	s := 1000000000

	ms := m / s
	fmt.Println(ms*ms >= s)

}
