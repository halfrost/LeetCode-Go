package leetcode

import (
	"fmt"
	"testing"
)

var tcs = []struct {
	n   int
	ans []string
}{

	{
		15,
		[]string{
			"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",
		},
	},
}

func Test_fizzBuzz(t *testing.T) {
	fmt.Printf("------------------------Leetcode Problem 412------------------------\n")
	for _, tc := range tcs {
		fmt.Printf("【output】:%v\n", tc)
	}
	fmt.Printf("\n\n\n")
}

func Benchmark_fizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fizzBuzz(tc.n)
		}
	}
}
