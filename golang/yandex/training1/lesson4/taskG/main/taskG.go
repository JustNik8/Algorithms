package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEPOSIT = "DEPOSIT"
const INCOME = "INCOME"
const WITHDRAW = "WITHDRAW"
const BALANCE = "BALANCE"
const TRANSFER = "TRANSFER"

func main() {
	accounts := make(map[string]int64)
	in := bufio.NewReader(os.Stdin)

	var sb strings.Builder

	for {
		var command string
		fmt.Fscan(in, &command)
		if len(command) == 0 {
			break
		}
		var name string
		if command != INCOME {
			fmt.Fscan(in, &name)
			if command != BALANCE {
				if _, ok := accounts[name]; !ok {
					accounts[name] = 0
				}
			}
		}

		switch command {
		case DEPOSIT:
			var amount int64
			fmt.Fscan(in, &amount)
			accounts[name] += amount
		case WITHDRAW:
			var amount int64
			fmt.Fscan(in, &amount)
			accounts[name] -= amount
		case BALANCE:
			balance, ok := accounts[name]
			if !ok {
				sb.WriteString("ERROR\n")
			} else {
				sb.WriteString(strconv.FormatInt(balance, 10))
				sb.WriteString("\n")
			}
		case TRANSFER:
			var recipient string
			var amount int64
			fmt.Fscan(in, &recipient, &amount)
			accounts[name] -= amount
			accounts[recipient] += amount
		case INCOME:
			var percent float64
			fmt.Fscan(in, &percent)
			for name, balance := range accounts {
				if balance > 0 {
					accounts[name] += int64(float64(balance) * (percent / 100))
				}
			}
		}
	}

	ans := sb.String()
	fmt.Print(ans[:len(ans)-1])
}
