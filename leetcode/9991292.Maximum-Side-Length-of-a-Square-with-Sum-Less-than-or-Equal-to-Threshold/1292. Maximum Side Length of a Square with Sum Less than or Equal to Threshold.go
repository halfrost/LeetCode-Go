package leetcode

// 解法一 二维前缀和。sum[i+1][j+1] 表示左上角到 (i,j) 的矩形和。
// 边长 ans 从 0 开始递增：每次检查是否存在一个边长为 ans+1、且和不超过 threshold 的正方形，
// 存在就把 ans 加一。由于答案具有单调性，整体相当于一次扫描，时间复杂度 O(m*n)。
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + mat[i][j]
		}
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 尝试以 (i,j) 为右下角、边长为 ans+1 的正方形
			k := ans + 1
			if i >= k && j >= k {
				area := sum[i][j] - sum[i-k][j] - sum[i][j-k] + sum[i-k][j-k]
				if area <= threshold {
					ans++
				}
			}
		}
	}
	return ans
}
