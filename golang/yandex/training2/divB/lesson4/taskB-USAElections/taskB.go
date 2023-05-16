package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/28970/problems/B/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	presidents := make(map[string]int)
	for scanner.Scan() {
		str := scanner.Text()
		args := strings.Split(str, " ")

		name := args[0]
		cnt, _ := strconv.Atoi(args[1])

		presidents[name] += cnt
	}

	sortedKeys := make([]string, 0)

	for president := range presidents {
		sortedKeys = append(sortedKeys, president)
	}

	sort.Strings(sortedKeys)
	for _, key := range sortedKeys {
		fmt.Printf("%s %d\n", key, presidents[key])
	}
}
