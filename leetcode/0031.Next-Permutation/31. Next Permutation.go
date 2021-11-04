// https://leetcode.com/problems/next-permutation/discuss/1554932/Go-Submission-with-Explanation
// Time O(N) , Space: O(1)

package leetcode

// [2,(3),6,5,4,1] -> 2,(4),6,5,(3),1 -> 2,4, 1,3,5,6
func nextPermutation(nums []int) {
	var n = len(nums)
	var pIdx = checkPermutationPossibility(nums)
	if pIdx == -1 {
		reverse(nums, 0, n-1)
		return
	}

	var rp = len(nums) - 1
	// start from right most to leftward,find the first number which is larger than PIVOT
	for rp > 0 {
		if nums[rp] > nums[pIdx] {
			nums[rp], nums[pIdx] = nums[pIdx], nums[rp]
			break
		} else {
			rp--
		}
	}
	// Finally, Reverse all elements which are right from pivot
	reverse(nums, pIdx+1, n-1)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int, s int, e int) {
	for s < e {
		swap(nums, s, e)
		s++
		e--
	}
}

// checkPermutationPossibility returns 1st occurrence Index where
// value is not in increasing order(from right to left)
// returns -1 if not found(it's already in its last permutation)
func checkPermutationPossibility(nums []int) (idx int) {
	// search right to left for 1st number(from right) that is not in increasing order
	var rp = len(nums) - 1
	for rp > 0 {
		if nums[rp-1] < nums[rp] {
			idx = rp - 1
			return idx
		}
		rp--
	}
	return -1
}
