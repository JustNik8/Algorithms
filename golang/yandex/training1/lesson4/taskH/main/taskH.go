package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	dict := make(map[string]map[string]bool)
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)
		key := strings.ToLower(word)
		_, ok := dict[key]
		if !ok {
			dict[key] = make(map[string]bool)
		}
		dict[key][word] = true
	}
	in.ReadString('\n')
	taskString, _ := in.ReadString('\n')
	taskString = taskString[:len(taskString)-1]
	if len(strings.TrimSpace(taskString)) == 0 {
		fmt.Println(0)
		return
	}
	taskWords := strings.Split(taskString, " ")

	mistakes := 0
	for _, word := range taskWords {
		key := strings.ToLower(word)
		upperCnt := 0
		for _, char := range word {
			if unicode.IsUpper(char) {
				upperCnt++
			}
		}
		if upperCnt != 1 {
			mistakes++
			continue
		}

		dictWords, hasWords := dict[key]
		if !hasWords {
			continue
		}
		_, isRight := dictWords[word]
		if !isRight {
			mistakes++
		}
	}

	fmt.Print(mistakes)
}
