package bishop

// Given the position of a Bishop (A, B) on an 8 * 8 chessboard.
// Your task is to count the total number of squares that can be visited by the Bishop in one move.
// The position of the Bishop is denoted using row and column number of the chessboard.

// Input 1:
//
//	A = 4
//	B = 4
//
// Output 1:
//
//	13
func BishopMoves(A, B int) int {

	var count int

	// total squares to top left
	count += min(A, B) - 1

	// total squares to top right
	count += min(A, 9-B) - 1

	// total squares to bottom left
	count += 8 - max(A, 9-B)

	// total squares to bottom right
	count += 8 - max(A, B)

	return count
}

func min(A, B int) int {
	if A < B {
		return A
	}
	return B
}

func max(A, B int) int {
	if A > B {
		return A
	}
	return B
}
