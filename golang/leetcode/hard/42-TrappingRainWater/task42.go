package main

func main() {

}

func trap(height []int) int {
	size := len(height)

	left, right := 0, size-1
	leftMax, rightMax := height[left], height[right]

	ans := 0
	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			ans += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			ans += rightMax - height[right]
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
