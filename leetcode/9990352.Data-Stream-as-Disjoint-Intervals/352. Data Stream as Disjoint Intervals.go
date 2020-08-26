package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// Interval define
type Interval = structures.Interval

// SummaryRanges define
type SummaryRanges struct {
	intervals []Interval
}

// Constructor352 define
func Constructor352() SummaryRanges {
	return SummaryRanges{}
}

// AddNum define
func (sr *SummaryRanges) AddNum(val int) {
	if sr.intervals == nil {
		sr.intervals = []Interval{
			{
				Start: val,
				End:   val,
			},
		}
		return
	}

	low, high := 0, len(sr.intervals)-1
	for low <= high {
		mid := low + (high-low)>>1
		if sr.intervals[mid].Start <= val && val <= sr.intervals[mid].End {
			return
		} else if val < sr.intervals[mid].Start {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low == 0 {
		if sr.intervals[0].Start-1 == val {
			sr.intervals[0].Start--
			return
		}
		ni := Interval{Start: val, End: val}
		sr.intervals = append(sr.intervals, ni)
		copy(sr.intervals[1:], sr.intervals)
		sr.intervals[0] = ni
		return
	}

	if low == len(sr.intervals) {
		if sr.intervals[low-1].End+1 == val {
			sr.intervals[low-1].End++
			return
		}
		sr.intervals = append(sr.intervals, Interval{Start: val, End: val})
		return
	}

	if sr.intervals[low-1].End+1 < val && val < sr.intervals[low].Start-1 {
		sr.intervals = append(sr.intervals, Interval{})
		copy(sr.intervals[low+1:], sr.intervals[low:])
		sr.intervals[low] = Interval{Start: val, End: val}
		return
	}

	if sr.intervals[low-1].End == val-1 && val+1 == sr.intervals[low].Start {
		sr.intervals[low-1].End = sr.intervals[low].End
		n := len(sr.intervals)
		copy(sr.intervals[low:], sr.intervals[low+1:])
		sr.intervals = sr.intervals[:n-1]
		return
	}

	if sr.intervals[low-1].End == val-1 {
		sr.intervals[low-1].End++
	} else {
		sr.intervals[low].Start--
	}
}

// GetIntervals define
func (sr *SummaryRanges) GetIntervals() [][]int {
	intervals := [][]int{}
	for _, interval := range sr.intervals {
		intervals = append(intervals, []int{interval.Start, interval.End})
	}
	return intervals
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
