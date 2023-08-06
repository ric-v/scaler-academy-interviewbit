package levelorder

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

func LevelOrder(A *treeNode) [][]int {
	var result [][]int
	if A == nil {
		return result
	}
	var queue []*treeNode
	queue = append(queue, A)
	for len(queue) > 0 {
		var level []int
		var size = len(queue)
		for i := 0; i < size; i++ {
			var node = queue[0]
			queue = queue[1:]
			level = append(level, node.value)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}

		}
		result = append(result, level)
	}
	return result
}
