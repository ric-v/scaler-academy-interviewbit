package sumofallsubarrays

// SumOfAllSubarrays returns the sum of all subarrays of an array
// A = [1, 2, 3] => 1 + 2 + 3 + 1 + 2 + 3 + 1 + 2 + 3 = 18
//
// i+1 * (n-i) = number of subarrays starting at i and ending at j
// A[i] = value of the subarray
// hence, sum of all subarrays = (i+1) * (n-i) * A[i] for all i in [0, n-1]
// result = result + (i+1) * (n-i) * A[i] for all i in [0, n-1]
//
// Time complexity: O(n)
// Space complexity: O(1)
func SumOfAllSubarrays(A []int) int64 {
	var result int64

	// i+1 * (n-i) = number of subarrays starting at i and ending at j
	for i := 0; i < len(A); i++ {
		result += int64(((i + 1) * (len(A) - i)) * A[i])
	}
	return result
}

// SumOfAllSubarrays_CarryForward returns the sum of all subarrays of an array
// A = [1, 2, 3] => 1 + 1+2 + 1+2+3 + 2 + 2+3 + 3 = 18
//
// # Carry forward the sum of the previous subarray to the next subarray
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func SumOfAllSubarrays_CarryForward(A []int) int64 {
	var result int64
 
	for i := 0; i < len(A); i++ {
		var sum int64
		for j := i; j < len(A); j++ {
			sum += int64(A[j])
			result += sum
		}
	}
	return result
}

// SumOfAllSubarrays_PrefixSum returns the sum of all subarrays of an array
// A = [1, 2, 3] => 1 + 2 + 3 + 1 + 2 + 3 + 1 + 2 + 3 = 18
//
// prefixSum = [1, 3, 6] => prefixSum[j] = sum of all elements from A[0] to A[j]
// prefixSum[i-1] = sum of all elements from A[0] to A[i-1]
// prefixSum[j] - prefixSum[i-1] = sum of all elements from A[i] to A[j]
// hence, sum of all subarrays = prefixSum[j] - prefixSum[i-1] for all i in [0, n-1]
// result = result + prefixSum[j] - prefixSum[i-1] for all i in [0, n-1]
//
// Time complexity: O(n^2)
// Space complexity: O(n)
func SumOfAllSubarrays_PrefixSum(A []int) int64 {
	var result int64

	// prefixSum = [1, 3, 6] => prefixSum[j] = sum of all elements from A[0] to A[j]
	var prefixSum []int = []int{A[0]}
	for i := 1; i < len(A); i++ {
		prefixSum = append(prefixSum, A[i]+prefixSum[i-1])
	}

	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {

			// prefixSum[j] - prefixSum[i-1] = sum of all elements from A[i] to A[j]
			if i == 0 {
				result += int64(prefixSum[j])
			} else {
				result += int64(prefixSum[j] - prefixSum[i-1])
			}
		}
	}
	return result
}

// SumOfAllSubarrays_BruteForce returns the sum of all subarrays of an array
// A = [1, 2, 3] => 1 + 2 + 3 + 1 + 2 + 3 + 1 + 2 + 3 = 18
//
// # Brute force approach to find the sum of all subarrays of an array
//
// Time complexity: O(n^3)
// Space complexity: O(1)
func SumOfAllSubarrays_BruteForce(A []int) int64 {
	var result int64

	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			var sum int
			for k := i; k <= j; k++ {
				sum += A[k]
			}
			result += int64(sum)
		}
	}
	return result
}
