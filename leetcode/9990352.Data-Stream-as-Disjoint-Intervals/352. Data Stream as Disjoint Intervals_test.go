package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem352(t *testing.T) {
	obj := Constructor352()
	fmt.Printf("obj = %v\n", obj)
	obj.AddNum(1)
	fmt.Printf("Intervals = %v\n", obj.GetIntervals()) // [1,1]
	obj.AddNum(3)
	fmt.Printf("Intervals = %v\n", obj.GetIntervals()) // [1,1] [3,3]
	obj.AddNum(7)
	fmt.Printf("Intervals = %v\n", obj.GetIntervals()) // [1, 1], [3, 3], [7, 7]
	obj.AddNum(2)
	fmt.Printf("Intervals = %v\n", obj.GetIntervals()) // [1, 3], [7, 7]
	obj.AddNum(6)
	fmt.Printf("Intervals = %v\n", obj.GetIntervals()) // [1, 3], [6, 7]

	// Cover remaining branches.
	obj2 := Constructor352()
	obj2.AddNum(5)                                      // nil init -> [5,5]
	obj2.AddNum(5)                                      // already inside an interval (early return)
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [5,5]
	obj2.AddNum(4)                                      // low==0 and intervals[0].Start-1==val -> [4,5]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [4,5]
	obj2.AddNum(1)                                      // low==0 prepend new interval -> [1,1],[4,5]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[4,5]
	obj2.AddNum(10)                                     // low==len append new interval -> [1,1],[4,5],[10,10]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[4,5],[10,10]
	obj2.AddNum(11)                                     // low==len End+1==val -> extend -> [10,11]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[4,5],[10,11]
	obj2.AddNum(7)                                      // gap insert in middle -> [1,1],[4,5],[7,7],[10,11]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[4,5],[7,7],[10,11]
	obj2.AddNum(6)                                      // merge both sides: [4,5]+6+[7,7] -> [4,7]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[4,7],[10,11]
	obj2.AddNum(8)                                      // intervals[low-1].End==val-1 only (extend left end) -> [4,8]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[4,8],[10,11]
	obj2.AddNum(3)                                      // else branch: intervals[low].Start-- -> [3,8]
	fmt.Printf("Intervals = %v\n", obj2.GetIntervals()) // [1,1],[3,8],[10,11]

	got := obj2.GetIntervals()
	want := [][]int{{1, 1}, {3, 8}, {10, 11}}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
