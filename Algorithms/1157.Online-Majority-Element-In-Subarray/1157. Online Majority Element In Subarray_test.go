package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1157(t *testing.T) {
	arr := []int{1, 1, 2, 2, 1, 1}
	obj := Constructor1157(arr)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("query(0,5,4) = %v\n", obj.Query(0, 5, 4)) //1
	fmt.Printf("query(0,3,3) = %v\n", obj.Query(0, 3, 3)) //-1
	fmt.Printf("query(2,3,2) = %v\n", obj.Query(2, 3, 2)) //2
}
