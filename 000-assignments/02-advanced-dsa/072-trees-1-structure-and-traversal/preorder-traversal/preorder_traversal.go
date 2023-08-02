package preordertraversal

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

// PreorderTraversal returns the preorder traversal of a binary tree
//
// Time complexity: O(n)
// Space complexity: O(n)
func PreorderTraversal(A *treeNode) []int {

	// base case
	if A == nil {
		return []int{}
	}

	// recursive case
	var result []int
	preorderT(A, &result)
	return result
}

// preorder is a helper function for PreorderTraversal
// It appends the preorder traversal of a binary tree to the result slice
// It is a recursive function
func preorderT(A *treeNode, result *[]int) {
	if A == nil {
		return
	}

	*result = append(*result, A.value)
	preorderT(A.left, result)
	preorderT(A.right, result)
}
