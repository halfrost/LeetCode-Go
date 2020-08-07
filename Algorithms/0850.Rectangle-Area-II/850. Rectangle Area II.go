package leetcode

import (
	"sort"
)

func rectangleArea(rectangles [][]int) int {
	sat, res := SegmentAreaTree{}, 0
	posXMap, posX, posYMap, posY, lines := discretization850(rectangles)
	tmp := make([]int, len(posYMap))
	for i := 0; i < len(tmp)-1; i++ {
		tmp[i] = posY[i+1] - posY[i]
	}
	sat.Init(tmp, func(i, j int) int {
		return i + j
	})
	for i := 0; i < len(posY)-1; i++ {
		tmp[i] = posY[i+1] - posY[i]
	}
	for i := 0; i < len(posX)-1; i++ {
		for _, v := range lines[posXMap[posX[i]]] {
			sat.Update(posYMap[v.start], posYMap[v.end], v.state)
		}
		res += ((posX[i+1] - posX[i]) * sat.Query(0, len(posY)-1)) % 1000000007
	}
	return res % 1000000007
}

func discretization850(positions [][]int) (map[int]int, []int, map[int]int, []int, map[int][]LineItem) {
	tmpXMap, tmpYMap, posXArray, posXMap, posYArray, posYMap, lines := map[int]int{}, map[int]int{}, []int{}, map[int]int{}, []int{}, map[int]int{}, map[int][]LineItem{}
	for _, pos := range positions {
		tmpXMap[pos[0]]++
		tmpXMap[pos[2]]++
	}
	for k := range tmpXMap {
		posXArray = append(posXArray, k)
	}
	sort.Ints(posXArray)
	for i, pos := range posXArray {
		posXMap[pos] = i
	}

	for _, pos := range positions {
		tmpYMap[pos[1]]++
		tmpYMap[pos[3]]++
		tmp1 := lines[posXMap[pos[0]]]
		tmp1 = append(tmp1, LineItem{start: pos[1], end: pos[3], state: 1})
		lines[posXMap[pos[0]]] = tmp1
		tmp2 := lines[posXMap[pos[2]]]
		tmp2 = append(tmp2, LineItem{start: pos[1], end: pos[3], state: -1})
		lines[posXMap[pos[2]]] = tmp2
	}
	for k := range tmpYMap {
		posYArray = append(posYArray, k)
	}
	sort.Ints(posYArray)
	for i, pos := range posYArray {
		posYMap[pos] = i
	}
	return posXMap, posXArray, posYMap, posYArray, lines
}

// LineItem define
type LineItem struct { // 垂直于 x 轴的线段
	start, end, state int // state = 1 代表进入，-1 代表离开
}

// SegmentItem define
type SegmentItem struct {
	count int
	val   int
}

// SegmentAreaTree define
type SegmentAreaTree struct {
	data        []int
	tree        []SegmentItem
	left, right int
	merge       func(i, j int) int
}

// Init define
func (sat *SegmentAreaTree) Init(nums []int, oper func(i, j int) int) {
	sat.merge = oper
	data, tree := make([]int, len(nums)), make([]SegmentItem, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	sat.data, sat.tree = data, tree
	if len(nums) > 0 {
		sat.buildSegmentTree(0, 0, len(nums)-1)
	}
}

// 在 treeIndex 的位置创建 [left....right] 区间的线段树
func (sat *SegmentAreaTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right-1 {
		sat.tree[treeIndex] = SegmentItem{count: 0, val: sat.data[left]}
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, sat.leftChild(treeIndex), sat.rightChild(treeIndex)
	sat.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	sat.buildSegmentTree(rightTreeIndex, midTreeIndex, right)
	sat.pushUp(treeIndex, leftTreeIndex, rightTreeIndex)
}

func (sat *SegmentAreaTree) pushUp(treeIndex, leftTreeIndex, rightTreeIndex int) {
	newCount, newValue := sat.merge(sat.tree[leftTreeIndex].count, sat.tree[rightTreeIndex].count), 0
	if sat.tree[leftTreeIndex].count > 0 && sat.tree[rightTreeIndex].count > 0 {
		newValue = sat.merge(sat.tree[leftTreeIndex].val, sat.tree[rightTreeIndex].val)
	} else if sat.tree[leftTreeIndex].count > 0 && sat.tree[rightTreeIndex].count == 0 {
		newValue = sat.tree[leftTreeIndex].val
	} else if sat.tree[leftTreeIndex].count == 0 && sat.tree[rightTreeIndex].count > 0 {
		newValue = sat.tree[rightTreeIndex].val
	}
	sat.tree[treeIndex] = SegmentItem{count: newCount, val: newValue}
}

func (sat *SegmentAreaTree) leftChild(index int) int {
	return 2*index + 1
}

func (sat *SegmentAreaTree) rightChild(index int) int {
	return 2*index + 2
}

// 查询 [left....right] 区间内的值

// Query define
func (sat *SegmentAreaTree) Query(left, right int) int {
	if len(sat.data) > 0 {
		return sat.queryInTree(0, 0, len(sat.data)-1, left, right)
	}
	return 0
}

func (sat *SegmentAreaTree) queryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, sat.leftChild(treeIndex), sat.rightChild(treeIndex)
	if left > queryRight || right < queryLeft { // segment completely outside range
		return 0 // represents a null node
	}
	if queryLeft <= left && queryRight >= right { // segment completely inside range
		if sat.tree[treeIndex].count > 0 {
			return sat.tree[treeIndex].val
		}
		return 0
	}
	if queryLeft > midTreeIndex {
		return sat.queryInTree(rightTreeIndex, midTreeIndex, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return sat.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	// merge query results
	return sat.merge(sat.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		sat.queryInTree(rightTreeIndex, midTreeIndex, right, midTreeIndex, queryRight))
}

// Update define
func (sat *SegmentAreaTree) Update(updateLeft, updateRight, val int) {
	if len(sat.data) > 0 {
		sat.updateInTree(0, 0, len(sat.data)-1, updateLeft, updateRight, val)
	}
}

func (sat *SegmentAreaTree) updateInTree(treeIndex, left, right, updateLeft, updateRight, val int) {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, sat.leftChild(treeIndex), sat.rightChild(treeIndex)
	if left > right || left >= updateRight || right <= updateLeft { // 由于叶子节点的区间不在是 left == right 所以这里判断需要增加等号的判断
		return // out of range. escape.
	}

	if updateLeft <= left && right <= updateRight { // segment is fully within update range
		if left == right-1 {
			sat.tree[treeIndex].count = sat.merge(sat.tree[treeIndex].count, val)
		}
		if left != right-1 { // update lazy[] for children
			sat.updateInTree(leftTreeIndex, left, midTreeIndex, updateLeft, updateRight, val)
			sat.updateInTree(rightTreeIndex, midTreeIndex, right, updateLeft, updateRight, val)
			sat.pushUp(treeIndex, leftTreeIndex, rightTreeIndex)
		}
		return
	}
	sat.updateInTree(leftTreeIndex, left, midTreeIndex, updateLeft, updateRight, val)
	sat.updateInTree(rightTreeIndex, midTreeIndex, right, updateLeft, updateRight, val)
	// merge updates
	sat.pushUp(treeIndex, leftTreeIndex, rightTreeIndex)
}
