package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1656(t *testing.T) {
	obj := Constructor(5)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Insert(3, "ccccc")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param1 = obj.Insert(1, "aaaaa")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param1 = obj.Insert(2, "bbbbb")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param1 = obj.Insert(5, "eeeee")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param1 = obj.Insert(4, "ddddd")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
}
