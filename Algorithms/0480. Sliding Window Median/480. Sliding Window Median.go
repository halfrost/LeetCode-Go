package leetcode

import (
	"container/heap"
	"container/list"
	"sort"
)

// 解法一 用链表按照题意实现 时间复杂度 O(n * k) 空间复杂度 O(k)
func medianSlidingWindow(nums []int, k int) []float64 {
	var res []float64
	w := getWindowList(nums[:k], k)
	res = append(res, getMedian(w, k))

	for p1 := k; p1 < len(nums); p1++ {
		w = removeFromWindow(w, nums[p1-k])
		w = insertInWindow(w, nums[p1])
		res = append(res, getMedian(w, k))
	}
	return res
}

func getWindowList(nums []int, k int) *list.List {
	s := make([]int, k)
	copy(s, nums)
	sort.Ints(s)
	l := list.New()
	for _, n := range s {
		l.PushBack(n)
	}
	return l
}

func removeFromWindow(w *list.List, n int) *list.List {
	for e := w.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == n {
			w.Remove(e)
			return w
		}
	}
	return w
}

func insertInWindow(w *list.List, n int) *list.List {
	for e := w.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= n {
			w.InsertBefore(n, e)
			return w
		}
	}
	w.PushBack(n)
	return w
}

func getMedian(w *list.List, k int) float64 {
	e := w.Front()
	for i := 0; i < k/2; e, i = e.Next(), i+1 {
	}
	if k%2 == 1 {
		return float64(e.Value.(int))
	}
	p := e.Prev()
	return (float64(e.Value.(int)) + float64(p.Value.(int))) / 2
}

// 解法二 用两个堆实现 时间复杂度 O(n * log k) 空间复杂度 O(k)
// 用两个堆记录窗口内的值
// 大顶堆里面的元素都比小顶堆里面的元素小
// 如果 k 是偶数，那么两个堆都有 k/2 个元素，中间值就是两个堆顶的元素
// 如果 k 是奇数，那么小顶堆比大顶堆多一个元素，中间值就是小顶堆的堆顶元素
// 删除一个元素，元素都标记到删除的堆中，取 top 的时候注意需要取出没有删除的元素
func medianSlidingWindow1(nums []int, k int) []float64 {
	ans := []float64{}
	minH := MinHeapR{}
	maxH := MaxHeapR{}
	if minH.Len() > maxH.Len()+1 {
		maxH.Push(minH.Pop())
	} else if minH.Len() < maxH.Len() {
		minH.Push(maxH.Pop())
	}
	for i := range nums {
		if minH.Len() == 0 || nums[i] >= minH.Top() {
			minH.Push(nums[i])
		} else {
			maxH.Push(nums[i])
		}
		if i >= k {
			if nums[i-k] >= minH.Top() {
				minH.Remove(nums[i-k])
			} else {
				maxH.Remove(nums[i-k])
			}
		}
		if minH.Len() > maxH.Len()+1 {
			maxH.Push(minH.Pop())
		} else if minH.Len() < maxH.Len() {
			minH.Push(maxH.Pop())
		}
		if minH.Len()+maxH.Len() == k {
			if k%2 == 0 {
				ans = append(ans, float64(minH.Top()+maxH.Top())/2.0)
			} else {
				ans = append(ans, float64(minH.Top()))
			}
		}
		// fmt.Printf("%+v, %+v\n", minH, maxH)
	}
	return ans
}

// IntHeap define
type IntHeap struct {
	data []int
}

// Len define
func (h IntHeap) Len() int { return len(h.data) }

// Swap define
func (h IntHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

// Push define
func (h *IntHeap) Push(x interface{}) { h.data = append(h.data, x.(int)) }

// Pop define
func (h *IntHeap) Pop() interface{} {
	x := h.data[h.Len()-1]
	h.data = h.data[0 : h.Len()-1]
	return x
}

// Top defines
func (h IntHeap) Top() int {
	return h.data[0]
}

// MinHeap define
type MinHeap struct {
	IntHeap
}

// Less define
func (h MinHeap) Less(i, j int) bool { return h.data[i] < h.data[j] }

// MaxHeap define
type MaxHeap struct {
	IntHeap
}

// Less define
func (h MaxHeap) Less(i, j int) bool { return h.data[i] > h.data[j] }

// MinHeapR define
type MinHeapR struct {
	hp, hpDel MinHeap
}

// Len define
func (h MinHeapR) Len() int { return h.hp.Len() - h.hpDel.Len() }

// Top define
func (h *MinHeapR) Top() int {
	for h.hpDel.Len() > 0 && h.hp.Top() == h.hpDel.Top() {
		heap.Pop(&h.hp)
		heap.Pop(&h.hpDel)
	}
	return h.hp.Top()
}

// Pop define
func (h *MinHeapR) Pop() int {
	x := h.Top()
	heap.Pop(&h.hp)
	return x
}

// Push define
func (h *MinHeapR) Push(x int) { heap.Push(&h.hp, x) }

// Remove define
func (h *MinHeapR) Remove(x int) { heap.Push(&h.hpDel, x) }

// MaxHeapR define
type MaxHeapR struct {
	hp, hpDel MaxHeap
}

// Len define
func (h MaxHeapR) Len() int { return h.hp.Len() - h.hpDel.Len() }

// Top define
func (h *MaxHeapR) Top() int {
	for h.hpDel.Len() > 0 && h.hp.Top() == h.hpDel.Top() {
		heap.Pop(&h.hp)
		heap.Pop(&h.hpDel)
	}
	return h.hp.Top()
}

// Pop define
func (h *MaxHeapR) Pop() int {
	x := h.Top()
	heap.Pop(&h.hp)
	return x
}

// Push define
func (h *MaxHeapR) Push(x int) { heap.Push(&h.hp, x) }

// Remove define
func (h *MaxHeapR) Remove(x int) { heap.Push(&h.hpDel, x) }
