package leetcode

func checkStraightLine(coordinates [][]int) bool {
	dx0 := coordinates[1][0] - coordinates[0][0]
	dy0 := coordinates[1][1] - coordinates[0][1]
	for i := 1; i < len(coordinates)-1; i++ {
		dx := coordinates[i+1][0] - coordinates[i][0]
		dy := coordinates[i+1][1] - coordinates[i][1]
		if dy*dx0 != dy0*dx { // check cross product
			return false
		}
	}
	return true
}
