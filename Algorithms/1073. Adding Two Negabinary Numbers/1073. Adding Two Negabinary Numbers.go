package leetcode

// 解法一 模拟进位
func addNegabinary(arr1 []int, arr2 []int) []int {
	carry, ans := 0, []int{}
	for i, j := len(arr1)-1, len(arr2)-1; i >= 0 || j >= 0 || carry != 0; {
		if i >= 0 {
			carry += arr1[i]
			i--
		}
		if j >= 0 {
			carry += arr2[j]
			j--
		}
		ans = append([]int{carry & 1}, ans...)
		carry = -(carry >> 1)
	}
	for idx, num := range ans { // 去掉前导 0
		if num != 0 {
			return ans[idx:]
		}
	}
	return []int{0}
}

// 解法二 标准的模拟，但是这个方法不能 AC，因为测试数据超过了 64 位，普通数据类型无法存储
func addNegabinary1(arr1 []int, arr2 []int) []int {
	return intToNegabinary(negabinaryToInt(arr1) + negabinaryToInt(arr2))
}

func negabinaryToInt(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(arr)-1; i++ {
		if res == 0 {
			res += (-2) * arr[i]
		} else {
			res = res * (-2)
			res += (-2) * arr[i]
		}
	}
	return res + 1*arr[len(arr)-1]
}

func intToNegabinary(num int) []int {
	if num == 0 {
		return []int{0}
	}
	res := []int{}

	for num != 0 {
		remainder := num % (-2)
		num = num / (-2)
		if remainder < 0 {
			remainder += 2
			num++
		}
		res = append([]int{remainder}, res...)
	}
	return res
}
