package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem676(t *testing.T) {
	dict := []string{"hello", "leetcode"}
	obj := Constructor676()
	obj.BuildDict(dict)
	fmt.Printf("obj = %v\n", obj)
	fmt.Println(obj.Search("hello"))
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("leetcode"))
	fmt.Println(obj.Search("leetcoded"))
	fmt.Println(obj.Search("hhllo"))
	fmt.Println(obj.Search("hell"))
}
