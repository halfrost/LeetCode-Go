package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem208(t *testing.T) {
	obj := Constructor208()
	fmt.Printf("obj = %v\n", obj)
	obj.Insert("apple")
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Search("apple")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.Search("app")
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
	param3 := obj.StartsWith("app")
	fmt.Printf("param_3 = %v obj = %v\n", param3, obj)
	obj.Insert("app")
	fmt.Printf("obj = %v\n", obj)
	param4 := obj.Search("app")
	fmt.Printf("param_4 = %v obj = %v\n", param4, obj)
}
