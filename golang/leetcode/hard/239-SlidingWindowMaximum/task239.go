package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	// store indexes
	deque := make([]int, 0)

	ans := make([]int, 0)
	left, right := 0, 0

	for right < len(nums) {
		elem := nums[right]
		for len(deque) > 0 && nums[deque[len(deque)-1]] < elem {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)

		if deque[0] < left {
			deque = deque[1:]
		}

		if left+right+1 >= k {
			ans = append(ans, nums[deque[0]])
			left++
		}
		right++
	}

	return ans
}
