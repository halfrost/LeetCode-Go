package leetcode

import (
	"strconv"
)

func decodeString(s string) string {
	stack, res := []string{}, ""
	for _, str := range s {
		if len(stack) == 0 || (len(stack) > 0 && str != ']') {
			stack = append(stack, string(str))
		} else {
			tmp := ""
			for stack[len(stack)-1] != "[" {
				tmp = stack[len(stack)-1] + tmp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			index, repeat := 0, ""
			for index = len(stack) - 1; index >= 0; index-- {
				if stack[index] >= "0" && stack[index] <= "9" {
					repeat = stack[index] + repeat
				} else {
					break
				}
			}
			nums, _ := strconv.Atoi(repeat)
			copyTmp := tmp
			for i := 0; i < nums-1; i++ {
				tmp += copyTmp
			}
			for i := 0; i < len(repeat)-1; i++ {
				stack = stack[:len(stack)-1]
			}
			stack[index+1] = tmp
		}
	}
	for _, s := range stack {
		res += s
	}
	return res
}
