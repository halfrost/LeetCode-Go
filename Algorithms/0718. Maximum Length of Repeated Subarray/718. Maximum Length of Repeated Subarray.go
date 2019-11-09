package leetcode

const primeRK = 16777619

// 解法一 二分搜索 + Rabin-Karp
func findLength(A []int, B []int) int {
	low, high := 0, min(len(A), len(B))
	for low < high {
		mid := (low + high + 1) >> 1
		if hasRepeated(A, B, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func hashSlice(arr []int, length int) []int {
	// hash 数组里面记录 arr 比 length 长出去部分的 hash 值
	hash, pl, h := make([]int, len(arr)-length+1), 1, 0
	for i := 0; i < length-1; i++ {
		pl *= primeRK
	}
	for i, v := range arr {
		h = h*primeRK + v
		if i >= length-1 {
			hash[i-length+1] = h
			h -= pl * arr[i-length+1]
		}
	}
	return hash
}

func hasSamePrefix(A, B []int, length int) bool {
	for i := 0; i < length; i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}

func hasRepeated(A, B []int, length int) bool {
	hs := hashSlice(A, length)
	hashToOffset := make(map[int][]int, len(hs))
	for i, h := range hs {
		hashToOffset[h] = append(hashToOffset[h], i)
	}
	for i, h := range hashSlice(B, length) {
		if offsets, ok := hashToOffset[h]; ok {
			for _, offset := range offsets {
				if hasSamePrefix(A[offset:], B[i:], length) {
					return true
				}
			}
		}
	}
	return false
}

// 解法二 DP 动态规划
func findLength1(A []int, B []int) int {
	res, dp := 0, make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}
