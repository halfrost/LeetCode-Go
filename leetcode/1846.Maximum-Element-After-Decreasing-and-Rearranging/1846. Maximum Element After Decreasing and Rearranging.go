package leetcode

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	cnt := make([]int, n+1)
	for _, v := range arr {
		cnt[min(v, n)]++
	}
	miss := 0
	for _, c := range cnt[1:] {
		if c == 0 {
			miss++
		} else {
			miss -= min(c-1, miss)
		}
	}
	return n - miss
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
