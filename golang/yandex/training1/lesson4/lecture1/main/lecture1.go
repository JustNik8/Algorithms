package main

import "fmt"

func main() {

	const k = 10
	var firstDigits [k]int
	var secondDigits [k]int

	var first int
	fmt.Scan(&first)
	var second int
	fmt.Scan(&second)

	for first > 0 {
		ost := first % 10
		firstDigits[ost] += 1
		first /= 10
	}

	for second > 0 {
		ost := second % 10
		secondDigits[ost] += 1
		second /= 10
	}

	res := true
	for i := 0; i < 10; i++ {
		if firstDigits[i] != secondDigits[i] {
			res = false
			break
		}
	}

	fmt.Println(res)
}
