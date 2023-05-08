package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type topic struct {
	name string
	cnt  int
}

// https://contest.yandex.ru/contest/28970/problems/E/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	topics := make(map[int]topic)
	messages := make(map[int]int)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		number, _ := strconv.Atoi(scanner.Text())

		if number == 0 {
			scanner.Scan()
			topicName := scanner.Text()

			t := topic{name: topicName, cnt: 1}
			topics[i+1] = t
		} else {
			var topicNum int
			_, found := topics[number]
			if found {
				topicNum = number
			} else {
				topicNum = messages[number]
			}
			t := topics[topicNum]
			t.cnt += 1
			topics[topicNum] = t
			messages[i+1] = topicNum

		}

		scanner.Scan()

	}

	maxTopic := topic{name: "", cnt: 0}
	topicPos := -1

	for pos, t := range topics {
		if t.cnt > maxTopic.cnt {
			maxTopic = t
			topicPos = pos
		} else if t.cnt == maxTopic.cnt && (pos < topicPos) {
			maxTopic = t
			topicPos = pos
		}
	}

	fmt.Println(maxTopic.name)
}
