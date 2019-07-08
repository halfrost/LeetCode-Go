package leetcode

// 解法一 暴力打表法
func countArrangement1(N int) int {
	res := []int{0, 1, 2, 3, 8, 10, 36, 41, 132, 250, 700, 750, 4010, 4237, 10680, 24679, 87328, 90478, 435812}
	return res[N]
}

// 解法二 DFS 回溯
func countArrangement(N int) int {
	if N == 0 {
		return 0
	}
	nums, used, p, res := make([]int, N), make([]bool, N), []int{}, [][]int{}
	for i := range nums {
		nums[i] = i + 1
	}
	generatePermutation526(nums, 0, p, &res, &used)
	return len(res)
}

func generatePermutation526(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if !(checkDivisible(nums[i], len(p)+1) || checkDivisible(len(p)+1, nums[i])) { // 关键的剪枝条件
				continue
			}
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation526(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func checkDivisible(num, d int) bool {
	tmp := num / d
	if int(tmp)*int(d) == num {
		return true
	}
	return false
}
