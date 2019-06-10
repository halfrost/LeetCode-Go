package leetcode

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	longRow, longCol, result := max(abs(r0-0), abs(R-r0)), max(abs(c0-0), abs(C-c0)), make([][]int, 0)
	maxDistance := longRow + longCol
	bucket := make([][][]int, maxDistance+1)
	for i := 0; i <= maxDistance; i++ {
		bucket[i] = make([][]int, 0)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			distance := abs(r-r0) + abs(c-c0)
			tmp := []int{r, c}
			bucket[distance] = append(bucket[distance], tmp)
		}
	}
	for i := 0; i <= maxDistance; i++ {
		for _, buk := range bucket[i] {
			result = append(result, buk)
		}
	}
	return result
}
