package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem173(t *testing.T) {

	root := Ints2TreeNode([]int{9, 7, 15, 3, NULL, NULL, 20})
	obj := Constructor173(root)
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param2 := obj.HasNext()
	fmt.Printf("param_2 = %v\n", param2)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param1)
	param2 = obj.HasNext()
	fmt.Printf("param_2 = %v\n", param2)
}
