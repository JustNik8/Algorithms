package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/27794/problems/B/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		fmt.Fscan(reader, &number)
		numbers[i] = number
	}

	var left = 0
	var right = 0
	var cnt = 0
	var nowSum = numbers[0]

	for right < n {
		if nowSum == k {
			cnt++
		}
		if left == right || nowSum < k {
			right++
			if right < n {
				nowSum += numbers[right]
			} else {
				break
			}
		} else {
			nowSum -= numbers[left]
			left++
		}
	}

	fmt.Print(cnt)
}
