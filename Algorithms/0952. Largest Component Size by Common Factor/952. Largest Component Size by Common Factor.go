package leetcode

// 解法一 并查集 UnionFind
func largestComponentSize(A []int) int {
	maxElement, uf, countMap, res := 0, UnionFind{}, map[int]int{}, 1
	for _, v := range A {
		maxElement = max(maxElement, v)
	}
	uf.init(maxElement + 1)
	for _, v := range A {
		for k := 2; k*k <= v; k++ {
			if v%k == 0 {
				uf.union(v, k)
				uf.union(v, v/k)
			}
		}
	}
	for _, v := range A {
		countMap[uf.find(v)]++
		res = max(res, countMap[uf.find(v)])
	}
	return res
}

// 解法二 UnionFindCount
func largestComponentSize1(A []int) int {
	uf, factorMap := UnionFindCount{}, map[int]int{}
	uf.init(len(A))
	for i, v := range A {
		for k := 2; k*k <= v; k++ {
			if v%k == 0 {
				if _, ok := factorMap[k]; !ok {
					factorMap[k] = i
				} else {
					uf.union(i, factorMap[k])
				}
				if _, ok := factorMap[v/k]; !ok {
					factorMap[v/k] = i
				} else {
					uf.union(i, factorMap[v/k])
				}
			}
		}
		if _, ok := factorMap[v]; !ok {
			factorMap[v] = i
		} else {
			uf.union(i, factorMap[v])
		}
	}
	return uf.maxUnionCount
}
