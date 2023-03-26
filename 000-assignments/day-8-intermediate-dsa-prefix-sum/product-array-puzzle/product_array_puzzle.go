package productarraypuzzle

// ProductArrayPuzzle returns an array of products of all elements except the element at that index
//
// Time complexity: O(n)
// Space complexity: O(1)
func ProductArrayPuzzle(A []int) []int {

	// 1. Calculate the product of all elements in the array
	var arrProduct int = 1
	for _, num := range A {
		arrProduct *= num
	}

	// 2. Divide the product by each element in the array
	//   to get the product of all elements except the element at that index
	//  and store it in the same array
	for i, num := range A {
		A[i] = arrProduct / num
	}
	return A
}
