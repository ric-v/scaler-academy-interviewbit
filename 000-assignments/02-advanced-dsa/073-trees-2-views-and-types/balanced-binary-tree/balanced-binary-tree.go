package balancedbinarytree

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

func BalancedBinaryTree(A *treeNode) int {

	if A == nil {
		return 1
	}

	leftHeight := height(A.left)
	rightHeight := height(A.right)

	if abs(leftHeight-rightHeight) <= 1 && BalancedBinaryTree(A.left) == 1 && BalancedBinaryTree(A.right) == 1 {
		return 1
	}

	return 0
}

func height(A *treeNode) int {
	if A == nil {
		return 0
	}

	return 1 + max(height(A.left), height(A.right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
