package leetcode

// 解法一 二叉排序树
// Event define
type Event struct {
	start, end  int
	left, right *Event
}

// Insert define
func (e *Event) Insert(curr *Event) bool {
	if e.end > curr.start && curr.end > e.start {
		return false
	}
	if curr.start < e.start {
		if e.left == nil {
			e.left = curr
		} else {
			return e.left.Insert(curr)
		}
	} else {
		if e.right == nil {
			e.right = curr
		} else {
			return e.right.Insert(curr)
		}
	}
	return true
}

// MyCalendar define
type MyCalendar struct {
	root *Event
}

// Constructor729 define
func Constructor729() MyCalendar {
	return MyCalendar{
		root: nil,
	}
}

// Book define
func (this *MyCalendar) Book(start int, end int) bool {
	curr := &Event{start: start, end: end, left: nil, right: nil}
	if this.root == nil {
		this.root = curr
		return true
	}
	return this.root.Insert(curr)
}

// 解法二 快排 + 二分
// MyCalendar define
// type MyCalendar struct {
// 	calendar []Interval
// }

// // Constructor729 define
// func Constructor729() MyCalendar {
// 	calendar := []Interval{}
// 	return MyCalendar{calendar: calendar}
// }

// // Book define
// func (this *MyCalendar) Book(start int, end int) bool {
// 	if len(this.calendar) == 0 {
// 		this.calendar = append(this.calendar, Interval{Start: start, End: end})
// 		return true
// 	}
// 	// 快排
// 	quickSort(this.calendar, 0, len(this.calendar)-1)
// 	// 二分
// 	pos := searchLastLessInterval(this.calendar, start, end)
// 	// 如果找到最后一个元素，需要判断 end
// 	if pos == len(this.calendar)-1 && this.calendar[pos].End <= start {
// 		this.calendar = append(this.calendar, Interval{Start: start, End: end})
// 		return true
// 	}
// 	// 如果不是开头和结尾的元素，还需要判断这个区间是否能插入到原数组中(要看起点和终点是否都能插入)
// 	if pos != len(this.calendar)-1 && pos != -1 && this.calendar[pos].End <= start && this.calendar[pos+1].Start >= end {
// 		this.calendar = append(this.calendar, Interval{Start: start, End: end})
// 		return true
// 	}
// 	// 如果元素比开头的元素还要小，要插入到开头
// 	if this.calendar[0].Start >= end {
// 		this.calendar = append(this.calendar, Interval{Start: start, End: end})
// 		return true
// 	}
// 	return false
// }

// func searchLastLessInterval(intervals []Interval, start, end int) int {
// 	low, high := 0, len(intervals)-1
// 	for low <= high {
// 		mid := low + ((high - low) >> 1)
// 		if intervals[mid].Start <= start {
// 			if (mid == len(intervals)-1) || (intervals[mid+1].Start > start) { // 找到最后一个小于等于 target 的元素
// 				return mid
// 			}
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}
// 	return -1
// }

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
