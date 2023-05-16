package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type time struct {
	hours   int
	minutes int
	seconds int
}

// https://contest.yandex.ru/contest/45468/problems/7/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)

	aStamp := parseTime(a)
	cStamp := parseTime(c)

	if cStamp < aStamp {
		cStamp += 24 * 60 * 60
	}

	diff := int(math.Round((float64(cStamp) - float64(aStamp)) / 2))

	bStamp := parseTime(b) + diff

	fmt.Println(parseStamp(bStamp))
}

func parseTime(str string) int {
	timeStr := strings.Split(str, ":")
	hours, _ := strconv.Atoi(timeStr[0])
	minutes, _ := strconv.Atoi(timeStr[1])
	seconds, _ := strconv.Atoi(timeStr[2])

	return seconds + minutes*60 + hours*60*60
}

func parseStamp(stamp int) string {
	if stamp >= 24*60*60 {
		stamp -= 24 * 60 * 60
	}
	hours := stamp / (60 * 60)
	stamp -= hours * 60 * 60

	minutes := stamp / 60
	stamp -= minutes * 60

	seconds := stamp

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

/*
Tc 2
11:37:00
23:51:00
23:59:59

Output
06:02:30

Tc 9
21:00:00
22:00:00
06:00:00

Output
02:30:00
*/
