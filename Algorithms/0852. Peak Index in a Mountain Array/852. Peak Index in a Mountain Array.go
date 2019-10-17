package leetcode

// 解法一 二分
func peakIndexInMountainArray(A []int) int {
	low, high := 0, len(A)-1
	for low <= high {
		mid := low + (high-low)>>1
		if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
			return mid
		}
		if A[mid] > A[mid+1] && A[mid] < A[mid-1] {
			high = mid - 1
		}
		if A[mid] < A[mid+1] && A[mid] > A[mid-1] {
			low = mid + 1
		}
	}
	return 0
}

// 解法二 二分
func peakIndexInMountainArray1(A []int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := low + (high-low)>>1
		// 如果 mid 较大，则左侧存在峰值，high = m，如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
		if A[mid] > A[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
