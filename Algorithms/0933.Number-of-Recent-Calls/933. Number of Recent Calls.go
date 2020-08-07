package leetcode

import "sort"

type RecentCounter struct {
	list []int
}

func Constructor933() RecentCounter {
	return RecentCounter{
		list: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.list = append(this.list, t)
	index := sort.Search(len(this.list), func(i int) bool { return this.list[i] >= t-3000 })
	if index < 0 {
		index = -index - 1
	}
	return len(this.list) - index
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
