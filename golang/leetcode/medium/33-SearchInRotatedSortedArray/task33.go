package main

func main() {

}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1

	for l <= r {
		m := (l + r) / 2

		if nums[m] == target {
			ans = m
			return ans
		}

		if nums[m] < nums[r] {
			if nums[m] <= target && nums[r] >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[l] <= target && nums[m] >= target {
				r = m - 1
			} else {
				l = m + 1
			}
		}

	}

	return ans
}

// one else solution
func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			// Left portion
			if target < nums[l] || target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			// Right portion
			if target > nums[r] || target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
