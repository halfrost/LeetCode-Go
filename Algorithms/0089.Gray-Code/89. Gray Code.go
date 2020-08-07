package leetcode

// 解法一 递归方法，时间复杂度和空间复杂度都较优
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{}
	num := make([]int, n)
	generateGrayCode(int(1<<uint(n)), 0, &num, &res)
	return res
}

func generateGrayCode(n, step int, num *[]int, res *[]int) {
	if n == 0 {
		return
	}
	*res = append(*res, convertBinary(*num))

	if step%2 == 0 {
		(*num)[len(*num)-1] = flipGrayCode((*num)[len(*num)-1])
	} else {
		index := len(*num) - 1
		for ; index >= 0; index-- {
			if (*num)[index] == 1 {
				break
			}
		}
		if index == 0 {
			(*num)[len(*num)-1] = flipGrayCode((*num)[len(*num)-1])
		} else {
			(*num)[index-1] = flipGrayCode((*num)[index-1])
		}
	}
	generateGrayCode(n-1, step+1, num, res)
	return
}

func convertBinary(num []int) int {
	res, rad := 0, 1
	for i := len(num) - 1; i >= 0; i-- {
		res += num[i] * rad
		rad *= 2
	}
	return res
}

func flipGrayCode(num int) int {
	if num == 0 {
		return 1
	}
	return 0
}

// 解法二 直译
func grayCode1(n int) []int {
	var l uint = 1 << uint(n)
	out := make([]int, l)
	for i := uint(0); i < l; i++ {
		out[i] = int((i >> 1) ^ i)
	}
	return out
}
