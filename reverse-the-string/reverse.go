package reverse

import (
	"strings"
)

func Reverse(A string) string {

	var out string

	// split A by whitespace
	split := strings.Split(A, " ")

	// for each word in reverse, append to output
	for i := len(split) - 1; i >= 0; i-- {

		// if split[i] is empty string, skip
		if split[i] != "" {
			out += " " + split[i]
		}
	}

	// trim any whitespace
	return strings.TrimSpace(out)
}
