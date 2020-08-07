package leetcode

import (
	"math"
)

// 解法一 二分搜索
func maxSumSubmatrix(matrix [][]int, k int) int {
	// 转换为前缀和
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	sum, absMax, absMaxFound := make([]int, len(matrix)), 0, false
	for y1 := 0; y1 < len(matrix[0]); y1++ {
		for y2 := y1; y2 < len(matrix[0]); y2++ {
			for x := 0; x < len(matrix); x++ {
				if y1 == 0 {
					sum[x] = matrix[x][y2]
				} else {
					sum[x] = matrix[x][y2] - matrix[x][y1-1]
				}
			}
			curMax := kadaneK(sum, k)
			if !absMaxFound || curMax > absMax {
				absMax = curMax
				absMaxFound = true
			}
		}
	}
	return absMax
}

func kadaneK(a []int, k int) int {
	sum, sums, maxSoFar := 0, []int{}, math.MinInt32
	for _, v := range a {
		// 第一次循环会先插入 0，因为 sum 有可能等于 k
		sums = insertSort(sums, sum)
		sum += v
		pos := binarySearchOfKadane(sums, sum-k)
		if pos < len(sums) && sum-sums[pos] > maxSoFar {
			maxSoFar = sum - sums[pos]
		}
	}
	return maxSoFar
}

func binarySearchOfKadane(a []int, v int) int {
	low, high := 0, len(a)
	for low < high {
		mid := low + (high-low)>>1
		if a[mid] < v {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func insertSort(a []int, v int) []int {
	// 类似插入排序，将元素按照从小到大的顺序插入数组
	p := binarySearchOfKadane(a, v)
	a = append(a, 0)
	// 把 p 后面的元素全部往后移，把 p 位置空出来放 v
	copy(a[p+1:], a[p:len(a)-1])
	a[p] = v
	return a
}

// 解法二 暴力解法，超时
func maxSumSubmatrix1(matrix [][]int, k int) int {
	minNum := math.MaxInt64
	for row := range matrix {
		for col := 1; col < len(matrix[row]); col++ {
			if matrix[row][col] < minNum {
				minNum = matrix[row][col]
			}
		}
	}
	for row := range matrix {
		for col := 1; col < len(matrix[row]); col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}
	for i := k; ; i-- {
		if findSumSubmatrix(matrix, i) > 0 {
			return i
		}
	}
}

func findSumSubmatrix(matrix [][]int, target int) int {
	m, n, res := len(matrix), len(matrix[0]), 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			counterMap, sum := make(map[int]int, m), 0
			counterMap[0] = 1 // 题目保证一定有解，所以这里初始化是 1
			for row := 0; row < m; row++ {
				if i > 0 {
					sum += matrix[row][j] - matrix[row][i-1]
				} else {
					sum += matrix[row][j]
				}
				res += counterMap[sum-target]
				counterMap[sum]++
			}
		}
	}
	return res
}
