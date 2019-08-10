package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem705(t *testing.T) {
	obj := Constructor705()
	obj.Add(7)
	fmt.Printf("Contains 7 = %v\n", obj.Contains(7))
	obj.Remove(10)
	fmt.Printf("Contains 10 = %v\n", obj.Contains(10))
	obj.Add(20)
	fmt.Printf("Contains 20 = %v\n", obj.Contains(20))
	obj.Remove(30)
	fmt.Printf("Contains 30 = %v\n", obj.Contains(30))
	obj.Add(8)
	fmt.Printf("Contains 8 = %v\n", obj.Contains(8))
	obj.Remove(8)
	fmt.Printf("Contains 8 = %v\n", obj.Contains(8))
	param1 := obj.Contains(7)
	fmt.Printf("param1 = %v\n", param1)
}
