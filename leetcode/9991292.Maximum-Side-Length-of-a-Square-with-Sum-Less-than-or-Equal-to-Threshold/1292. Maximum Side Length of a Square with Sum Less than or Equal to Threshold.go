package leetcode

func maxSideLength(mat [][]int, threshold int) int {
	return 0
	// m, n, sum := len(mat), len(mat[0]), make([][]int, len(mat[0])+1, len(mat[0])+1)
	// for i := range sum {
	// 	sum[i] = make([]int, n+1, n+1)
	// }
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + mat[i][j]
	// 	}
	// }
	// low, high := 0, min(m, n)
	// for low < high {
	// 	mid := low + (high-low)>>1
	// 	if !inThreshold(&sum, threshold, mid) {
	// 		high = mid + 1
	// 	} else {
	// 		low = mid
	// 	}
	// }
	// if inThreshold(&sum, threshold, high) {
	// 	return high
	// }
	// return low
}

// func min(a int, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// func inThreshold(sum *[][]int, threshold int, length int) bool {
// 	for i := 1; i < len(*sum); i++ {
// 		for j := 1; j < len((*sum)[0]); j++ {
// 			if i < length || j < length {
// 				continue
// 			}
// 			if (*sum)[i][j]+(*sum)[i-length][j-length]-(*sum)[i-length][j]-(*sum)[i][j-length] <= threshold {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
