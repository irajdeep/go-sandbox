package main

import (
	"container/heap"
	"fmt"
)

//Element : ...
type Element struct {
	Val      int
	Priority int
}

type pqList []Element

func (pq pqList) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq pqList) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq pqList) Len() int           { return len(pq) }

func (pq *pqList) Push(x interface{}) {
	*pq = append(*pq, x.(Element))
}

func (pq *pqList) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	(*pq) = (*pq)[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	q := &pqList{}
	heap.Init(q)

	for _, n := range nums {
		count[n]++
	}
	ans := make([]int, k, k)

	for n := range count {
		heap.Push(q, Element{n, count[n]})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	pos := 1

	for q.Len() > 0 {
		if x, ok := heap.Pop(q).(Element); ok {
			ans[k-pos] = x.Val
			pos++
		}
	}
	return ans
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	res := topKFrequent(nums, k)
	fmt.Print(res)
}
