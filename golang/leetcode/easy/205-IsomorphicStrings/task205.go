package main

func isIsomorphic(s string, t string) bool {
	size := len(s)
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < size; i++ {
		sChar, sExists := sMap[s[i]]
		tChar, tExists := tMap[t[i]]

		if !sExists && !tExists {
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
			continue
		}

		if sChar != t[i] && tChar != s[i] {
			return false
		}
	}

	return true
}
