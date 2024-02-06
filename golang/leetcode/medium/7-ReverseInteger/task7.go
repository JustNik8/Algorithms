package main

func reverse(x int) int {
	prev := int32(0)
	ans := int32(0)

	for x != 0 {
		digit := int32(x % 10)
		x /= 10

		ans *= 10
		ans += digit

		if ans/10 != prev {
			return 0
		}

		prev = ans
	}

	return int(ans)
}
