package leetcode

// 解法一 stack
func removeDuplicates(s string, k int) string {
	stack, arr := [][2]int{}, []byte{}
	for _, c := range s {
		i := int(c - 'a')
		if len(stack) > 0 && stack[len(stack)-1][0] == i {
			stack[len(stack)-1][1]++
			if stack[len(stack)-1][1] == k {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, [2]int{i, 1})
		}
	}
	for _, pair := range stack {
		c := byte(pair[0] + 'a')
		for i := 0; i < pair[1]; i++ {
			arr = append(arr, c)
		}
	}
	return string(arr)
}

// 解法二 暴力
func removeDuplicates1(s string, k int) string {
	arr, count, tmp := []rune{}, 0, '#'
	for _, v := range s {
		arr = append(arr, v)
		for len(arr) > 0 {
			count = 0
			tmp = arr[len(arr)-1]
			for i := len(arr) - 1; i >= 0; i-- {
				if arr[i] != tmp {
					break
				}
				count++
			}
			if count == k {
				arr = arr[:len(arr)-k]
			} else {
				break
			}
		}
	}
	return string(arr)
}
