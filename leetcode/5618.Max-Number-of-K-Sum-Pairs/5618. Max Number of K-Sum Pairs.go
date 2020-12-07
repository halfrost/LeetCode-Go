package leetcode

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	c, res, ans, used := []int{}, [][]int{}, 0, make([]bool, len(nums))
	sort.Ints(nums)
	findcombinationSum(nums, k, 0, c, &res, &used)
	fmt.Printf("res = %v\n", res)
	for i := 0; i < len(res); i++ {
		ans = max(ans, len(res[i]))
	}
	return ans
}

func findcombinationSum(nums []int, k, index int, c []int, res *[][]int, used *[]bool) {
	if k <= 0 {
		if k == 0 && len(c) == 2 {
			fmt.Printf("used = %v nums = %v\n", used, nums)
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}

	for i := index; i < len(nums); i++ {
		if !(*used)[i] {
			if nums[i] > k { // 这里可以剪枝优化
				break
			}
			// if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // 这里是去重的关键逻辑
			// 	continue
			// }
			(*used)[i] = true
			c = append(c, nums[i])
			findcombinationSum(nums, k-nums[i], i+1, c, res, used) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
			c = c[:len(c)-1]
			(*used)[i] = false
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
