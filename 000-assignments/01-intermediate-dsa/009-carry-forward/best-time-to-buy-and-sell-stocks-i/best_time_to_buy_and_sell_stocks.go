package besttimetobuyandsellstocksi

// BestTimeToBuyAndSellStocks returns the maximum profit that can be made by buying and selling a stock
// once. The input array represents the price of the stock on each day.
//
// Time complexity: O(n)
// Space complexity: O(1)
func BestTimeToBuyAndSellStocks(A []int) int {
	// If the array is empty, return 0
	if len(A) == 0 {
		return 0
	}

	// Initialize the minimum price so far to the first element of the array
	// and the maximum profit to 0
	minSoFar := A[0]
	maxProfit := 0

	// Iterate over the array
	for i := 1; i < len(A); i++ {

		// If the current element is less than the minimum price so far, update the minimum price so far
		// Else if the current element minus the minimum price so far is greater than the maximum profit,
		// update the maximum profit
		if A[i] < minSoFar {
			minSoFar = A[i]
		} else if A[i]-minSoFar > maxProfit {
			maxProfit = A[i] - minSoFar
		}
	}
	return maxProfit
}
