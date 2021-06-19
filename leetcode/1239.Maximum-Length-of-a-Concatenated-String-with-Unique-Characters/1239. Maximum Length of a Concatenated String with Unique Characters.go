package leetcode

import (
	"math/bits"
)

func maxLength(arr []string) int {
	c, res := []uint32{}, 0
	for _, s := range arr {
		var mask uint32
		for _, c := range s {
			mask = mask | 1<<(c-'a')
		}
		if len(s) != bits.OnesCount32(mask) { // 如果字符串本身带有重复的字符，需要排除
			continue
		}
		c = append(c, mask)
	}
	dfs(c, 0, 0, &res)
	return res
}

func dfs(c []uint32, index int, mask uint32, res *int) {
	*res = max(*res, bits.OnesCount32(mask))
	for i := index; i < len(c); i++ {
		if mask&c[i] == 0 {
			dfs(c, i+1, mask|c[i], res)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
