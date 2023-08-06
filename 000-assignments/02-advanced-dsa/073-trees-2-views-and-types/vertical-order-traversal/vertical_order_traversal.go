package verticalordertraversal

import "math"

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

func verticalOrderTraversal(A *treeNode) [][]int {
	// Queue to perform level order traversal of the binary tree
	// Here we store the TreeNode and Vertical Level as pair in the queue
	// Root node has Vertical Level 0 and nodes on the left have decreasing Vertical Level and
	// nodes on the right have increasing Vertical Level
	q := make([]pair, 0)
	minVLevel := math.MaxInt32
	maxVLevel := math.MinInt32
	// Map to store the node values at each Vertical Level
	vLevelMap := make(map[int][]int)
	// Output of the vertical order traversal
	output := make([][]int, 0)
	// Start processing the tree nodes only if input is not null
	if A != nil {
		q = append(q, pair{A, 0})
		for len(q) > 0 {
			// Get the pair from queue and assign its contents to TreeNode and Vertical Level
			currentPair := q[0]
			q = q[1:]
			currentNode := currentPair.first
			currentVLevel := currentPair.second
			// Store the min and max of Vertical Level
			minVLevel = min(minVLevel, currentVLevel)
			maxVLevel = max(maxVLevel, currentVLevel)
			// Add the vertical level and node value to the map
			vLevelMap[currentVLevel] = append(vLevelMap[currentVLevel], currentNode.value)
			// Add the left node to the queue with decreasing vertical level
			if currentNode.left != nil {
				q = append(q, pair{currentNode.left, currentVLevel - 1})
			}
			// Add the right node to the queue with increasing vertical level
			if currentNode.right != nil {
				q = append(q, pair{currentNode.right, currentVLevel + 1})
			}
		}
		// Iterate thru’ the min and max keys and add to the output array
		for i := minVLevel; i <= maxVLevel; i++ {
			output = append(output, vLevelMap[i])
		}
	}
	return output
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair struct {
	first  *treeNode
	second int
}
