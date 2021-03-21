package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1603(t *testing.T) {
	obj := Constructor(1, 1, 0)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj = %v\n", obj.AddCar(1))
	fmt.Printf("obj = %v\n", obj.AddCar(2))
	fmt.Printf("obj = %v\n", obj.AddCar(3))
	fmt.Printf("obj = %v\n", obj.AddCar(1))
}
