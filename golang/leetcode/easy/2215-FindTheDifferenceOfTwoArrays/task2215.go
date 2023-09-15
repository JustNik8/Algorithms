package main

func main() {

}

type empty struct{}
type set map[int]empty

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(set)
	set2 := make(set)

	for i := 0; i < len(nums1); i++ {
		set1[nums1[i]] = empty{}
	}
	for i := 0; i < len(nums2); i++ {
		set2[nums2[i]] = empty{}
	}

	ans1 := make([]int, 0)
	for num, _ := range set1 {
		_, ok := set2[num]
		if !ok {
			ans1 = append(ans1, num)
		}
	}

	ans2 := make([]int, 0)
	for num, _ := range set2 {
		_, ok := set1[num]
		if !ok {
			ans2 = append(ans2, num)
		}
	}

	return [][]int{ans1, ans2}
}
