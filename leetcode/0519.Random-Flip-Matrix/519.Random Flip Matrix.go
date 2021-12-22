package leetcode

import (
	"math/rand"
)

type Solution struct {
	r     int
	c     int
	total int
	mp    map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{
		r:     m,
		c:     n,
		total: m * n,
		mp:    map[int]int{},
	}
}

func (this *Solution) Flip() []int {
	k := rand.Intn(this.total)
	val := k
	if v, ok := this.mp[k]; ok {
		val = v
	}
	if _, ok := this.mp[this.total-1]; ok {
		this.mp[k] = this.mp[this.total-1]
	} else {
		this.mp[k] = this.total - 1
	}
	delete(this.mp, this.total-1)
	this.total--
	newR, newC := val/this.c, val%this.c
	return []int{newR, newC}
}

func (this *Solution) Reset() {
	this.total = this.r * this.c
	this.mp = map[int]int{}
}
