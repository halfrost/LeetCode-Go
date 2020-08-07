package structures

// Point 定义了一个二维坐标点
type Point struct {
	X, Y int
}

// Intss2Points 把 [][]int 转换成 []Point
func Intss2Points(points [][]int) []Point {
	res := make([]Point, len(points))
	for i, p := range points {
		res[i] = Point{
			X: p[0],
			Y: p[1],
		}
	}
	return res
}

// Points2Intss 把 []Point 转换成　[][]int
func Points2Intss(points []Point) [][]int {
	res := make([][]int, len(points))
	for i, p := range points {
		res[i] = []int{p.X, p.Y}
	}
	return res
}
