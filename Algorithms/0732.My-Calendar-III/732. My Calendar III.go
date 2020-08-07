package leetcode

// SegmentTree732 define
type SegmentTree732 struct {
	start, end, count int
	left, right       *SegmentTree732
}

// MyCalendarThree define
type MyCalendarThree struct {
	st        *SegmentTree732
	maxHeight int
}

// Constructor732 define
func Constructor732() MyCalendarThree {
	st := &SegmentTree732{
		start: 0,
		end:   1e9,
	}
	return MyCalendarThree{
		st: st,
	}
}

// Book define
func (mct *MyCalendarThree) Book(start int, end int) int {
	mct.st.book(start, end, &mct.maxHeight)
	return mct.maxHeight
}

func (st *SegmentTree732) book(start, end int, maxHeight *int) {
	if start == end {
		return
	}
	if start == st.start && st.end == end {
		st.count++
		if st.count > *maxHeight {
			*maxHeight = st.count
		}
		if st.left == nil {
			return
		}
	}
	if st.left == nil {
		if start == st.start {
			st.left = &SegmentTree732{start: start, end: end, count: st.count}
			st.right = &SegmentTree732{start: end, end: st.end, count: st.count}
			st.left.book(start, end, maxHeight)
			return
		}
		st.left = &SegmentTree732{start: st.start, end: start, count: st.count}
		st.right = &SegmentTree732{start: start, end: st.end, count: st.count}
		st.right.book(start, end, maxHeight)
		return
	}
	if start >= st.right.start {
		st.right.book(start, end, maxHeight)
	} else if end <= st.left.end {
		st.left.book(start, end, maxHeight)
	} else {
		st.left.book(start, st.left.end, maxHeight)
		st.right.book(st.right.start, end, maxHeight)
	}
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
