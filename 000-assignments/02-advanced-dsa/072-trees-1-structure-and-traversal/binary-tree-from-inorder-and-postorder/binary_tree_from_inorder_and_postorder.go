package binarytreefrominorderandpostorder

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

// BinaryTreeFromInorderAndPostorder constructs a binary tree from its inorder and postorder traversals.
// The inorder traversal is the sequence of nodes visited in an inorder traversal of the tree.
// The postorder traversal is the sequence of nodes visited in a postorder traversal of the tree.
//
// Time complexity: O(n)
// Space complexity: O(n)
func BinaryTreeFromInorderAndPostorder(A, B []int) *treeNode {
	if len(A) == 0 {
		return nil
	}

	// The last element of B is the root of the tree.
	// e.g.: A = [2, 1, 3], B = [2, 3, 1]
	// The root is 1.
	var root *treeNode = treeNode_new(B[len(B)-1])

	// Find the root in A. The elements to the left of the root in A are the left subtree.
	// The elements to the right of the root in A are the right subtree.
	// e.g.: A = [2, 1, 3], B = [2, 3, 1] => root = 1
	// rootIndex = 1
	var rootIndex int
	for i, v := range A {
		if v == root.value {
			rootIndex = i
			break
		}
	}

	// The elements to the left of the root in B are the left subtree.
	// e.g.: A = [2, 1, 3], B = [2, 3, 1] => root = 1
	// rootIndex = 1
	// root.Left = BinaryTreeFromInorderAndPostorder(A[:rootIndex], B[:rootIndex])
	// => root.Left = BinaryTreeFromInorderAndPostorder([2], [2])
	root.left = BinaryTreeFromInorderAndPostorder(A[:rootIndex], B[:rootIndex])

	// The elements to the right of the root in A are the right subtree.
	// e.g.: A = [2, 1, 3], B = [2, 3, 1] => root = 1
	// rootIndex = 1
	// root.Right = BinaryTreeFromInorderAndPostorder(A[rootIndex+1:], B[rootIndex:len(B)-1])
	// => root.Right = BinaryTreeFromInorderAndPostorder([3], [3])
	root.right = BinaryTreeFromInorderAndPostorder(A[rootIndex+1:], B[rootIndex:len(B)-1])
	return root
}
