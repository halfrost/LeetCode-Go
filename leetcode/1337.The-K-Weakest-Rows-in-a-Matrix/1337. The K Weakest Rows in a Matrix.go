package leetcode

func kWeakestRows(mat [][]int, k int) []int {
	res := []int{}
	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if mat[i][j] == 0 && ((j == 0) || (mat[i][j-1] != 0)) {
				res = append(res, i)
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		if mat[i][len(mat[0])-1] == 1 {
			res = append(res, i)
		}
	}
	return res[:k]
}
