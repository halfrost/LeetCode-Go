package leetcode

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	clockwiseDis, counterclockwiseDis, n := 0, 0, len(distance)
	for i := start; i != destination; i = (i + 1) % n {
		clockwiseDis += distance[i]
	}
	for i := destination; i != start; i = (i + 1) % n {
		counterclockwiseDis += distance[i]
	}
	if clockwiseDis < counterclockwiseDis {
		return clockwiseDis
	}
	return counterclockwiseDis
}
