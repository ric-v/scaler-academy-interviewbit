package reverseinarange

// ReverseInARange function reverses the elements of an array in a given range
//
// Time Complexity: O(n)
// Space Complexity: O(1)
//
// Example:
// [1, 2, 3, 4, 5, 6, 7, 8, 9] Start = 0, End = 8
// [9, 8, 7, 6, 5, 4, 3, 2, 1]
//
// [1, 2, 3, 4, 5, 6, 7, 8, 9] Start = 2, End = 6
// [1, 2, 7, 6, 5, 4, 3, 8, 9]
func ReverseInARange(Arr []int, Start int, End int) []int {

	// iterate over the array from Start to End until Start < End
	for Start < End {

		// swap the elements
		Arr[Start], Arr[End] = Arr[End], Arr[Start]
		Start++
		End--
	}
	return Arr
}
