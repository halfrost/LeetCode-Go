package leetcode

import "math/rand"

type Solution struct {
	M        int
	BlackMap map[int]int
}

func Constructor710(N int, blacklist []int) Solution {
	blackMap := map[int]int{}
	for i := 0; i < len(blacklist); i++ {
		blackMap[blacklist[i]] = 1
	}
	M := N - len(blacklist)
	for _, value := range blacklist {
		if value < M {
			for {
				if _, ok := blackMap[N-1]; ok {
					N--
				} else {
					break
				}
			}
			blackMap[value] = N - 1
			N--
		}
	}
	return Solution{BlackMap: blackMap, M: M}
}

func (this *Solution) Pick() int {
	idx := rand.Intn(this.M)
	if _, ok := this.BlackMap[idx]; ok {
		return this.BlackMap[idx]
	}
	return idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
