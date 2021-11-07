package leetcode

// 解法一
func rotate(matrix [][]int) {
	length := len(matrix)
	// rotate by diagonal 对角线变换
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// rotate by vertical centerline 竖直轴对称翻转
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}

// 解法二
func rotate1(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}
	/* rotate clock-wise = 1. transpose matrix => 2. reverse(matrix[i])

	1   2  3  4      1   5  9  13        13  9  5  1
	5   6  7  8  =>  2   6  10 14  =>    14  10 6  2
	9  10 11 12      3   7  11 15        15  11 7  3
	13 14 15 16      4   8  12 16        16  12 8  4

	*/

	for i := 0; i < n; i++ {
		// transpose, i=rows, j=columns
		// j = i+1, coz diagonal elements didn't change in a square matrix
		for j := i + 1; j < n; j++ {
			swap(matrix, i, j)
		}
		// reverse each row of the image
		matrix[i] = reverse(matrix[i])
	}
}

// swap changes original slice's i,j position
func swap(nums [][]int, i, j int) {
	nums[i][j], nums[j][i] = nums[j][i], nums[i][j]
}

// reverses a row of image, matrix[i]
func reverse(nums []int) []int {
	var lp, rp = 0, len(nums) - 1

	for lp < rp {
		nums[lp], nums[rp] = nums[rp], nums[lp]
		lp++
		rp--
	}
	return nums
}
