package leetcode

import "sort"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	return sort.Search(n, func(x int) bool { return guess(x) <= 0 })
}

func guess(num int) int {
	return 0
}
