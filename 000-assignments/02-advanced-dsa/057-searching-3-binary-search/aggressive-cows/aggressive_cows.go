package aggressivecows

import "sort"

// AggresssiveCowns returns the maximum distance between two cows if we place B cows
// in A such that the minimum distance between any two cows is maximum.
//
// Time complexity: O(nlogn)
// Space complexity: O(1)
func AggresssiveCowns(A []int, B int) int {
	sort.IntSlice(A).Sort()

	var low, high int = 0, -1
	ans := -1

	// Find the maximum distance between any two cows.
	for _, arr := range A {
		high = max(arr, high)
	}

	// Binary search for the maximum distance between any two cows.
	for low <= high {
		mid := (low + high) / 2

		// If it is possible to place cows in A such that the minimum distance between
		// any two cows is mid, then we try to find a greater value of mid.
		// Otherwise, we try to find a smaller value of mid.
		if isPossible(A, B, mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

// isPossible returns true if it is possible to place cows in A such that the minimum
// distance between any two cows is mid.
// Time complexity: O(n)
// Space complexity: O(1)
func isPossible(arr []int, cows int, mid int) bool {
	var lastPos, cowCount int = arr[0], 1

	// Place cows in A such that the minimum distance between any two cows is mid.
	for i := 0; i < len(arr); i++ {

		// If it is possible to place cows in A such that the minimum distance between
		// any two cows is mid, then we try to find a greater value of mid.
		// Otherwise, we try to find a smaller value of mid.
		if arr[i]-lastPos >= mid {
			cowCount++
			if cowCount == cows {
				return true
			}
			lastPos = arr[i]
		}
	}
	return false
}

// max returns the maximum of two integers.
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
