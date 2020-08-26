package structures

// Interval 提供区间表示
type Interval struct {
	Start int
	End   int
}

// Interval2Ints 把 Interval 转换成 整型切片
func Interval2Ints(i Interval) []int {
	return []int{i.Start, i.End}
}

// IntervalSlice2Intss 把 []Interval 转换成 [][]int
func IntervalSlice2Intss(is []Interval) [][]int {
	res := make([][]int, 0, len(is))
	for i := range is {
		res = append(res, Interval2Ints(is[i]))
	}
	return res
}

// Intss2IntervalSlice 把 [][]int 转换成 []Interval
func Intss2IntervalSlice(intss [][]int) []Interval {
	res := make([]Interval, 0, len(intss))
	for _, ints := range intss {
		res = append(res, Interval{Start: ints[0], End: ints[1]})
	}
	return res
}

// QuickSort define
func QuickSort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}

func partitionSort(a []Interval, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
