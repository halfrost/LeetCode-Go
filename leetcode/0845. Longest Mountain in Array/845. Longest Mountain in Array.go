package leetcode

func longestMountain(A []int) int {
	left, right, res, isAscending := 0, 0, 0, true
	for left < len(A) {
		if right+1 < len(A) && ((isAscending == true && A[right+1] > A[left] && A[right+1] > A[right]) || (right != left && A[right+1] < A[right])) {
			if A[right+1] < A[right] {
				isAscending = false
			}
			right++
		} else {
			if right != left && isAscending == false {
				res = max(res, right-left+1)
			}
			left++
			if right < left {
				right = left
			}
			if right == left {
				isAscending = true
			}
		}
	}
	return res
}
