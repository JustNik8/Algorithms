package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	numsCnt := make(map[int]int)
	bucket := make([][]int, len(nums)+1)

	for _, num := range nums {
		numsCnt[num] += 1
	}

	for num, cnt := range numsCnt {
		bucket[cnt] = append(bucket[cnt], num)
	}

	ans := make([]int, k)
	cnt := 0

	for i := len(bucket) - 1; i >= 0; i-- {
		for j := 0; j < len(bucket[i]); j++ {
			ans[cnt] = bucket[i][j]
			cnt++

			if cnt == k {
				return ans
			}
		}
	}

	return ans
}
