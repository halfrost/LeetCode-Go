package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if newColor == color {
		return image
	}
	dfs733(image, sr, sc, newColor)
	return image
}

func dfs733(image [][]int, x, y int, newColor int) {
	if image[x][y] == newColor {
		return
	}
	oldColor := image[x][y]
	image[x][y] = newColor
	for i := 0; i < 4; i++ {
		if (x+dir[i][0] >= 0 && x+dir[i][0] < len(image)) && (y+dir[i][1] >= 0 && y+dir[i][1] < len(image[0])) && image[x+dir[i][0]][y+dir[i][1]] == oldColor {
			dfs733(image, x+dir[i][0], y+dir[i][1], newColor)
		}
	}
}
