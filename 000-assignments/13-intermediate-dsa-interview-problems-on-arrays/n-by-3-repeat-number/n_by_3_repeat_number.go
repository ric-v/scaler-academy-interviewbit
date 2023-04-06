package nby3repeatnumber

// MajorityElement returns the majority element in an array.
// The majority element is the element that appears more than n/3 times.
// If there is no majority element, return -1.
//
// The algorithm is based on the Boyer-Moore majority vote algorithm.
//
// Time complexity: O(n)
// Space complexity: O(1)
func MajorityElement(nums []int) int {

	// Find the two majority elements. If there are no majority elements, return -1.
	// The two majority elements can be the same element since the majority element appears more than n/3 times.
	var major1, major2, freq1, freq2 = nums[0], 0, 1, 0

	for _, n := range nums[1:] {

		// If the current element is equal to the first majority element,
		if n == major1 {
			// increment the frequency of the first majority element.
			freq1++
		} else if n == major2 {
			// If the current element is equal to the second majority element,
			// increment the frequency of the second majority element.
			freq2++
		} else if freq1 == 0 {

			// If the current element is not equal to the first or second majority element,
			// and the frequency of the first majority element is 0,
			major1 = n
			freq1 = 1

		} else if freq2 == 0 {

			// If the current element is not equal to the first or second majority element,
			// and the frequency of the second majority element is 0,
			major2 = n
			freq2 = 1
		} else {

			// If the current element is not equal to the first or second majority element,
			// and the frequency of the second majority element is not 0,
			freq1--
			freq2--
		}
	}

	// Check if the majority element appears more than n/3 times.
	freq1, freq2 = 0, 0
	for _, n := range nums {
		if n == major1 {
			freq1++
		}
		if n == major2 {
			freq2++
		}
	}

	// If the majority element appears more than n/3 times, return the majority element.
	// Otherwise, return -1.
	if freq1 > len(nums)/3 {
		return major1
	}
	if freq2 > len(nums)/3 {
		return major2
	}
	return -1
}
