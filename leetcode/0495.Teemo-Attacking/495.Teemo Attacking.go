package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	var ans int
	for i := 1; i < len(timeSeries); i++ {
		t := timeSeries[i-1]
		end := t + duration - 1
		if end < timeSeries[i] {
			ans += duration
		} else {
			ans += timeSeries[i] - t
		}
	}
	ans += duration
	return ans
}
