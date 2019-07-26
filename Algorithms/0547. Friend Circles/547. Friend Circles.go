package leetcode

// 解法一 并查集

// UninonSet defind
type UninonSet struct {
	roots []int
}

func (us UninonSet) init() {
	for i := range us.roots {
		us.roots[i] = i
	}
}

func (us UninonSet) findRoot(i int) int {
	root := i
	for root != us.roots[root] {
		root = us.roots[root]
	}
	for i != us.roots[i] {
		tmp := us.roots[i]
		us.roots[i] = root
		i = tmp
	}
	return root
}

func (us UninonSet) union(p, q int) {
	qroot := us.findRoot(q)
	proot := us.findRoot(p)
	us.roots[proot] = qroot
}

func findCircleNum(M [][]int) int {
	n, count := len(M), 0
	if n == 0 {
		return 0
	}
	us := UninonSet{}
	us.roots = make([]int, n+1)
	us.init()
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 {
				x, y := us.findRoot(i), us.findRoot(j)
				if x != y {
					us.union(x, y)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if us.roots[i] == i {
			count++
		}
	}
	return count
}

// 解法二 FloodFill DFS 暴力解法
func findCircleNum1(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	visited := make([]bool, len(M))
	res := 0
	for i := range M {
		if !visited[i] {
			dfs547(M, i, visited)
			res++
		}
	}
	return res
}

func dfs547(M [][]int, cur int, visited []bool) {
	visited[cur] = true
	for j := 0; j < len(M[cur]); j++ {
		if !visited[j] && M[cur][j] == 1 {
			dfs547(M, j, visited)
		}
	}
}
