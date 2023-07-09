package main

func main() {

}

func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		m := (right + left) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			right = m
		} else {
			left = m + 1
		}
	}

	return -1
}
