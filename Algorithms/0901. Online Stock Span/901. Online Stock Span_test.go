package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem901(t *testing.T) {
	obj := Constructor901()
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Next(100)
	fmt.Printf("param_1 = %v\n", param1)
	param2 := obj.Next(80)
	fmt.Printf("param_2 = %v\n", param2)
	param3 := obj.Next(60)
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj.Next(70)
	fmt.Printf("param_4 = %v\n", param4)
	param5 := obj.Next(60)
	fmt.Printf("param_5 = %v\n", param5)
	param6 := obj.Next(75)
	fmt.Printf("param_6 = %v\n", param6)
	param7 := obj.Next(85)
	fmt.Printf("param_7 = %v\n", param7)
}
