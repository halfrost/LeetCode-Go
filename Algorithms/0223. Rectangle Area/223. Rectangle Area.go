package leetcode

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	X0, Y0, X1, Y1 := max(A, E), max(B, F), min(C, G), min(D, H)
	return area(A, B, C, D) + area(E, F, G, H) - area(X0, Y0, X1, Y1)
}

func area(x0, y0, x1, y1 int) int {
	l, h := x1-x0, y1-y0
	if l <= 0 || h <= 0 {
		return 0
	}
	return l * h
}
