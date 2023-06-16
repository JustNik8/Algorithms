package main

func main() {

}

func moveZeroes(nums []int) {
	shift := 0

	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i] == 0 {
			shift++
		} else {
			nums[i-shift] = nums[i]
		}
	}

	for i := size - shift; i < size; i++ {
		nums[i] = 0
	}
}
