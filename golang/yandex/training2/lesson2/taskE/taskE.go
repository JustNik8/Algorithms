package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var folderCnt int
	fmt.Fscan(reader, &folderCnt)

	folders := make([]int, folderCnt)
	for i := 0; i < folderCnt; i++ {
		var diploma int
		fmt.Fscan(reader, &diploma)
		folders[i] = diploma
	}

	maxDiplomasIndex := 0
	for i := 1; i < folderCnt; i++ {
		if folders[i] > folders[maxDiplomasIndex] {
			maxDiplomasIndex = i
		}
	}

	secCount := 0
	for i := 0; i < folderCnt; i++ {
		if i != maxDiplomasIndex {
			secCount += folders[i]
		}
	}

	fmt.Print(secCount)
}
