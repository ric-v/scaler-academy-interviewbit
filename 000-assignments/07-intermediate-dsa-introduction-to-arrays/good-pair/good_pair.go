package goodpair

// GoodPair returns 1 if there is a pair of numbers in A that add up to B, 0 otherwise
//
// Time complexity: O(n^2)
// Space complexity: O(1)
//
// Example:
//
//	GoodPair([]int{1, 2, 3, 4}, 7) // 1
//
//	i=[0->n]     j=[i+1->n]    	sum=A[i]+A[j]
//
// -------------------------------------------
//
//	0   		i+1=1  			1+2=3
//	1   		i+1=2  			2+3=5
//	1   		i+1=2  			2+3=5
//	2   		i+1=3  			3+4=7  <--- 7==7 return 1
func GoodPair(A []int, B int) int {

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i]+A[j] == B {
				return 1
			}
		}
	}
	return 0
}

func GoodPair_Unoptimized(A []int, B int) int {

	for i := 0; i < len(A); i++ {

		for j := 0; j < len(A); j++ {

			if i != j && A[i]+A[j] == B {
				return 1
			}
		}
	}
	return 0
}
