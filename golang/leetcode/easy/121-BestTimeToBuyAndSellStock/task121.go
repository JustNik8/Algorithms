package main

func main() {

}

func maxProfit(prices []int) int {
	bestProfit := 0
	minCost := prices[0]

	for i := 0; i < len(prices); i++ {
		profit := prices[i] - minCost

		if profit > bestProfit {
			bestProfit = profit
		}

		if prices[i] < minCost {
			minCost = prices[i]
		}
	}

	return bestProfit
}
