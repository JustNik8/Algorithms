package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pickedNumber = 10

func guess(num int) int {
	if num == pickedNumber {
		return 0
	} else if num > pickedNumber {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	l, r := 1, n

	for l <= r {
		m := (l + r) / 2

		res := guess(m)
		if res == 0 {
			return m
		}
		if res == -1 {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return 0
}
