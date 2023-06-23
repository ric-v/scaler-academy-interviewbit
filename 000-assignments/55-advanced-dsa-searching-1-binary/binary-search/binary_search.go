package binarysearch

// BinarySearch returns the element and index of the query in the array
// If the query is not found, the element is -1 and the index is -1
//
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(arr []int, query int) (element int, index int) {
	lo := 0
	hi := len(arr) - 1

	// Binary Search Loop Invariant:
	// 1. lo <= hi
	// 2. arr[0..lo-1] < query
	// 3. arr[hi+1..n-1] > query
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] == query {
			return arr[mid], mid
		} else if arr[mid] < query {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	// query not found
	return query, -1
}
