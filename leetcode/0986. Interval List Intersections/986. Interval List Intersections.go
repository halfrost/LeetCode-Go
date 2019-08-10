package leetcode

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func intervalIntersection(A []Interval, B []Interval) []Interval {
	res := []Interval{}
	for i, j := 0, 0; i < len(A) && j < len(B); {
		start := max(A[i].Start, B[j].Start)
		end := min(A[i].End, B[j].End)
		if start <= end {
			res = append(res, Interval{Start: start, End: end})
		}
		if A[i].End <= B[j].End {
			i++
		} else {
			j++
		}
	}
	return res
}
