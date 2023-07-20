package main

func main() {

}

// Мое решение
func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		if nums[m] < nums[r] {
			r = m
			break
		} else {
			l = m + 1
		}
	}

	minElem := nums[r]
	for l <= r {
		m := (l + r) / 2

		if nums[m] < nums[r] {
			minElem = nums[m]
			r = m
		} else {
			l = m + 1
		}
	}

	return minElem
}

// Neetcode решение
func findMinNeetcode(nums []int) int {
	l, r := 0, len(nums)-1

	minElem := nums[l]
	for l <= r {
		if nums[l] < nums[r] {
			minElem = min(minElem, nums[l])
			break
		}

		m := (l + r) / 2
		minElem = min(minElem, nums[m])

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}

	}

	return minElem
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
