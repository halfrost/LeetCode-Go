package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem732(t *testing.T) {
	obj := Constructor732()
	fmt.Printf("book = %v\n\n", obj.Book(10, 20)) // returns 1
	fmt.Printf("book = %v\n\n", obj.Book(50, 60)) // returns 1
	fmt.Printf("book = %v\n\n", obj.Book(10, 40)) // returns 2
	fmt.Printf("book = %v\n\n", obj.Book(5, 15))  // returns 3
	fmt.Printf("book = %v\n\n", obj.Book(5, 10))  // returns 3
	fmt.Printf("book = %v\n\n", obj.Book(25, 55)) // returns 3
}
