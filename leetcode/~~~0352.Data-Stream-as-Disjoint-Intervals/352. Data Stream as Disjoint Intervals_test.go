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
}
