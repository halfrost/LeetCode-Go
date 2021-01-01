package leetcode

func canFormArray(arr []int, pieces [][]int) bool {
	arrMap := map[int]int{}
	for i, v := range arr {
		arrMap[v] = i
	}
	for i := 0; i < len(pieces); i++ {
		order := -1
		for j := 0; j < len(pieces[i]); j++ {
			if _, ok := arrMap[pieces[i][j]]; !ok {
				return false
			}
			if order == -1 {
				order = arrMap[pieces[i][j]]
			} else {
				if arrMap[pieces[i][j]] == order+1 {
					order = arrMap[pieces[i][j]]
				} else {
					return false
				}
			}
		}
	}
	return true
}
