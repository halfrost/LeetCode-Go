package leetcode

// RangeModule define
type RangeModule struct {
	Root *SegmentTreeNode
}

// SegmentTreeNode define
type SegmentTreeNode struct {
	Start, End  int
	Tracked     bool
	Lazy        int
	Left, Right *SegmentTreeNode
}

// Constructor715 define
func Constructor715() RangeModule {
	return RangeModule{&SegmentTreeNode{0, 1e9, false, 0, nil, nil}}
}

// AddRange define
func (rm *RangeModule) AddRange(left int, right int) {
	update(rm.Root, left, right-1, true)
}

// QueryRange define
func (rm *RangeModule) QueryRange(left int, right int) bool {
	return query(rm.Root, left, right-1)
}

// RemoveRange define
func (rm *RangeModule) RemoveRange(left int, right int) {
	update(rm.Root, left, right-1, false)
}

func lazyUpdate(node *SegmentTreeNode) {
	if node.Lazy != 0 {
		node.Tracked = node.Lazy == 2
	}
	if node.Start != node.End {
		if node.Left == nil || node.Right == nil {
			m := node.Start + (node.End-node.Start)/2
			node.Left = &SegmentTreeNode{node.Start, m, node.Tracked, 0, nil, nil}
			node.Right = &SegmentTreeNode{m + 1, node.End, node.Tracked, 0, nil, nil}
		} else if node.Lazy != 0 {
			node.Left.Lazy = node.Lazy
			node.Right.Lazy = node.Lazy
		}
	}
	node.Lazy = 0
}

func update(node *SegmentTreeNode, start, end int, track bool) {
	lazyUpdate(node)
	if start > end || node == nil || end < node.Start || node.End < start {
		return
	}
	if start <= node.Start && node.End <= end {
		// segment completely covered by the update range
		node.Tracked = track
		if node.Start != node.End {
			if track {
				node.Left.Lazy = 2
				node.Right.Lazy = 2
			} else {
				node.Left.Lazy = 1
				node.Right.Lazy = 1
			}
		}
		return
	}
	update(node.Left, start, end, track)
	update(node.Right, start, end, track)
	node.Tracked = node.Left.Tracked && node.Right.Tracked
}

func query(node *SegmentTreeNode, start, end int) bool {
	lazyUpdate(node)
	if start > end || node == nil || end < node.Start || node.End < start {
		return true
	}
	if start <= node.Start && node.End <= end {
		// segment completely covered by the update range
		return node.Tracked
	}
	return query(node.Left, start, end) && query(node.Right, start, end)
}

// 解法二 BST
// type RangeModule struct {
// 	Root *BSTNode
// }

// type BSTNode struct {
// 	Interval    []int
// 	Left, Right *BSTNode
// }

// func Constructor715() RangeModule {
// 	return RangeModule{}
// }

// func (this *RangeModule) AddRange(left int, right int) {
// 	interval := []int{left, right - 1}
// 	this.Root = insert(this.Root, interval)
// }

// func (this *RangeModule) RemoveRange(left int, right int) {
// 	interval := []int{left, right - 1}
// 	this.Root = delete(this.Root, interval)
// }

// func (this *RangeModule) QueryRange(left int, right int) bool {
// 	return query(this.Root, []int{left, right - 1})
// }

// func (this *RangeModule) insert(root *BSTNode, interval []int) *BSTNode {
// 	if root == nil {
// 		return &BSTNode{interval, nil, nil}
// 	}
// 	if root.Interval[0] <= interval[0] && interval[1] <= root.Interval[1] {
// 		return root
// 	}
// 	if interval[0] < root.Interval[0] {
// 		root.Left = insert(root.Left, []int{interval[0], min(interval[1], root.Interval[0]-1)})
// 	}
// 	if root.Interval[1] < interval[1] {
// 		root.Right = insert(root.Right, []int{max(interval[0], root.Interval[1]+1), interval[1]})
// 	}
// 	return root
// }

// func (this *RangeModule) delete(root *BSTNode, interval []int) *BSTNode {
// 	if root == nil {
// 		return nil
// 	}
// 	if interval[0] < root.Interval[0] {
// 		root.Left = delete(root.Left, []int{interval[0], min(interval[1], root.Interval[0]-1)})
// 	}
// 	if root.Interval[1] < interval[1] {
// 		root.Right = delete(root.Right, []int{max(interval[0], root.Interval[1]+1), interval[1]})
// 	}
// 	if interval[1] < root.Interval[0] || root.Interval[1] < interval[0] {
// 		return root
// 	}
// 	if interval[0] <= root.Interval[0] && root.Interval[1] <= interval[1] {
// 		if root.Left == nil {
// 			return root.Right
// 		} else if root.Right == nil {
// 			return root.Left
// 		} else {
// 			pred := root.Left
// 			for pred.Right != nil {
// 				pred = pred.Right
// 			}
// 			root.Interval = pred.Interval
// 			root.Left = delete(root.Left, pred.Interval)
// 			return root
// 		}
// 	}
// 	if root.Interval[0] < interval[0] && interval[1] < root.Interval[1] {
// 		left := &BSTNode{[]int{root.Interval[0], interval[0] - 1}, root.Left, nil}
// 		right := &BSTNode{[]int{interval[1] + 1, root.Interval[1]}, nil, root.Right}
// 		left.Right = right
// 		return left
// 	}
// 	if interval[0] <= root.Interval[0] {
// 		root.Interval[0] = interval[1] + 1
// 	}
// 	if root.Interval[1] <= interval[1] {
// 		root.Interval[1] = interval[0] - 1
// 	}
// 	return root
// }

// func (this *RangeModule) query(root *BSTNode, interval []int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	if interval[1] < root.Interval[0] {
// 		return query(root.Left, interval)
// 	}
// 	if root.Interval[1] < interval[0] {
// 		return query(root.Right, interval)
// 	}
// 	left := true
// 	if interval[0] < root.Interval[0] {
// 		left = query(root.Left, []int{interval[0], root.Interval[0] - 1})
// 	}
// 	right := true
// 	if root.Interval[1] < interval[1] {
// 		right = query(root.Right, []int{root.Interval[1] + 1, interval[1]})
// 	}
// 	return left && right
// }

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
