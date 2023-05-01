package firstrepeatingelement

// FirstRepeatingElement returns the index of the first repeating element in the array.
// If there are no repeating elements, it returns -1.
//
// Time complexity: O(n)
// Space complexity: O(n)
func FirstRepeatingElement(A []int) int {

	// The idea is to store the index of the first occurrence of every element in a map.
	smallindex := 1<<31 - 1
	occMap := make(map[int]int)

	// If the element is already present in the map, check if the index is smaller than the current smallest index.
	for i, a := range A {
		if val, ok := occMap[a]; ok {
			if val < smallindex {
				// If yes, update the smallest index.
				smallindex = val
			}
		} else {
			// If the element is not present in the map, add it.
			occMap[a] = i
		}
	}

	// If the map size is equal to the array size, it means there are no repeating elements.
	if len(occMap) == len(A) {
		return -1
	}
	return A[smallindex]
}
