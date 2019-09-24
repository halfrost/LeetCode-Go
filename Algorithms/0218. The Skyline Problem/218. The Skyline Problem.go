package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/template"
)

// 解法一 线段树 Segment Tree，时间复杂度 O(n log n)
func getSkyline(buildings [][]int) [][]int {
	st, ans, lastHeight, check := template.SegmentTree{}, [][]int{}, 0, false
	posMap, pos := discretization218(buildings)
	tmp := make([]int, len(posMap))
	st.Init(tmp, func(i, j int) int {
		return max(i, j)
	})
	for _, b := range buildings {
		st.UpdateLazy(posMap[b[0]], posMap[b[1]-1], b[2])
	}
	for i := 0; i < len(pos); i++ {
		h := st.QueryLazy(posMap[pos[i]], posMap[pos[i]])
		if check == false && h != 0 {
			ans = append(ans, []int{pos[i], h})
			check = true
		} else if i > 0 && h != lastHeight {
			ans = append(ans, []int{pos[i], h})
		}
		lastHeight = h
	}
	return ans
}

func discretization218(positions [][]int) (map[int]int, []int) {
	tmpMap, posArray, posMap := map[int]int{}, []int{}, map[int]int{}
	for _, pos := range positions {
		tmpMap[pos[0]]++
		tmpMap[pos[1]-1]++
		tmpMap[pos[1]]++
	}
	for k := range tmpMap {
		posArray = append(posArray, k)
	}
	sort.Ints(posArray)
	for i, pos := range posArray {
		posMap[pos] = i
	}
	return posMap, posArray
}

// 解法二 扫描线 Sweep Line，时间复杂度 O(n log n)
func getSkyline1(buildings [][]int) [][]int {
	size := len(buildings)
	es := make([]E, 0)
	for i, b := range buildings {
		l := b[0]
		r := b[1]
		h := b[2]
		// 1-- enter
		el := NewE(i, l, h, 0)
		es = append(es, el)
		// 0 -- leave
		er := NewE(i, r, h, 1)
		es = append(es, er)
	}
	skyline := make([][]int, 0)
	sort.Slice(es, func(i, j int) bool {
		if es[i].X == es[j].X {
			if es[i].T == es[j].T {
				if es[i].T == 0 {
					return es[i].H > es[j].H
				}
				return es[i].H < es[j].H
			}
			return es[i].T < es[j].T
		}
		return es[i].X < es[j].X
	})
	pq := NewIndexMaxPQ(size)
	for _, e := range es {
		curH := pq.Front()
		if e.T == 0 {
			if e.H > curH {
				skyline = append(skyline, []int{e.X, e.H})
			}
			pq.Enque(e.N, e.H)
		} else {
			pq.Remove(e.N)
			h := pq.Front()
			if curH > h {
				skyline = append(skyline, []int{e.X, h})
			}
		}
	}
	return skyline
}

// 扫面线伪代码
// events = {{x: L , height: H , type: entering},
// 		  {x: R , height: H , type: leaving}}
// event.SortByX()
// ds = new DS()

// for e in events:
// 	if entering(e):
// 		if e.height > ds.max(): ans += [e.height]
// 		ds.add(e.height)
// 	if leaving(e):
// 		ds.remove(e.height)
// 		if e.height > ds.max(): ans += [ds.max()]

// E define
type E struct { // 定义一个 event 事件
	N int // number 编号
	X int // x 坐标
	H int // height 高度
	T int // type  0-进入 1-离开
}

// NewE define
func NewE(n, x, h, t int) E {
	return E{
		N: n,
		X: x,
		H: h,
		T: t,
	}
}

// IndexMaxPQ define
type IndexMaxPQ struct {
	items []int
	pq    []int
	qp    []int
	total int
}

// NewIndexMaxPQ define
func NewIndexMaxPQ(n int) IndexMaxPQ {
	qp := make([]int, n)
	for i := 0; i < n; i++ {
		qp[i] = -1
	}
	return IndexMaxPQ{
		items: make([]int, n),
		pq:    make([]int, n+1),
		qp:    qp,
	}
}

// Enque define
func (q *IndexMaxPQ) Enque(key, val int) {
	q.total++
	q.items[key] = val
	q.pq[q.total] = key
	q.qp[key] = q.total
	q.swim(q.total)
}

// Front define
func (q *IndexMaxPQ) Front() int {
	if q.total < 1 {
		return 0
	}
	return q.items[q.pq[1]]
}

// Remove define
func (q *IndexMaxPQ) Remove(key int) {
	rank := q.qp[key]
	q.exch(rank, q.total)
	q.total--
	q.qp[key] = -1
	q.sink(rank)
}

func (q *IndexMaxPQ) sink(n int) {
	for 2*n <= q.total {
		k := 2 * n
		if k < q.total && q.less(k, k+1) {
			k++
		}
		if q.less(k, n) {
			break
		}
		q.exch(k, n)
		n = k
	}
}

func (q *IndexMaxPQ) swim(n int) {
	for n > 1 {
		k := n / 2
		if q.less(n, k) {
			break
		}
		q.exch(n, k)
		n = k
	}
}

func (q *IndexMaxPQ) exch(i, j int) {
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
	q.qp[q.pq[i]] = i
	q.qp[q.pq[j]] = j
}

func (q *IndexMaxPQ) less(i, j int) bool {
	return q.items[q.pq[i]] < q.items[q.pq[j]]
}
