package main

func main() {

}

func maxArea(height []int) int {
	size := len(height)

	maxArea := 0
	left, right := 0, size-1

	for left < right {
		minHeight := min(height[left], height[right])

		area := minHeight * (right - left)

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else if height[right] < height[left] {
			right--
		} else {
			rightNeig := height[right-1]
			leftNeig := height[left+1]

			if leftNeig > rightNeig {
				left++
			} else {
				right--
			}
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
