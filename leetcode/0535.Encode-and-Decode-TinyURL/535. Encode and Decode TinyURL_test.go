package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem535(t *testing.T) {
	obj := Constructor()
	fmt.Printf("obj = %v\n", obj)
	e := obj.encode("https://leetcode.com/problems/design-tinyurl")
	fmt.Printf("obj encode = %v\n", e)
	d := obj.decode(e)
	fmt.Printf("obj decode = %v\n", d)
}
