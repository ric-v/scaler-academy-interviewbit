package longestconsecutivesequence

func LongestConsecutiveSequence(A []int) int {
	var longestChain int = 0
	hs := make(map[int]bool)

	// Add all elements to a hash set
	for _, v := range A {
		hs[v] = true
	}

	// Check if the element is the start of a chain
	for _, v := range A {

		// If the element is the start of a chain, check the length of the chain
		// by checking if the next element is in the hash set
		var x int = v
		if _, ok := hs[x-1]; !ok {

			// If the next element is not in the hash set, then the chain is of length 1
			chain := 1
			y := x + 1
			for {

				// Check if the next element is in the hash set
				if _, ok := hs[y]; ok {
					chain++
					y++
				} else {
					break
				}
			}

			// Update the longest chain
			if chain > longestChain {
				longestChain = chain
			}
		}
	}
	return longestChain
}
