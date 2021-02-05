package leetcode

import "container/heap"

type Value struct {
	left  int
	right int
	value int
}

type PQueue []*Value

func (p PQueue) Len() int {
	return len(p)
}

func (p PQueue) Less(a, b int) bool {
	return p[a].value < p[b].value
}

func (p PQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQueue) Push(x interface{}) {
	point := x.(*Value)
	*p = append(*p, point)
}

func (p *PQueue) Pop() interface{} {
	old := *p
	n := len(old)
	// point := old[n-1]
	*p = old[0 : n-1]
	return old[n-1]
}

func kthSmallest(mat [][]int, k int) int {
	nums1 := mat[0]
	for _, nums2 := range mat[1:] {
		nums1 = helper(nums1, nums2, k)
	}

	return nums1[k-1]
}

func helper(nums1, nums2 []int, k int) []int {
	var seen = make(map[[2]int]bool)
	var res []int
	nums1Len, nums2Len := len(nums1), len(nums2)
	pq := PQueue{}
	p := &Value{
		left:  0,
		right: 0,
		value: nums1[0] + nums2[0],
	}
	heap.Push(&pq, p)
	for k > 0 && pq.Len() > 0 {
		v := heap.Pop(&pq)
		l := v.(*Value).left
		r := v.(*Value).right
		res = append(res, v.(*Value).value)
		k -= 1
		if _, ok := seen[[2]int{l + 1, r}]; !ok && l+1 < nums1Len {
			seen[[2]int{l + 1, r}] = true
			heap.Push(&pq, &Value{
				left:  l + 1,
				right: r,
				value: nums1[l+1] + nums2[r],
			})
		}
		if _, ok := seen[[2]int{l, r + 1}]; !ok && r+1 < nums2Len {
			seen[[2]int{l, r + 1}] = true
			heap.Push(&pq, &Value{
				left:  l,
				right: r + 1,
				value: nums1[l] + nums2[r+1],
			})
		}
	}
	return res
}
