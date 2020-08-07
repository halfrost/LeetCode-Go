package leetcode

import (
	"math/rand"
)

// Solution528 define
type Solution528 struct {
	prefixSum []int
}

// Constructor528 define
func Constructor528(w []int) Solution528 {
	prefixSum := make([]int, len(w))
	for i, e := range w {
		if i == 0 {
			prefixSum[i] = e
			continue
		}
		prefixSum[i] = prefixSum[i-1] + e
	}
	return Solution528{prefixSum: prefixSum}
}

// PickIndex define
func (so *Solution528) PickIndex() int {
	n := rand.Intn(so.prefixSum[len(so.prefixSum)-1]) + 1
	low, high := 0, len(so.prefixSum)-1
	for low < high {
		mid := low + (high-low)>>1
		if so.prefixSum[mid] == n {
			return mid
		} else if so.prefixSum[mid] < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
