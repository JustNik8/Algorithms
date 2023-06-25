package main

func main() {

}

func longestConsecutive(nums []int) int {
	marked := make(map[int]bool)

	for _, num := range nums {
		marked[num] = false
	}

	maxLen := 0
	for num, _ := range marked {
		if marked[num] {
			continue
		}
		marked[num] = true

		currLen := 1
		currNum := num

		_, hasNext := marked[currNum+1]
		for hasNext {
			currNum = currNum + 1
			currLen += 1

			marked[currNum] = true

			_, hasNext = marked[currNum+1]
		}

		if currLen > maxLen {
			maxLen = currLen
		}

	}

	return maxLen
}
