package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem497(t *testing.T) {
	w := [][]int{{1, 1, 5, 5}}
	sol := Constructor497(w)
	fmt.Printf("1.Pick = %v\n", sol.Pick())
	fmt.Printf("2.Pick = %v\n", sol.Pick())
	fmt.Printf("3.Pick = %v\n", sol.Pick())
	fmt.Printf("4.Pick = %v\n", sol.Pick())
	fmt.Printf("5.Pick = %v\n", sol.Pick())
	fmt.Printf("6.Pick = %v\n", sol.Pick())

	// Multiple rectangles to exercise the prefix-sum accumulation (i > 0)
	// and the binary-search branches inside Pick.
	w2 := [][]int{{-2, -2, -1, -1}, {1, 0, 3, 0}, {-2, 0, 0, 5}, {10, 10, 20, 20}}
	sol2 := Constructor497(w2)
	for i := 0; i < 200; i++ {
		p := sol2.Pick()
		if !inAnyRect(w2, p) {
			t.Fatalf("Pick returned %v which is not inside any rectangle %v", p, w2)
		}
	}

	// Rectangle whose computed (x2-x1+1)*(y2-y1+1) is negative, to cover
	// the area = -area normalization branch.
	w3 := [][]int{{0, 0, -3, 2}}
	sol3 := Constructor497(w3)
	if sol3.arr[0] <= 0 {
		t.Fatalf("expected positive normalized area, got %d", sol3.arr[0])
	}
}

func inAnyRect(rects [][]int, p []int) bool {
	for _, r := range rects {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		if p[0] >= x1 && p[0] <= x2 && p[1] >= y1 && p[1] <= y2 {
			return true
		}
	}
	return false
}
