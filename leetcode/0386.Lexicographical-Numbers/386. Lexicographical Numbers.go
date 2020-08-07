package leetcode

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	dfs386(1, n, &res)
	return res
}

func dfs386(x, n int, res *[]int) {
	limit := (x + 10) / 10 * 10
	for x <= n && x < limit {
		*res = append(*res, x)
		if x*10 <= n {
			dfs386(x*10, n, res)
		}
		x++
	}
}
