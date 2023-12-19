package main

import (
	"fmt"
)

func maxTastiness(A []int, B int) int {
	low, high := 0, 0
	for _, x := range A {
		high += x
	}

	for low < high {
		mid := (low + high + 1) / 2
		partitions := countPartitions(A, mid)

		if partitions >= B {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

func countPartitions(A []int, tastinessLimit int) int {
	partitions := 1
	currentTastiness := 0

	for _, x := range A {
		currentTastiness += x
		if currentTastiness > tastinessLimit {
			partitions++
			currentTastiness = x
		}
	}

	return partitions
}

func main() {
	input1 := []int{1, 2, 3, 4, 5}
	B1 := 2
	output1 := maxTastiness(input1, B1)
	fmt.Println(output1) // Output: 15

	input2 := []int{2, 2, 2, 2}
	B2 := 3
	output2 := maxTastiness(input2, B2)
	fmt.Println(output2) // Output: 12

	input3 := []int{2, 1, 8, 7, 24, 49, 12, 21, 1}
	B3 := 6
	output3 := maxTastiness(input3, B3)
	fmt.Println(output3) // Output: 223
}
