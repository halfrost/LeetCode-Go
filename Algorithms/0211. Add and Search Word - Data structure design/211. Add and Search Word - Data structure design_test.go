package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem211(t *testing.T) {
	obj := Constructor_211()
	fmt.Printf("obj = %v\n", obj)
	obj.AddWord("bad")
	fmt.Printf("obj = %v\n", obj)
	obj.AddWord("dad")
	fmt.Printf("obj = %v\n", obj)
	obj.AddWord("mad")
	fmt.Printf("obj = %v\n", obj)

	param_1 := obj.Search("pad")
	fmt.Printf("param_1 = %v obj = %v\n", param_1, obj)
	param_2 := obj.Search("bad")
	fmt.Printf("param_2 = %v obj = %v\n", param_2, obj)
	param_3 := obj.Search(".ad")
	fmt.Printf("param_3 = %v obj = %v\n", param_3, obj)
	param_4 := obj.Search("b..")
	fmt.Printf("param_4 = %v obj = %v\n", param_4, obj)
}
