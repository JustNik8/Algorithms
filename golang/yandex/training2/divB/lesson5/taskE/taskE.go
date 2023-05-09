package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/29075/problems/E/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var s int
	fmt.Fscan(reader, &s)

	a := readArr(reader)
	b := readArr(reader)
	c := readArr(reader)

	nums := make(map[int]int)

	for k := 0; k < len(c); k++ {
		num := c[k]

		_, found := nums[num]
		if !found {
			nums[num] = k
		}
	}

	ans := [3]int{-1, -1, -1}
	isFound := false
	for i := 0; i < len(a); i++ {
		if isFound {
			break
		}
		for j := 0; j < len(b); j++ {
			residue := s - (a[i] + b[j])
			if residue <= 0 {
				continue
			}

			k, found := nums[residue]
			if found {
				ans[0] = i
				ans[1] = j
				ans[2] = k
				isFound = true
				break
			}
		}
	}

	if isFound {
		fmt.Println(ans[0], ans[1], ans[2])
	} else {
		fmt.Println(-1)
	}

}

func readArr(reader *bufio.Reader) []int {
	var size int
	fmt.Fscan(reader, &size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		var num int
		fmt.Fscan(reader, &num)
		arr[i] = num
	}

	return arr
}
