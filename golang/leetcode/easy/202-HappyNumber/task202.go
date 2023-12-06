package main

type empty struct{}

type set map[int]empty

// Решение через set
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

// Решение через поиска цикла в связанном цикле с помощью slow и fast
func isHappy2(n int) bool {
	slow, fast := n, squareDigits(n)

	for slow != fast {
		fast = squareDigits(fast)
		fast = squareDigits(fast)
		slow = squareDigits(slow)
	}

	return fast == 1
}

func squareDigits(n int) int {
	sum := 0
	for n > 0 {
		part := n % 10
		sum += part * part
		n /= 10
	}

	return sum
}
