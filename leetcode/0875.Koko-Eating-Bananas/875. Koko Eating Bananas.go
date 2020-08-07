package leetcode

import "math"

func minEatingSpeed(piles []int, H int) int {
	low, high := 1, maxInArr(piles)
	for low < high {
		mid := low + (high-low)>>1
		if !isPossible(piles, mid, H) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func isPossible(piles []int, h, H int) bool {
	res := 0
	for _, p := range piles {
		res += int(math.Ceil(float64(p) / float64(h)))
	}
	return res <= H
}

func maxInArr(xs []int) int {
	res := 0
	for _, x := range xs {
		if res < x {
			res = x
		}
	}
	return res
}
