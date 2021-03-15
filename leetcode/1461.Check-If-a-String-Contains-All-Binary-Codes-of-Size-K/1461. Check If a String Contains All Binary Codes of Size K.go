package leetcode

import "math"

func hasAllCodes(s string, k int) bool {
	need := int(math.Pow(2.0, float64(k)))
	visited, mask, curr := make([]bool, need), (1<<k)-1, 0
	for i := 0; i < len(s); i++ {
		curr = ((curr << 1) | int(s[i]-'0')) & mask
		if i >= k-1 { // mask 有效位达到了 k 位
			if !visited[curr] {
				need--
				visited[curr] = true
				if need == 0 {
					return true
				}
			}
		}
	}
	return false
}
