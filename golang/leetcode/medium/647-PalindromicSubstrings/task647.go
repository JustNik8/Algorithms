package main

func main() {

}

func countSubstrings(s string) int {
	size := len(s)
	cnt := 0

	for i := 0; i < size; i++ {
		left, right := i, i
		cnt += calculate(s, left, right)

		left, right = i, i+1
		cnt += calculate(s, left, right)
	}

	return cnt
}

func calculate(s string, left, right int) int {
	size := len(s)
	cnt := 0
	for left >= 0 && right < size {
		if s[left] != s[right] {
			break
		}

		cnt++
		left--
		right++
	}

	return cnt
}
