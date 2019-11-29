package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem981(t *testing.T) {
	obj := Constructor981()
	obj.Set("foo", "bar", 1)
	fmt.Printf("Get = %v\n", obj.Get("foo", 1))
	fmt.Printf("Get = %v\n", obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Printf("Get = %v\n", obj.Get("foo", 4))
	fmt.Printf("Get = %v\n", obj.Get("foo", 5))
}
