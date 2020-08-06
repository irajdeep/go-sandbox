func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, n := range nums {
		seen[n]++
	}

	q := &pqList{}
	heap.Init(q)
	for val, cnt := range seen {
		heap.Push(q, Element{val, cnt})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]int, k, k)
	for i := 1; i <= k; i++ {
		v := heap.Pop(q)
		fmt.Println(v)
		if s, ok := v.(Element); ok {
			ans[k-i] = s.Val
		}
	}
	return ans
}