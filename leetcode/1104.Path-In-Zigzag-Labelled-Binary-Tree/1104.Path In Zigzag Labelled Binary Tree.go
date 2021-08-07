package leetcode

func pathInZigZagTree(label int) []int {
	level := getLevel(label)
	ans := []int{label}
	curIndex := label - (1 << level)
	parent := 0
	for level >= 1 {
		parent, curIndex = getParent(curIndex, level)
		ans = append(ans, parent)
		level--
	}
	ans = reverse(ans)
	return ans
}

func getLevel(label int) int {
	level := 0
	nums := 0
	for {
		nums += 1 << level
		if nums >= label {
			return level
		}
		level++
	}
}

func getParent(index int, level int) (parent int, parentIndex int) {
	parentIndex = 1<<(level-1) - 1 + (index/2)*(-1)
	parent = 1<<(level-1) + parentIndex
	return
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
