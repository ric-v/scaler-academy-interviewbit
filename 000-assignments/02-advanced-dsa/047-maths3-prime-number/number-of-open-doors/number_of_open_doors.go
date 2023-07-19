package numberofopendoors

// NumberOfOpenDoors returns the number of doors that are open after A passes
//
// Time complexity: O(sqrt(A))
// Space complexity: O(1)
func NumberOfOpenDoors(A int) int {
	var result int
	for i := 1; i*i <= A; i++ {
		result++
	}
	return result
}
