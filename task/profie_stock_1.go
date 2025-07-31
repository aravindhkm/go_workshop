package task

import "fmt"

func findMaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	prevPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {

		// check if loss
		if prevPrice > prices[i] {
			prevPrice = prices[i]
		} else {
			// find the profit
			currProfit := prices[i] - prevPrice

			if currProfit > maxProfit {
				maxProfit = currProfit
			}
		}
	}
	return maxProfit
}

func ProfitStockOne() {
	price := []int{6, 5, 4, 8, 6, 8, 7, 8, 9, 4, 5}

	fmt.Println(findMaxProfit(price))
}

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.
