package josephusproblem2

// JosephusProblem2 solves the Josephus problem using recursion
//
// The Josephus problem is a theoretical problem related to a certain counting-out game.
// People are standing in a circle waiting to be executed. Counting begins at a specified point in the circle and proceeds around the circle in a specified direction.
// After a specified number of people are skipped, the next person is executed. The procedure is repeated with the remaining people, starting with the next person,
// going in the same direction and skipping the same number of people, until only one person remains, and is freed.
//
// Time complexity: O(n)
// Space complexity: O(1)
func JosephusProblem2(A int) int {
	var result int

	for i := 1; i <= A; i++ {

		// The position of the survivor is 2*(n - 2^floor(log2(n))) + 1
		result = (result + 2) % i
	}
	return result + 1
}
