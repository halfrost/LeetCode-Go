package leetcode

// 解法一 二分思想
func maxDepthAfterSplit(seq string) []int {
	stack, maxDepth, res := 0, 0, []int{}
	for _, v := range seq {
		if v == '(' {
			stack++
			maxDepth = max(stack, maxDepth)
		} else {
			stack--
		}
	}
	stack = 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			stack++
			if stack <= maxDepth/2 {
				res = append(res, 0)
			} else {
				res = append(res, 1)
			}
		} else {
			if stack <= maxDepth/2 {
				res = append(res, 0)
			} else {
				res = append(res, 1)
			}
			stack--
		}
	}
	return res
}

// 解法二 模拟
func maxDepthAfterSplit1(seq string) []int {
	stack, top, res := make([]int, len(seq)), -1, make([]int, len(seq))
	for i, r := range seq {
		if r == ')' {
			res[i] = res[stack[top]]
			top--
			continue
		}
		top++
		stack[top] = i
		res[i] = top % 2
	}
	return res
}
