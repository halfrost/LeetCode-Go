package leetcode

// 解法一 array 优化版
func subarrayBitwiseORs(A []int) int {
	res, cur, isInMap := []int{}, []int{}, make(map[int]bool)
	cur = append(cur, 0)
	for _, v := range A {
		var cur2 []int
		for _, vv := range cur {
			tmp := v | vv
			if !inSlice(cur2, tmp) {
				cur2 = append(cur2, tmp)
			}
		}
		if !inSlice(cur2, v) {
			cur2 = append(cur2, v)
		}
		cur = cur2
		for _, vv := range cur {
			if _, ok := isInMap[vv]; !ok {
				isInMap[vv] = true
				res = append(res, vv)
			}
		}
	}
	return len(res)
}

func inSlice(A []int, T int) bool {
	for _, v := range A {
		if v == T {
			return true
		}
	}
	return false
}

// 解法二 map 版
func subarrayBitwiseORs1(A []int) int {
	res, t := map[int]bool{}, map[int]bool{}
	for _, num := range A {
		r := map[int]bool{}
		r[num] = true
		for n := range t {
			r[(num | n)] = true
		}
		t = r
		for n := range t {
			res[n] = true
		}
	}
	return len(res)
}
