package main

func main() {

}

// https://leetcode.com/problems/two-sum/?envType=list&envId=oceyii8d
func twoSum(nums []int, target int) []int {
	indices := make(map[int]int)

	for index, num := range nums {
		diff := target - num

		another, ok := indices[diff]
		if ok {
			return []int{index, another}
		}

		indices[num] = index
	}

	return []int{0, 0}
}
