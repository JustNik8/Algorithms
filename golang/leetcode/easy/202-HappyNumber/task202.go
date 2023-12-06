package main

type empty struct{}

type set map[int]empty

// Я так понимаю это
func isHappy(n int) bool {
	used := make(set)

	sum := 0

	for {
		for n > 0 {
			part := n % 10
			n /= 10
			sum += part * part
		}

		if sum == 1 {
			return true
		}

		if _, isUsed := used[sum]; isUsed {
			return false
		}

		used[sum] = empty{}

		n = sum
		sum = 0
	}
}
