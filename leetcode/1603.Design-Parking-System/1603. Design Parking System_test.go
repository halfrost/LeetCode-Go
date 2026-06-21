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

	obj2 := Constructor(0, 0, 1)
	fmt.Printf("obj2 = %v\n", obj2)
	fmt.Printf("obj2 = %v\n", obj2.AddCar(2)) // medium full -> false
	fmt.Printf("obj2 = %v\n", obj2.AddCar(3)) // small available -> true
	fmt.Printf("obj2 = %v\n", obj2.AddCar(3)) // small full -> false
	fmt.Printf("obj2 = %v\n", obj2.AddCar(4)) // invalid carType -> false
}
