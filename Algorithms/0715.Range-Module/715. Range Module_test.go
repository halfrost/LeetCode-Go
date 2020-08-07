package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem715(t *testing.T) {
	obj := Constructor715()
	obj.AddRange(10, 20)
	obj.RemoveRange(14, 16)
	fmt.Printf("query = %v\n", obj.QueryRange(10, 14))   // returns true
	fmt.Printf("query = %v\n", obj.QueryRange(13, 15))   // returns false
	fmt.Printf("query = %v\n\n", obj.QueryRange(16, 17)) // returns true

	obj1 := Constructor715()
	obj1.AddRange(10, 180)
	obj1.AddRange(150, 200)
	obj1.AddRange(250, 500)
	fmt.Printf("query = %v\n", obj1.QueryRange(50, 100))   // returns true
	fmt.Printf("query = %v\n", obj1.QueryRange(180, 300))  // returns false
	fmt.Printf("query = %v\n", obj1.QueryRange(600, 1000)) // returns false
	obj1.RemoveRange(50, 150)
	fmt.Printf("query = %v\n\n", obj1.QueryRange(50, 100)) // returns false

	obj2 := Constructor715()
	obj2.AddRange(6, 8)
	obj2.RemoveRange(7, 8)
	obj2.RemoveRange(8, 9)
	obj2.AddRange(8, 9)
	obj2.RemoveRange(1, 3)
	obj2.AddRange(1, 8)
	fmt.Printf("query = %v\n", obj2.QueryRange(2, 4)) // returns true
	fmt.Printf("query = %v\n", obj2.QueryRange(2, 9)) // returns true
	fmt.Printf("query = %v\n", obj2.QueryRange(4, 6)) // returns true
}
