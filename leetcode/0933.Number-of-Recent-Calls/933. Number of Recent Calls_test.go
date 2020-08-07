package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem933(t *testing.T) {
	obj := Constructor933()
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.Ping(1)
	fmt.Printf("param = %v\n", param1)
	param1 = obj.Ping(100)
	fmt.Printf("param = %v\n", param1)
	param1 = obj.Ping(3001)
	fmt.Printf("param = %v\n", param1)
	param1 = obj.Ping(3002)
	fmt.Printf("param = %v\n", param1)
}
