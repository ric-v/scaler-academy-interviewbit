package bclosestpointstoorigin

import (
	"math"
	"sort"
)

// ClosestPointsToOrigin returns B closest points to origin
// A is a 2D array of points
// B is the number of closest points to return
//
// Time complexity: O(n log n)
// Space complexity: O(n)
func ClosestPointsToOrigin(A [][]int, B int) [][]int {
	// calculate euclidean distance
	for i, a := range A {
		distance := int(math.Pow(0-float64(a[0]), 2) + math.Pow(0-float64(a[1]), 2))
		A[i] = append([]int{distance}, a...)
	}

	// sort by distance
	sort.Slice(A, func(i, j int) bool {
		return A[i][0] < A[j][0]
	})

	// remove distance from result
	var result [][]int
	for i := 0; i < B; i++ {
		result = append(result, A[i][1:])
	}
	return result
}
