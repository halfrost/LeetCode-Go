package leetcode

// 解法一 模拟法
func maxTurbulenceSize(A []int) int {
	inc, dec := 1, 1
	maxLen := min(1, len(A))
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			inc = dec + 1
			dec = 1
		} else if A[i-1] > A[i] {
			dec = inc + 1
			inc = 1
		} else {
			inc = 1
			dec = 1
		}
		maxLen = max(maxLen, max(inc, dec))
	}
	return maxLen
}

// 解法二 滑动窗口
func maxTurbulenceSize1(A []int) int {
	if len(A) == 1 {
		return 1
	}
	// flag > 0 代表下一个数要大于前一个数，flag < 0 代表下一个数要小于前一个数
	res, left, right, flag, lastNum := 0, 0, 0, A[1]-A[0], A[0]
	for left < len(A) {
		if right < len(A)-1 && ((A[right+1] > lastNum && flag > 0) || (A[right+1] < lastNum && flag < 0) || (right == left)) {
			right++
			flag = lastNum - A[right]
			lastNum = A[right]
		} else {
			if right != left && flag != 0 {
				res = max(res, right-left+1)
			}
			left++
		}
	}
	return max(res, 1)
}
