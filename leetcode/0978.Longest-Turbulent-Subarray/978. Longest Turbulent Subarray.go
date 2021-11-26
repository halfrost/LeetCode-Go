package leetcode

// 解法一 模拟法
func maxTurbulenceSize(arr []int) int {
	inc, dec := 1, 1
	maxLen := min(1, len(arr))
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			inc = dec + 1
			dec = 1
		} else if arr[i-1] > arr[i] {
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 解法二 滑动窗口
func maxTurbulenceSize1(arr []int) int {
	var maxLength int
	if len(arr) == 2 && arr[0] != arr[1] {
		maxLength = 2
	} else {
		maxLength = 1
	}
	left := 0
	for right := 2; right < len(arr); right++ {
		if arr[right] == arr[right-1] {
			left = right
		} else if (arr[right]-arr[right-1])^(arr[right-1]-arr[right-2]) >= 0 {
			left = right - 1
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
