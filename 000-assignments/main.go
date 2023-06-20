package main

import "fmt"

func main() {

	// find rank among all possible pemutations
	// of a given string
	// return ans % 100000000007
	// 1 <= len(s) <= 1000000
	// A = acb -> 2
	// B = bacd -> 7

	fmt.Println(findRank("acb"))
	fmt.Println(findRank("bacd"))
}

func findRank(A string) int {
	var ans int

	// create a map of runes to count their occurance
	// in the string
	runeMap := make(map[rune]int)
	for _, r := range A {
		runeMap[r]++
	}

	// create a map of runes to find their occurance
	// in the string
	// for each rune in the runeMap, find the number
	// of permutations possible
	runeOccurance := make(map[rune]int)
	for _, r := range A {
		ans += runeMap[r] * findFactorial(len(A)-1)
		runeOccurance[r]++
	}

	// for each rune in runeOccurance, find the number
	// of permutations possible
	for _, r := range A {
		ans /= findFactorial(runeOccurance[r])
	}

	return ans % 1000000007
}

func findFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * findFactorial(n-1)
}
