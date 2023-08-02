package inordertraversal

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

// InorderTraversal traverses a binary tree in inorder fashion
// and returns the values in a slice
// 
// Time complexity: O(n)
// Space complexity: O(n)
func InorderTraversal(A *treeNode) []int {

	// base case
	if A == nil {
		return []int{}
	}

	// recursive case
	var result []int
	inorder(A, &result)
	return result
}

// inorder traverses a binary tree in inorder fashion 
// and appends the values to the result slice
func inorder(A *treeNode, result *[]int) {
	if A == nil {
		return
	}

	inorder(A.left, result)
	*result = append(*result, A.value)
	inorder(A.right, result)
}
