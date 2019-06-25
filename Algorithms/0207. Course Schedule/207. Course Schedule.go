package leetcode

// AOV 网的拓扑排序
func canFinish(n int, pre [][]int) bool {
	in := make([]int, n)
	frees := make([][]int, n)
	next := make([]int, 0, n)
	for _, v := range pre {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}
	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]
		for _, vv := range v {
			in[vv]--
			if in[vv] == 0 {
				next = append(next, vv)
			}
		}
	}
	return len(next) == n
}
