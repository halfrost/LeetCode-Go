package leetcode

import "container/heap"

func topKFrequent(words []string, k int) []string {
	m := map[string]int{}
	for _, word := range words {
		m[word]++
	}
	pq := &PQ{}
	heap.Init(pq)
	for w, c := range m {
		heap.Push(pq, &wordCount{w, c})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		wc := heap.Pop(pq).(*wordCount)
		res[i] = wc.word
	}
	return res
}

type wordCount struct {
	word string
	cnt  int
}

type PQ []*wordCount

func (pq PQ) Len() int      { return len(pq) }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQ) Less(i, j int) bool {
	if pq[i].cnt == pq[j].cnt {
		return pq[i].word > pq[j].word
	}
	return pq[i].cnt < pq[j].cnt
}
func (pq *PQ) Push(x interface{}) {
	tmp := x.(*wordCount)
	*pq = append(*pq, tmp)
}
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
