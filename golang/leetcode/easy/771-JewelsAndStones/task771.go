package main

type empty struct{}

type set map[byte]empty

func numJewelsInStones(jewels string, stones string) int {
	m := make(set)

	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = empty{}
	}

	cnt := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			cnt++
		}
	}

	return cnt
}
