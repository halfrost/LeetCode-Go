package leetcode

func countBalls(lowLimit int, highLimit int) int {
	buckets, maxBall := [46]int{}, 0
	for i := lowLimit; i <= highLimit; i++ {
		t := 0
		for j := i; j > 0; {
			t += j % 10
			j = j / 10
		}
		buckets[t]++
		if buckets[t] > maxBall {
			maxBall = buckets[t]
		}
	}
	return maxBall
}
