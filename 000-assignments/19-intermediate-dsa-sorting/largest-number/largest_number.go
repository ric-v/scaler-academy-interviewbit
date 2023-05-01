package largestnumber

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(A []int) string {
	// Convert the array of integers to array of strings
	str := make([]string, len(A))
	for i, v := range A {
		str[i] = strconv.Itoa(v)
	}

	// Sort the array of strings in descending order
	sort.Slice(str, func(i, j int) bool {
		return str[i]+str[j] > str[j]+str[i]
	})

	// If the largest number is 0, then the entire number is 0
	if str[0] == "0" {
		return "0"
	}

	// Otherwise, concatenate the string
	return strings.Join(str, "")
}
