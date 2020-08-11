package leetcode

func imageSmoother(M [][]int) [][]int {
	res := make([][]int, len(M))
	for i := range M {
		res[i] = make([]int, len(M[0]))
	}
	for y := 0; y < len(M); y++ {
		for x := 0; x < len(M[0]); x++ {
			res[y][x] = smooth(x, y, M)
		}
	}
	return res
}

func smooth(x, y int, M [][]int) int {
	count, sum := 1, M[y][x]
	// Check bottom
	if y+1 < len(M) {
		sum += M[y+1][x]
		count++
	}
	// Check Top
	if y-1 >= 0 {
		sum += M[y-1][x]
		count++
	}
	// Check left
	if x-1 >= 0 {
		sum += M[y][x-1]
		count++
	}
	// Check Right
	if x+1 < len(M[y]) {
		sum += M[y][x+1]
		count++
	}
	// Check Coners
	// Top Left
	if y-1 >= 0 && x-1 >= 0 {
		sum += M[y-1][x-1]
		count++
	}
	// Top Right
	if y-1 >= 0 && x+1 < len(M[0]) {
		sum += M[y-1][x+1]
		count++
	}
	// Bottom Left
	if y+1 < len(M) && x-1 >= 0 {
		sum += M[y+1][x-1]
		count++
	}
	//Bottom Right
	if y+1 < len(M) && x+1 < len(M[0]) {
		sum += M[y+1][x+1]
		count++
	}
	return sum / count
}
