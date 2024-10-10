package main

import "fmt"

// write a program to print hello world
func main() {
	fmt.Println("Hello World")
}

/*
Problem Description

Write a program to input an integer N and a N*N matrix Mat from user and print the matrix in wave form (column wise)

See example for clarifications regarding wave print.

Note: Ensure there is a space character (' ') at the end of the line.

# Problem Constraints

1 <= N <= 103

0 <= Mat[i][j] <= 109

# Input Format

# First line is an integer N

# Next N lines contain N space separated integers representing the matrix Mat

# Output Format

A single line containing N*N integers of matrix Mat in wave form (column wise)

# Example Input

Input 1:

3
4 1 2
7 4 4
3 7 4
Input 2:

2
1 2
3 4

# Example Output

Output 1:

4 7 3 7 4 1 2 4 4
Output 2:

1 3 4 2

# Example Explanation

For Input 1:
We will first iterate the 1st column from top to bottom and print the elements,
then we will iterate the 2nd column from botton to top and print the elements,
then we will iterate the 3rd column from top to bottom and print the elements.
For Input 2:
We will first iterate the 1st column from top to bottom and print the elements,
then we will iterate the 2nd column from bottom to top and print the elements.
*/
func main() {
	var N int
	fmt.Scan(&N)

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		mat[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	for j := 0; j < N; j++ {
		if j%2 == 0 {
			for i := 0; i < N; i++ {
				fmt.Printf("%d ", mat[i][j])
			}
		} else {
			for i := N - 1; i >= 0; i-- {
				fmt.Printf("%d ", mat[i][j])
			}
		}
	}

	fmt.Println()

}
