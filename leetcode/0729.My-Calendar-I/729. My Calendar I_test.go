package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem729(t *testing.T) {
	obj := Constructor729()
	param1 := obj.Book(10, 20)
	fmt.Printf("param = %v obj = %v\n", param1, obj)
	param1 = obj.Book(15, 25)
	fmt.Printf("param = %v obj = %v\n", param1, obj)
	param1 = obj.Book(20, 30)
	fmt.Printf("param = %v obj = %v\n", param1, obj)

	obj1 := Constructor729()
	param2 := obj1.Book(47, 50)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(33, 41)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(39, 45)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(33, 42)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(25, 32)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(26, 35)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(19, 25)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(3, 8)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(8, 13)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
	param2 = obj1.Book(18, 27)
	fmt.Printf("param = %v obj = %v\n", param2, obj1)
}
