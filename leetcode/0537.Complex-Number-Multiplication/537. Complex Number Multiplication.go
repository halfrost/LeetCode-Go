package leetcode

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	realA, imagA := parse(a)
	realB, imagB := parse(b)
	real := realA*realB - imagA*imagB
	imag := realA*imagB + realB*imagA
	return strconv.Itoa(real) + "+" + strconv.Itoa(imag) + "i"
}

func parse(s string) (int, int) {
	ss := strings.Split(s, "+")
	r, _ := strconv.Atoi(ss[0])
	i, _ := strconv.Atoi(ss[1][:len(ss[1])-1])
	return r, i
}
