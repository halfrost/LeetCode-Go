package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1600(t *testing.T) {
	obj := Constructor("king")
	fmt.Printf("obj = %v\n", obj)
	obj.Birth("king", "andy")
	fmt.Printf("obj = %v\n", obj)
	obj.Birth("king", "bob")
	fmt.Printf("obj = %v\n", obj)
	obj.Birth("king", "catherine")
	fmt.Printf("obj = %v\n", obj)
	obj.Birth("andy", "matthew")
	fmt.Printf("obj = %v\n", obj)
	obj.Birth("bob", "alex")
	fmt.Printf("obj = %v\n", obj)
	obj.Birth("bob", "asha")
	fmt.Printf("obj = %v\n", obj)
	param2 := obj.GetInheritanceOrder()
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
	obj.Death("bob")
	fmt.Printf("obj = %v\n", obj)
	param2 = obj.GetInheritanceOrder()
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
}
