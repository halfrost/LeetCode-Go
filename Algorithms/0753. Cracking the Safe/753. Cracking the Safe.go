package leetcode

import "math"

const number = "0123456789"

func crackSafe(n int, k int) string {
	if n == 1 {
		return number[:k]
	}
	visit, total := map[string]bool{}, int(math.Pow(float64(k), float64(n)))
	str := make([]byte, 0, total+n-1)
	for i := 1; i != n; i++ {
		str = append(str, '0')
	}
	dfsCrackSafe(total, n, k, &str, &visit)
	return string(str)
}

func dfsCrackSafe(depth, n, k int, str *[]byte, visit *map[string]bool) bool {
	if depth == 0 {
		return true
	}
	for i := 0; i != k; i++ {
		*str = append(*str, byte('0'+i))
		cur := string((*str)[len(*str)-n:])
		if _, ok := (*visit)[cur]; ok != true {
			(*visit)[cur] = true
			if dfsCrackSafe(depth-1, n, k, str, visit) {
				// 只有这里不需要删除
				return true
			}
			delete(*visit, cur)
		}
		// 删除
		*str = (*str)[0 : len(*str)-1]
	}
	return false
}
