package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem65(t *testing.T) {

	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0",
			true,
		},

		{
			"e",
			false,
		},

		{
			".",
			false,
		},

		{
			".1",
			true,
		},
	}
	fmt.Printf("------------------------Leetcode Problem 65------------------------\n")

	for _, tc := range tcs {
		fmt.Printf("【input】:%v       【output】:%v\n", tc, isNumber(tc.s))
	}
	fmt.Printf("\n\n\n")
}
