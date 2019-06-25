package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem901(t *testing.T) {
	obj := Constructor_901()
	fmt.Printf("obj = %v\n", obj)
	param_1 := obj.Next(100)
	fmt.Printf("param_1 = %v\n", param_1)
	param_2 := obj.Next(80)
	fmt.Printf("param_2 = %v\n", param_2)
	param_3 := obj.Next(60)
	fmt.Printf("param_3 = %v\n", param_3)
	param_4 := obj.Next(70)
	fmt.Printf("param_4 = %v\n", param_4)
	param_5 := obj.Next(60)
	fmt.Printf("param_5 = %v\n", param_5)
	param_6 := obj.Next(75)
	fmt.Printf("param_6 = %v\n", param_6)
	param_7 := obj.Next(85)
	fmt.Printf("param_7 = %v\n", param_7)
}
