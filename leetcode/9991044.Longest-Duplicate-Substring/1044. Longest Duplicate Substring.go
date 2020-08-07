package leetcode

// 解法一 二分搜索 + Rabin-Karp
func longestDupSubstring(S string) string {
	// low, high := 0, len(S)
	// for low < high {
	// 	mid := (low + high + 1) >> 1
	// 	if hasRepeated("", B, mid) {
	// 		low = mid
	// 	} else {
	// 		high = mid - 1
	// 	}
	// }
	return "这个解法还有问题!"
}

// func hashSlice(arr []int, length int) []int {
// 	// hash 数组里面记录 arr 比 length 长出去部分的 hash 值
// 	hash, pl, h := make([]int, len(arr)-length+1), 1, 0
// 	for i := 0; i < length-1; i++ {
// 		pl *= primeRK
// 	}
// 	for i, v := range arr {
// 		h = h*primeRK + v
// 		if i >= length-1 {
// 			hash[i-length+1] = h
// 			h -= pl * arr[i-length+1]
// 		}
// 	}
// 	return hash
// }

// func hasSamePrefix(A, B []int, length int) bool {
// 	for i := 0; i < length; i++ {
// 		if A[i] != B[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func hasRepeated(A, B []int, length int) bool {
// 	hs := hashSlice(A, length)
// 	hashToOffset := make(map[int][]int, len(hs))
// 	for i, h := range hs {
// 		hashToOffset[h] = append(hashToOffset[h], i)
// 	}
// 	for i, h := range hashSlice(B, length) {
// 		if offsets, ok := hashToOffset[h]; ok {
// 			for _, offset := range offsets {
// 				if hasSamePrefix(A[offset:], B[i:], length) {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// 解法二 二分搜索 + 暴力匹配
func longestDupSubstring1(S string) string {
	res := ""
	low, high := 0, len(S)
	for low < high {
		mid := low + (high-low)>>1
		if isDuplicate(mid, S, &res) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return res
}

func isDuplicate(length int, str string, res *string) bool {
	visited := map[string]bool{}
	for i := 0; i+length <= len(str); i++ {
		subStr := str[i : i+length]
		if visited[subStr] {
			*res = subStr
			return true
		}
		visited[subStr] = true
	}
	return false
}
