package priorityq

// Element : ...
type Element struct {
	Val      int
	Priority int
}

type pqList []Element

func (pq pqList) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq pqList) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq pqList) Len() int           { return len(pq) }

func (pq *pqList) Push(e Element) {
	*pq = append(*pq, e)
}

func (pq *pqList) Pop() Element {
	n := len(*pq)
	x := (*pq)[n-1]
	(*pq) = (*pq)[0 : n-1]
	return x
}
