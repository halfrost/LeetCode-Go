package leetcode

import (
	"sort"
)

type car struct {
	time     float64
	position int
}

// ByPosition define
type ByPosition []car

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].position > a[j].position }

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n <= 1 {
		return n
	}
	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{float64(target-position[i]) / float64(speed[i]), position[i]}
	}
	sort.Sort(ByPosition(cars))
	fleet, lastTime := 0, 0.0
	for i := 0; i < len(cars); i++ {
		if cars[i].time > lastTime {
			lastTime = cars[i].time
			fleet++
		}
	}
	return fleet
}
