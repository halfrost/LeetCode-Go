package leetcode

func numSimilarGroups(A []string) int {
	uf := UnionFind{}
	uf.init(len(A))
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if isSimilar(A[i], A[j]) {
				uf.union(i, j)
			}
		}
	}
	return uf.totalCount()
}

func isSimilar(a, b string) bool {
	var n int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			n++
			if n > 2 {
				return false
			}
		}
	}
	return true
}
