package leetcode

import "math/rand"

func rand10() int {
	rand10 := 10
	for rand10 >= 10 {
		rand10 = (rand7() - 1) + rand7()
	}
	return rand10%10 + 1
}

func rand7() int {
	return rand.Intn(7)
}

func rand101() int {
	rand40 := 40
	for rand40 >= 40 {
		rand40 = (rand7()-1)*7 + rand7() - 1
	}
	return rand40%10 + 1
}
