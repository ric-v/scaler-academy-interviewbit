package besttimetobuyandsellstocks1

// BestTimeToBuyAndSellStocks returns the maximum profit that can be made by buying and selling a stock once.
// The function assumes that there are at least 2 days.
// The function assumes that there is always a way to make a profit.
// The function assumes that the price of the stock on each day is an integer.
// The function assumes that the price of the stock on each day is non-negative.
//
// Time complexity: O(n)
// Space complexity: O(1)
func BestTimeToBuyAndSellStocks(A []int) int {

	// The function assumes that there are at least 2 days.
	if len(A) < 2 {
		return 0
	}

	// max profit so far and min price so far
	maxProfit := 0
	minPrice := A[0]

	// iterate over the array and update the max profit and min price
	for _, v := range A[1:] {

		// if the current price is less than the min price, update the min price
		// else, update the max profit
		if v < minPrice {
			minPrice = v
		} else {

			// calculate the profit and update the max profit
			profit := v - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
