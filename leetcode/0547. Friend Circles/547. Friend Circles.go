package leetcode

// 解法一 并查集

// UnionFind defind
type UnionFind struct {
	parent, rank []int
	count        int
}

func (uf *UnionFind) init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

func (uf *UnionFind) find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// compress path
	for p != uf.parent[p] {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

func (uf *UnionFind) union(p, q int) {
	proot := uf.find(p)
	qroot := uf.find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

func (uf *UnionFind) totalCount() int {
	return uf.count
}

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	uf := UnionFind{}
	uf.init(n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.count
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
