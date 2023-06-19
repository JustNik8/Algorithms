package main

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		var cnt [26]int

		str := strs[i]

		for j := 0; j < len(str); j++ {
			cnt[str[j]-'a']++
		}

		_, exists := groups[cnt]

		if !exists {
			groups[cnt] = make([]string, 0)
		}
		groups[cnt] = append(groups[cnt], str)
	}

	ans := make([][]string, len(groups))
	idx := 0
	for _, v := range groups {
		ans[idx] = v
		idx++
	}

	return ans
}
