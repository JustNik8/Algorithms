package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	size := len(numbers)

	left, right := 0, size-1

	for left < right {
		leftNum := numbers[left]
		rightNum := numbers[right]

		if leftNum+rightNum == target {
			return []int{left + 1, right + 1}
		} else if leftNum+rightNum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
