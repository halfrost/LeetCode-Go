package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func Test_Problem478(t *testing.T) {
	obj := Constructor(1, 0, 0)
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())

	obj = Constructor(10, 5, -7.5)
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())
	fmt.Printf("RandPoint() = %v\n", obj.RandPoint())

	a := angle()
	if a < 0 || a >= 2*math.Pi {
		t.Fatalf("angle() = %v, want in [0, 2*Pi)", a)
	}
	fmt.Printf("angle() = %v\n", a)
}
