package leetcode

func findRedundantConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	uf, res := UnionFind{}, []int{}
	uf.init(len(edges) + 1)
	for i := 0; i < len(edges); i++ {
		if uf.find(edges[i][0]) != uf.find(edges[i][1]) {
			uf.union(edges[i][0], edges[i][1])
		} else {
			res = append(res, edges[i][0])
			res = append(res, edges[i][1])
		}
	}
	return res
}
