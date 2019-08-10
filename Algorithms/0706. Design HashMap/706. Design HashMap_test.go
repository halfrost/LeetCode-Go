package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem706(t *testing.T) {
	obj := Constructor706()
	obj.Put(7, 10)
	fmt.Printf("Get 7 = %v\n", obj.Get(7))
	obj.Put(7, 20)
	fmt.Printf("Contains 7 = %v\n", obj.Get(7))
	param1 := obj.Get(100)
	fmt.Printf("param1 = %v\n", param1)
	obj.Remove(7)
	param1 = obj.Get(7)
	fmt.Printf("param1 = %v\n", param1)
}
