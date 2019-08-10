package leetcode

import (
	"container/list"
	"fmt"
	"strconv"
)

// 解法一
func calculate(s string) int {
	i, stack, result, sign := 0, list.New(), 0, 1 // 记录加减状态
	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else if s[i] <= '9' && s[i] >= '0' { // 获取一段数字
			base, v := 10, int(s[i]-'0')
			for i+1 < len(s) && s[i+1] <= '9' && s[i+1] >= '0' {
				v = v*base + int(s[i+1]-'0')
				i++
			}
			result += v * sign
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '(' { // 把之前计算结果及加减状态压栈，开始新的计算
			stack.PushBack(result)
			stack.PushBack(sign)
			result = 0
			sign = 1
			i++
		} else if s[i] == ')' { // 新的计算结果 * 前一个加减状态 + 之前计算结果
			result = result*stack.Remove(stack.Back()).(int) + stack.Remove(stack.Back()).(int)
			i++
		}
	}
	return result
}

// 解法二
func calculate1(s string) int {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == ')' {
			tmp, index := "", len(stack)-1
			for ; index >= 0; index-- {
				if stack[index] == '(' {
					break
				}
			}
			tmp = string(stack[index+1 : len(stack)])
			stack = stack[:index]
			res := strconv.Itoa(calculateStr(tmp))
			for j := 0; j < len(res); j++ {
				stack = append(stack, res[j])
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	fmt.Printf("stack = %v\n", string(stack))
	return calculateStr(string(stack))
}

func calculateStr(str string) int {
	s, nums, tmpStr, res := []byte{}, []int{}, "", 0
	// 处理符号的问题，++得+，--得+，+-、-+得-
	for i := 0; i < len(str); i++ {
		if len(s) > 0 && s[len(s)-1] == '+' && str[i] == '+' {
			continue
		} else if len(s) > 0 && s[len(s)-1] == '+' && str[i] == '-' {
			s[len(s)-1] = '-'
		} else if len(s) > 0 && s[len(s)-1] == '-' && str[i] == '+' {
			continue
		} else if len(s) > 0 && s[len(s)-1] == '-' && str[i] == '-' {
			s[len(s)-1] = '+'
		} else {
			s = append(s, str[i])
		}
	}
	str = string(s)
	s = []byte{}
	for i := 0; i < len(str); i++ {
		if isDigital(str[i]) {
			tmpStr += string(str[i])
		} else {
			num, _ := strconv.Atoi(tmpStr)
			nums = append(nums, num)
			tmpStr = ""
			s = append(s, str[i])
		}
	}
	if tmpStr != "" {
		num, _ := strconv.Atoi(tmpStr)
		nums = append(nums, num)
		tmpStr = ""
	}
	res = nums[0]
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			res += nums[i+1]
		} else {
			res -= nums[i+1]
		}
	}
	fmt.Printf("s = %v nums = %v res = %v\n", string(s), nums, res)
	return res
}
