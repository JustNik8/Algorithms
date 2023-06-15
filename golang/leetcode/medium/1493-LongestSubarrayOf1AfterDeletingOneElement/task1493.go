package main

func main() {

}

func longestSubarray(nums []int) int {
	left, right := -1, -1
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i] == 1 {
			left, right = i, i
			break
		}
	}

	if left == -1 && right == -1 {
		return 0
	}

	maxLen := 0
	canDelete := true

	for right < size && left < size {
		if nums[right] == 1 {
			right++
		} else if canDelete {
			canDelete = false
			right++
		} else {
			currLen := right - left - 1
			if currLen > maxLen {
				maxLen = currLen
			}

			if nums[left] == 0 {
				canDelete = true
			}
			left++
		}

	}

	currLen := right - left - 1
	if canDelete && currLen != len(nums)-1 {
		currLen++
	}

	if currLen > maxLen {
		maxLen = currLen
	}

	return maxLen
}
