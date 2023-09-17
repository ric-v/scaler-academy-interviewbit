package main

import (
	"container/heap"
	"fmt"
)

type Elem struct {
	r, c int
	val  int
}

type PQ []Elem

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool { return pq[i].val < pq[j].val }

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Elem))
}

func (pq *PQ) Pop() interface{} {
	top := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return top
}

func bthSmallest(A [][]int, B int) int {
	pq := make(PQ, 0, len(A))
	for r := 0; r < len(A); r++ {
		pq = append(pq, Elem{r, 0, A[r][0]})
	}
	heap.Init(&pq)

	currentSmallest := pq[0].val
	for pq.Len() > 0 && B > 0 {
		e := heap.Pop(&pq).(Elem)
		if e.val != currentSmallest {
			currentSmallest = e.val
			B--
		}
		c := e.c + 1
		if c < len(A[e.r]) {
			pq = append(pq, Elem{e.r, c, A[e.r][c]})
			heap.Init(&pq)
		}
	}

	return currentSmallest
}

func main() {
	A := [][]int{{9, 11, 15}, {10, 15, 17}}
	B := 6
	fmt.Println(bthSmallest(A, B))

	A = [][]int{{5, 9, 11}, {9, 11, 13}, {10, 12, 15}, {13, 14, 16}, {16, 20, 21}}
	B = 12
	fmt.Println(bthSmallest(A, B))
}
