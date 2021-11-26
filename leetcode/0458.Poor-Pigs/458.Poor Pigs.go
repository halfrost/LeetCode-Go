package leetcode

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	base := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log10(float64(buckets)) / math.Log10(float64(base))))
}
