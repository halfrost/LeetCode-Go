package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem208(t *testing.T) {
	obj := Constructor_208()
	fmt.Printf("obj = %v\n", obj)
	obj.Insert("apple")
	fmt.Printf("obj = %v\n", obj)
	param_1 := obj.Search("apple")
	fmt.Printf("param_1 = %v obj = %v\n", param_1, obj)
	param_2 := obj.Search("app")
	fmt.Printf("param_2 = %v obj = %v\n", param_2, obj)
	param_3 := obj.StartsWith("app")
	fmt.Printf("param_3 = %v obj = %v\n", param_3, obj)
	obj.Insert("app")
	fmt.Printf("obj = %v\n", obj)
	param_4 := obj.Search("app")
	fmt.Printf("param_4 = %v obj = %v\n", param_4, obj)
}
