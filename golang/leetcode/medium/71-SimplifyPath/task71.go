package main

import "strings"

func simplifyPath(path string) string {
	splitPath := strings.Split(path, "/")
	s := make([]string, 0)

	for _, part := range splitPath {
		if part == "." || part == "" {
			continue
		} else if part == ".." {
			if len(s) != 0 {
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, part)
		}
	}

	if len(s) == 0 {
		return "/"
	}

	var ans strings.Builder
	for _, part := range s {
		ans.WriteString("/")
		ans.WriteString(part)
	}

	return ans.String()
}
