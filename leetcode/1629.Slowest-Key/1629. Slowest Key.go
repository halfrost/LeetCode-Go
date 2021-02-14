package leetcode

func slowestKey(releaseTimes []int, keysPressed string) byte {
	longestDuration, key := releaseTimes[0], keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > longestDuration {
			longestDuration = duration
			key = keysPressed[i]
		} else if duration == longestDuration && keysPressed[i] > key {
			key = keysPressed[i]
		}
	}
	return key
}
