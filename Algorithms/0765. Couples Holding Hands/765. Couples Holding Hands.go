package leetcode

func minSwapsCouples(row []int) int {
	if len(row)&1 == 1 {
		return 0
	}
	uf := UnionFind{}
	uf.init(len(row))
	for i := 0; i < len(row)-1; i = i + 2 {
		uf.union(i, i+1)
	}
	for i := 0; i < len(row)-1; i = i + 2 {
		if uf.find(row[i]) != uf.find(row[i+1]) {
			uf.union(row[i], row[i+1])
		}
	}
	return len(row)/2 - uf.count
}
