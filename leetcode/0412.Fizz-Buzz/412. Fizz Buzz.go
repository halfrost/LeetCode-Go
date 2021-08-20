package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	solution := make([]string, n)
	for i := 1; i <= n; i++ {
		solution[i-1] = ""
		if i%3 == 0 {
			solution[i-1] += "Fizz"
		}
		if i%5 == 0 {
			solution[i-1] += "Buzz"
		}
		if solution[i-1] == "" {
			solution[i-1] = strconv.Itoa(i)
		}
	}
	return solution
}
