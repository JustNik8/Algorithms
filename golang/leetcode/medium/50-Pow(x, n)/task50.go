package main

func myPow(x float64, n int) float64 {
	ans := calc(x, abs(n))

	if n >= 0 {
		return ans
	} else {
		return 1 / ans
	}
}

func calc(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}

	res := calc(x*x, n/2)

	if n%2 == 0 {
		return res
	} else {
		return res * x
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
