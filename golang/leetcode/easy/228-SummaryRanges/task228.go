package main

import "fmt"

func main() {

}

func summaryRanges(nums []int) []string {
	size := len(nums)
	if size == 0 {
		return []string{}
	}
	if size == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	left := 0
	ans := make([]string, 0)

	for i := 1; i < size; i++ {
		if nums[i]-nums[i-1] == 1 {
			continue
		} else if i-left > 1 {
			numsRange := fmt.Sprintf("%d->%d", nums[left], nums[i-1])
			ans = append(ans, numsRange)
			left = i
		} else {
			ans = append(ans, fmt.Sprintf("%d", nums[i-1]))
			left = i
		}
	}

	last := size - 1
	if nums[last]-nums[last-1] == 1 {
		numsRange := fmt.Sprintf("%d->%d", nums[left], nums[last])
		ans = append(ans, numsRange)
	} else {
		ans = append(ans, fmt.Sprintf("%d", nums[last]))
	}

	return ans
}
