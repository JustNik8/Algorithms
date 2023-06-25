package main

func main() {

}

func productExceptSelf(nums []int) []int {
	size := len(nums)
	ans := make([]int, size)

	prefix := 1
	for i := 0; i < size; i++ {
		ans[i] = prefix
		prefix = prefix * nums[i]
	}

	postfix := 1
	for i := size - 1; i >= 0; i-- {
		ans[i] = ans[i] * postfix
		postfix = postfix * nums[i]
	}

	return ans
}
