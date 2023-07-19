package evensubarrays

// EvenSubarrays returns "YES" if there is an even number of even numbers in the array
// and "NO" otherwise.
//
// Time complexity: O(n)
// Space complexity: O(1)
func EvenSubarrays(A []int) string {

	// If the first and last elements are even and
	// If the length of the array is even, then the subarrays are even
	if len(A)%2 == 0 && A[0]%2 == 0 && A[len(A)-1]%2 == 0 {
		return "YES"
	}
	return "NO"
}
