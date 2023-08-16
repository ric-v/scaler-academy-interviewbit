package main

import "fmt"

func main() {
	// fmt.Println(countValidParanthesis("()()"))
	fmt.Println(waitingTime([]int{2, 3, 2}, 6))
	fmt.Println(waitingTime([]int{5, 1, 1, 1, 1}, 8))
}

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

// given a binary tryee represnt by root A, find the Bth smallest element in the tree
func kthsmallest(A *treeNode, B int) int {

	// use inorder traversal to get the sorted array
	var arr []int
	inorder(A, &arr)

	// return the Bth element
	return arr[B-1]
}

func inorder(A *treeNode, arr *[]int) {
	if A == nil {
		return
	}
	inorder(A.left, arr)
	*arr = append(*arr, A.value)
	inorder(A.right, arr)
}

// there are N people in a queue to buy tickets, where 0th person is at the front of the line and the (n-1)th person is at the back of the line
// you are given a 0 indexed integer array A and n where number of tickets that ith person would like to buy is A[i]
// each person takes exactly 1 second to buy a ticket. A person can only buy 1 ticket at a time and has to go back to the end of the line after buying the ticket
// which happens instantaneously. If a person does not have any tickets left to buy, the person will leave the line.
// return the time taken for the person at position B (0 indexed) to finish buying tickets
// A = [2, 3, 2], B = 2 => 6 (2nd person buys 1st ticket at 2 seconds, 2nd ticket at 4 seconds and 3rd ticket at 6 seconds)
// A = [5, 1, 1, 1, 1], B = 0 => 8 (0th person buys 1st ticket at 1 second, 2nd ticket at 2 seconds, 3rd ticket at 3 seconds, 4th ticket at 4 seconds and 5th ticket at 5 seconds)
func waitingTime(A []int, B int) int {
	var time int
	var queue []int
	for i := 0; i < len(A); i++ {
		queue = append(queue, A[i])
	}

	for len(queue) != 0 {
		if queue[0] == 1 {
			time++
			if B == 0 {
				return time - 1
			}
			queue = queue[1:]
			B--
		} else {
			time++
			queue[0]--
			queue = append(queue[1:], queue[0])
			B--
		}
	}
	return time
}

// get the count of longest valid parenthesis
// A = ()()) => 4
// A = (()())() => 8
func countValidParanthesis(A string) int {

	var stack []int
	stack = append(stack, -1)
	var max int

	for i := 0; i < len(A); i++ {
		if A[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				max = maxInt(max, i-stack[len(stack)-1])
			} else {
				stack = append(stack, i)
			}
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
