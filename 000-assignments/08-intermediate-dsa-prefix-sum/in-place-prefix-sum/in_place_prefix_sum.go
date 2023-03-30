package inplaceprefixsum

// InPlacePrefixSum calculates the prefix sum of an array in place.
// The prefix sum of an array is the sum of all elements in the array
// up to and including the current element.
//
// A[i] = A[i-1] + A[i]
//
// eg: A = [1, 2, 3, 4, 5]
//
// A[0] = A[0] = 1
// A[1] = A[0] + A[1] = 1 + 2 = 3
// A[2] = A[1] + A[2] = 3 + 3 = 6
// A[3] = A[2] + A[3] = 6 + 4 = 10
// A[4] = A[3] + A[4] = 10 + 5 = 15
//
// A = [1, 3, 6, 10, 15]
//
// Time complexity: O(n)
// Space complexity: O(1)
func InPlacePrefixSum(A []int) []int {

	for i := 1; i < len(A); i++ {
		A[i] = A[i-1] + A[i]
	}
	return A
}

// FilterPrefixSum calculates the prefix sum of an array in place.
// The prefix sum of an array is the sum of all elements in the array
// up to and including the current element.
// The filter function is applied to each element in the array before
// calculating the prefix sum.
// The filter function should return 0 if the element should be ignored
// and 1 if the element should be included or any other value to be added to the sum.
//
// eg: A = [1, 2, 3, 4, 5]
// filter = func(x int) int { if x%2 == 0 { return 1 } else { return 0 } } // filter even numbers
//
// A[0] = filter(A[0]) = 0
// A[1] = A[0] + filter(A[1]) = 0 + 1 = 1
// A[2] = A[1] + filter(A[2]) = 1 + 0 = 1
// A[3] = A[2] + filter(A[3]) = 1 + 1 = 2
// A[4] = A[3] + filter(A[4]) = 2 + 0 = 2
//
// A = [0, 1, 1, 2, 2]
//
// Time complexity: O(n)
// Space complexity: O(1)
func FilterPrefixSum(A []int, filter func(x int) int) []int {
	for i := 0; i < len(A); i++ {
		if i == 0 {
			A[i] = filter(A[i])
		} else {
			A[i] = A[i-1] + filter(A[i])
		}
	}
	return A
}
