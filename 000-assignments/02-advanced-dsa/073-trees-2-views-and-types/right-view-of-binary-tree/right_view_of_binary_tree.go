package rightviewofbinarytree

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

func RightViewOfBinaryTree(A *treeNode) []int {
	q := &queue{}
	q.add(A)
	ans := []int{}
	for !q.isEmpty() {
		size := len(q.q)
		ans = append(ans, q.peek().value)
		for size > 0 {
			rNode := q.poll()
			if rNode.right != nil {
				q.add(rNode.right)
			}

			if rNode.left != nil {
				q.add(rNode.left)
			}
			size--
		}
	}

	return ans
}

type queue struct {
	q []*treeNode
}

func (q *queue) add(a *treeNode) {
	q.q = append(q.q, a)
}

func (q *queue) poll() *treeNode {
	if len(q.q) == 0 {
		return nil
	}

	res := q.q[0]
	q.q = q.q[1:]
	return res
}

func (q *queue) peek() *treeNode {
	if len(q.q) == 0 {
		return nil
	}

	return q.q[0]
}

func (q *queue) isEmpty() bool {
	return len(q.q) == 0
}
