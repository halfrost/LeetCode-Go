package leetcode

func totalFruit(tree []int) int {
	if len(tree) == 0 {
		return 0
	}
	left, right, counter, res, freq := 0, 0, 1, 1, map[int]int{}
	freq[tree[0]]++
	for left < len(tree) {
		if right+1 < len(tree) && ((counter > 0 && tree[right+1] != tree[left]) || (tree[right+1] == tree[left] || freq[tree[right+1]] > 0)) {
			if counter > 0 && tree[right+1] != tree[left] {
				counter--
			}
			right++
			freq[tree[right]]++
		} else {
			if counter == 0 || (counter > 0 && right == len(tree)-1) {
				res = max(res, right-left+1)
			}
			freq[tree[left]]--
			if freq[tree[left]] == 0 {
				counter++
			}
			left++
		}
	}
	return res
}
