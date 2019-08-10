package leetcode

import "fmt"

func numSubarraysWithSum(A []int, S int) int {
	freq, sum, res := make([]int, len(A)+1), 0, 0
	freq[0] = 1
	for _, v := range A {
		t := sum + v - S
		if t >= 0 {
			// 总和有多余的，需要减去 t，除去的方法有 freq[t] 种
			res += freq[t]
		}
		sum += v
		freq[sum]++
		fmt.Printf("freq = %v sum = %v res = %v t = %v\n", freq, sum, res, t)
	}
	return res
}
