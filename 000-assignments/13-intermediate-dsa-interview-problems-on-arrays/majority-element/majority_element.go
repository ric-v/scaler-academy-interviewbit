package majorityelement

// MajorityElement returns the majority element in an array.
// The majority element is the element that appears more than n/2 times.
//
// The algorithm is based on the Boyer-Moore majority vote algorithm.
//
// Time complexity: O(n)
// Space complexity: O(1)
func MajorityElement(A []int) int {
	var (
		major = A[0] // The majority element is the first element at the beginning.
		freq  = 1    // The frequency of the majority element is 1 at the beginning.
	)

	// The majority element is the element that appears more than n/2 times.
	for _, n := range A[1:] {

		// If the current element is the same as the majority element,
		if n == major {
			// increase the frequency of the majority element.
			freq++
		} else {

			// If the current element is not the same as the majority element,
			// decrease the frequency of the majority element.
			if freq > 0 {
				freq--
			} else {

				// If the frequency of the majority element is 0,
				// the current element becomes the majority element.
				major = n
				freq = 1
			}
		}
	}

	// Check if the majority element appears more than n/2 times.
	freq = 0
	for _, n := range A {
		if n == major {
			freq++
		}
	}

	// If the majority element appears more than n/2 times, return it.
	// Otherwise, return -1.
	if freq <= len(A)/2 {
		return -1
	}
	return major
}
