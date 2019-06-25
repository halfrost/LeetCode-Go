package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem173(t *testing.T) {

	root := Ints2TreeNode([]int{9, 7, 15, 3, NULL, NULL, 20})
	obj := Constructor_173(root)
	fmt.Printf("obj = %v\n", obj)
	param_1 := obj.Next()
	fmt.Printf("param_1 = %v\n", param_1)
	param_2 := obj.HasNext()
	fmt.Printf("param_2 = %v\n", param_2)
	param_1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param_1)
	param_1 = obj.Next()
	fmt.Printf("param_1 = %v\n", param_1)
	param_2 = obj.HasNext()
	fmt.Printf("param_2 = %v\n", param_2)
}
