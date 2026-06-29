---
title: 3.1 Segment Tree
type: docs
weight: 1
---

# Segment Tree Segment Tree

A segment tree Segment tree is a binary tree data structure, invented by Jon Louis Bentley in 1977, used to store intervals or line segments and to allow fast queries for all intervals in the structure that contain a certain point.

For a segment tree containing {{< katex >}}n {{< /katex >}} intervals, the space complexity is {{< katex >}} O(n) {{< /katex >}}, and the query time complexity is {{< katex >}}O(log n+k) {{< /katex >}}, where {{< katex >}} k {{< /katex >}} is the number of intervals that meet the condition. The segment tree data structure can also be generalized to higher dimensions.

## I. What Is a Segment Tree

Take a one-dimensional segment tree as an example.

![](https://img.halfrost.com/Blog/ArticleImage/153_1.png)


Let S be a set of one-dimensional line segments. Sort the endpoint coordinates of these segments from small to large, and denote them as {{< katex >}}x_{1},x_{2},\cdots ,x_{m} {{< /katex >}}. We call each interval divided by these endpoints a "unit interval" (the position of each endpoint becomes a separate unit interval), including from left to right:

{{< katex display>}} 
(-\infty ,x_{1}),[x_{1},x_{1}],(x_{1},x_{2}),[x_{2},x_{2}],...,(x_{m-1},x_{m}),[x_{m},x_{m}],(x_{m},+\infty )
{{< /katex >}}

The structure of a segment tree is a binary tree. Each node represents a coordinate interval. The interval represented by node N is denoted as Int(N), and it must satisfy the following conditions:

- Each of its leaf nodes, from left to right, represents each unit interval.
- The interval represented by an internal node is the union of the intervals represented by its two children.
- Each node (including leaves) has a data structure that stores line segments. If the coordinate interval of a segment S contains Int(N) but does not contain Int(parent(N)), then segment S is stored in node N.


![](https://img.halfrost.com/Blog/ArticleImage/153_2.png)


A segment tree is a binary tree in which each node represents an interval. Usually, a node stores data for one or more merged intervals so that query operations can be performed.


## II. Why This Data Structure Is Needed

Many problems require us to produce results based on queries over ranges or intervals of available data. This can be a tedious and slow process, especially when there are numerous and repeated queries. Segment trees allow us to handle such queries efficiently in logarithmic time complexity.

Segment trees can be used in computational geometry and the field of [geographic information systems](https://en.wikipedia.org/wiki/Geographic_information_systems). For example, there may be a large number of points in space at certain distances from a central reference point/origin. Suppose we want to find points within a certain distance range from the origin. A normal lookup table would require a linear scan over all possible points or all possible distances (assuming it is a hash map). A segment tree enables us to achieve this requirement in logarithmic time while requiring much less space. Such problems are called [plane range searching](https://en.wikipedia.org/wiki/Range_searching). Efficiently solving such problems is crucial, especially when dealing with dynamic data that changes rapidly (for example, radar systems used for air traffic). Below, we will use a segment tree to solve the Range Sum Query problem as an example.

![](https://img.halfrost.com/Blog/ArticleImage/153_3.png)


The figure above is a segment tree used for range queries.

## III. Constructing a Segment Tree


Assume the data exists in arr[] with size n.

1. The root of the segment tree usually represents the entire data interval. Here it is arr[0:n-1].
2. Each leaf of the tree represents a range containing only one element. Therefore, the leaves represent arr[0], arr[1], and so on, up to arr[n-1].
3. The internal nodes of the tree represent the merged result or union result of their child nodes.
4. Each child node may represent roughly half of the range represented by its parent node. (The idea of binary splitting)

A segment tree for a range of n elements can be easily represented using an array of size {{< katex >}}\approx 4 \ast n {{< /katex >}}. ([Stack Overflow](http://stackoverflow.com/q/28470692/2844164) has a good discussion of why. If you are still not sure, do not worry. This article will discuss it later.)

A node with index i has two nodes with indices {{< katex >}}(2 \ast i + 1) {{< /katex >}} and {{< katex >}}(2 \ast i + 2){{< /katex >}}.

![](https://img.halfrost.com/Blog/ArticleImage/153_4.png)

A segment tree looks intuitive and is very suitable for recursive construction.

We will use the array tree[] to store the nodes of the segment tree (initialized to all zeros). Indices start from 0.

- The node of the tree is at index 0. Therefore tree[0] is the root of the tree.
- The children of tree[i] are in tree[2 * i + 1] and tree[2 * i + 2].
- Fill arr[] with extra 0 or null values so that {{< katex >}}n = 2^{k} {{< /katex >}} (where n is the total length of arr[], and k is a non-negative integer.)
- The index range of leaf nodes is {{< katex >}} \in [2^{k}-1, 2^{k+1}-2]{{< /katex >}}

![](https://img.halfrost.com/Blog/ArticleImage/153_5.png)

The code for constructing a segment tree is as follows:


```go
// SegmentTree define
type SegmentTree struct {
	data, tree, lazy []int
	left, right      int
	merge            func(i, j int) int
}

// Init define
func (st *SegmentTree) Init(nums []int, oper func(i, j int) int) {
	st.merge = oper
	data, tree, lazy := make([]int, len(nums)), make([]int, 4*len(nums)), make([]int, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	st.data, st.tree, st.lazy = data, tree, lazy
	if len(nums) > 0 {
		st.buildSegmentTree(0, 0, len(nums)-1)
	}
}

// Create a segment tree for the [left....right] interval at the position treeIndex
func (st *SegmentTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	st.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree) leftChild(index int) int {
	return 2*index + 1
}

func (st *SegmentTree) rightChild(index int) int {
	return 2*index + 2
}
```

The author turns the segment tree merge operation into a function. The merge operation changes according to the problem statement; common ones include addition, taking max, min, and so on.

We use arr[] = [18, 17, 13, 19, 15, 11, 20, 12, 33, 25 ] as an example to construct a segment tree:

![](https://img.halfrost.com/Blog/ArticleImage/153_6.png)

After the segment tree is constructed, the data in the array is:

```c
tree[] = [ 183, 82, 101, 48, 34, 43, 58, 35, 13, 19, 15, 31, 12, 33, 25, 18, 17, 0, 0, 0, 0, 0, 0, 11, 20, 0, 0, 0, 0, 0, 0 ]
```

The segment tree is padded with 0 to 4*n elements.


> The corresponding LeetCode problems are [218. The Skyline Problem](https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0218.The-Skyline-Problem/), [303. Range Sum Query - Immutable](https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0303.Range-Sum-Query-Immutable/), [307. Range Sum Query - Mutable](https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0307.Range-Sum-Query-Mutable/), and [699. Falling Squares](https://books.halfrost.com/leetcode/ChapterFour/0600~0699/0699.Falling-Squares/)

## IV. Querying a Segment Tree

There are two methods for querying a segment tree: one is direct query, and the other is lazy query.

### 1. Direct Query

When the query range completely matches the range represented by the current node, this method returns the result. Otherwise, it traverses deeper into the segment tree to find nodes that completely match part of the node.


```go
// Query the value in the [left....right] interval

// Query define
func (st *SegmentTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.queryInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// In the segment tree rooted at treeIndex, within the [left...right] range, search for the value of interval [queryLeft...queryRight]
func (st *SegmentTree) queryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if left == queryLeft && right == queryRight {
		return st.tree[treeIndex]
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if queryLeft > midTreeIndex {
		return st.queryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	return st.merge(st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.queryInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}
```


![](https://img.halfrost.com/Blog/ArticleImage/153_7.png)


In the example above, the query interval range is the sum of elements in [2, 8]. No single segment can fully represent the [2, 8] range. However, it can be observed that the four intervals [2, 2], [3, 4], [5, 7], and [8, 8] can be used to form [8, 8]. Quickly verifying, the sum of input elements at [2,8] is 13 + 19 + 15 + 11 + 20 + 12 + 33 = 123. The total sum of the nodes [2, 2], [3, 4], [5, 7], and [8, 8] is 13 + 34 + 43 + 33 = 123. The answer is correct.



### 2. Lazy Query

Lazy query corresponds to lazy update; the two are paired operations. During interval updates, not all nodes in the interval are updated directly. Instead, the value by which nodes in the interval increase or decrease is stored in the lazy array. The increase or decrease is then applied to specific nodes during the next query. This is also done to amortize time complexity and ensure that the time complexity of queries and updates remains at the O(log n) level, without degrading to the O(n) level.

Steps for querying a lazy node:

1. First determine whether the current node is a lazy node. This is determined by checking whether lazy[i] is 0. If it is a lazy node, apply its increase or decrease to this node and update its child nodes. This step is exactly the same as the first step of the update operation.
2. Recursively query child nodes to find suitable query nodes.

The specific code is as follows:

```go
// Query the value in the [left....right] interval

// QueryLazy define
func (st *SegmentTree) QueryLazy(left, right int) int {
	if len(st.data) > 0 {
		return st.queryLazyInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

func (st *SegmentTree) queryLazyInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if left > queryRight || right < queryLeft { // segment completely outside range
		return 0 // represents a null node
	}
	if st.lazy[treeIndex] != 0 { // this node is lazy
		// 幂等 merge（如 max/min）整段套用一次即可，O(1) 下推；
		// 按区间长度循环会退化成 O(区间)。求和语义则改成 (right-left+1)*st.lazy[treeIndex]。
		st.tree[treeIndex] = st.merge(st.tree[treeIndex], st.lazy[treeIndex])
		if left != right { // update lazy[] for children nodes
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], st.lazy[treeIndex])
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], st.lazy[treeIndex])
			// st.lazy[leftTreeIndex] += st.lazy[treeIndex]
			// st.lazy[rightTreeIndex] += st.lazy[treeIndex]
		}
		st.lazy[treeIndex] = 0 // current node processed. No longer lazy
	}
	if queryLeft <= left && queryRight >= right { // segment completely inside range
		return st.tree[treeIndex]
	}
	if queryLeft > midTreeIndex {
		return st.queryLazyInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
	} else if queryRight <= midTreeIndex {
		return st.queryLazyInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
	}
	// merge query results
	return st.merge(st.queryLazyInTree(leftTreeIndex, left, midTreeIndex, queryLeft, midTreeIndex),
		st.queryLazyInTree(rightTreeIndex, midTreeIndex+1, right, midTreeIndex+1, queryRight))
}
```


## V. Updating a Segment Tree

### 1. Single-Point Update

A single-point update is similar to `buildSegTree`. It updates the value of a leaf node of the tree, corresponding to the updated element. These updated values propagate their influence to the root through the upper nodes of the tree.


```go
// Update the value at position index

// Update define
func (st *SegmentTree) Update(index, val int) {
	if len(st.data) > 0 {
		st.updateInTree(0, 0, len(st.data)-1, index, val)
	}
}

// With treeIndex as the root, update the value at position index to val
func (st *SegmentTree) updateInTree(treeIndex, left, right, index, val int) {
	if left == right {
		st.tree[treeIndex] = val
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if index > midTreeIndex {
		st.updateInTree(rightTreeIndex, midTreeIndex+1, right, index, val)
	} else {
		st.updateInTree(leftTreeIndex, left, midTreeIndex, index, val)
	}
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}
```

![](https://img.halfrost.com/Blog/ArticleImage/153_8.png)

In this example, the elements at indices (in the original input data) 1, 3, and 6 are increased by +3, -1, and +2 respectively. You can see how the changes propagate along the tree all the way to the root.


### 2. Interval Update


A segment tree is very efficient for updating only a single element, with time complexity O(log n). But what if we want to update a series of elements? According to the current method, each element must be updated independently, and each element will take some time. Updating each leaf node separately means processing their common ancestors multiple times. Ancestor nodes may be updated multiple times. What should we do if we want to reduce this repeated computation?


![](https://img.halfrost.com/Blog/ArticleImage/153_11.png)

In the example above, the root node is updated three times, and the node numbered 82 is updated twice. This is because updating leaf nodes affects upper parent nodes. In the worst case, the queried interval does not contain frequently updated elements, so a lot of time is spent updating nodes that are rarely accessed. Adding an extra lazy array can reduce unnecessary computation and process nodes on demand.

Use another array lazy[], whose size is exactly the same as our segment tree array tree[], representing a lazy node. When this node is accessed or queried, lazy[i] retains the amount that needs to be added to or subtracted from this node tree[i]. When lazy[i] is 0, it means that the node tree[i] is not lazy and has no cached update.

Steps for updating nodes within an interval:

1. First determine whether the current node is a lazy node. This is determined by checking whether lazy[i] is 0. If it is a lazy node, apply its increase or decrease to this node and update its child nodes.
2. If the interval represented by the current node lies within the update range, apply the current update operation to the current node.
3. Recursively update child nodes.

The specific code is as follows:

```go

// Update the value at positions [updateLeft....updateRight]
// Note that the update value here increases or decreases based on the original value, rather than assigning all values in this interval to x. Interval updates are different from single-point updates
// The interval update here focuses on changes, while single-point update focuses on fixed values
// Of course, interval updates can also update everything to fixed values. If only interval updates to fixed values are used, then the lazy update strategy needs to change, and the merge strategy also needs to change. This will not be discussed in detail here

// UpdateLazy define
func (st *SegmentTree) UpdateLazy(updateLeft, updateRight, val int) {
	if len(st.data) > 0 {
		st.updateLazyInTree(0, 0, len(st.data)-1, updateLeft, updateRight, val)
	}
}

func (st *SegmentTree) updateLazyInTree(treeIndex, left, right, updateLeft, updateRight, val int) {
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	if st.lazy[treeIndex] != 0 { // this node is lazy
		// 幂等 merge（如 max/min）整段套用一次即可，O(1) 下推；
		// 按区间长度循环会退化成 O(区间)。求和语义则改成 (right-left+1)*st.lazy[treeIndex]。
		st.tree[treeIndex] = st.merge(st.tree[treeIndex], st.lazy[treeIndex])
		if left != right { // update lazy[] for children nodes
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], st.lazy[treeIndex])
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], st.lazy[treeIndex])
			// st.lazy[leftTreeIndex] += st.lazy[treeIndex]
			// st.lazy[rightTreeIndex] += st.lazy[treeIndex]
		}
		st.lazy[treeIndex] = 0 // current node processed. No longer lazy
	}

	if left > right || left > updateRight || right < updateLeft {
		return // out of range. escape.
	}

	if updateLeft <= left && right <= updateRight { // segment is fully within update range
		// 幂等 merge（如 max/min）整段套用一次即可，O(1) 下推；
		// 按区间长度循环会退化成 O(区间)。求和语义则改成 (right-left+1)*val。
		st.tree[treeIndex] = st.merge(st.tree[treeIndex], val)
		if left != right { // update lazy[] for children
			st.lazy[leftTreeIndex] = st.merge(st.lazy[leftTreeIndex], val)
			st.lazy[rightTreeIndex] = st.merge(st.lazy[rightTreeIndex], val)
			// st.lazy[leftTreeIndex] += val
			// st.lazy[rightTreeIndex] += val
		}
		return
	}
	st.updateLazyInTree(leftTreeIndex, left, midTreeIndex, updateLeft, updateRight, val)
	st.updateLazyInTree(rightTreeIndex, midTreeIndex+1, right, updateLeft, updateRight, val)
	// merge updates
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

```

> The corresponding LeetCode problems are [218. The Skyline Problem](https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0218.The-Skyline-Problem/) and [699. Falling Squares](https://books.halfrost.com/leetcode/ChapterFour/0600~0699/0699.Falling-Squares/)

## VI. Time Complexity Analysis

Let us look at the build process. We visit every leaf of the segment tree (corresponding to each element in the array arr[]). Therefore, we process about 2 * n nodes. This makes the time complexity of the build process O(n). For each recursive update process, half of the interval range is discarded to reach a leaf node in the tree. This is similar to binary search and only requires logarithmic time. After updating the leaf, the direct ancestors at each level of the tree are updated. This takes time linear in the height of the tree.

![](https://img.halfrost.com/Blog/ArticleImage/153_9.png)


4\*n nodes can ensure that the segment tree is built as a complete binary tree, so the height of the tree is the ceiling of log(4\*n + 1). The time complexity of reading and updating a segment tree is both O(log n).

## VII. Common Problem Types


### 1. Range Sum Queries

![](https://img.halfrost.com/Blog/ArticleImage/153_10.png)


Range Sum Queries are a subset of the [Range Queries](https://en.wikipedia.org/wiki/Range_query_(data_structures)) problem. Given an array or sequence of data elements, it is necessary to handle read and update queries consisting of ranges of elements. Both the segment tree Segment Tree and the binary indexed tree Binary Indexed Tree (a.k.a. Fenwick Tree)) can solve this kind of problem very quickly.

The Range Sum Query problem specifically handles the sum of elements within a query range. There are many variants of this problem, including [immutable data](https://leetcode.com/problems/range-sum-query-immutable/), [mutable data](https://leetcode.com/problems/range-sum-query-mutable/), [multiple updates, single query](https://leetcode.com/problems/range-addition/) and [multiple updates, multiple queries](https://leetcode.com/problems/range-sum-query-2d-mutable/).



### 2. Single-Point Update    
- [HDU 1166 Enemy Troops Formation](http://acm.hdu.edu.cn/showproblem.php?pid=1166) update:single-point increase/decrease query:interval sum  
- [HDU 1754 I Hate It](http://acm.hdu.edu.cn/showproblem.php?pid=1754) update:single-point replacement query:interval extremum  
- [HDU 1394 Minimum Inversion Number](http://acm.hdu.edu.cn/showproblem.php?pid=1394) update:single-point increase/decrease query:interval sum  
- [HDU 2795 Billboard](http://acm.hdu.edu.cn/showproblem.php?pid=2795) query:find the position of the maximum value in the interval (directly perform the update operation inside query)

### 3. Interval Update  

- [HDU 1698 Just a Hook](http://acm.hdu.edu.cn/showproblem.php?pid=1698) update:segment replacement (since the total interval is queried only once, the information of node 1 can be output directly)  
- [POJ 3468 A Simple Problem with Integers](http://poj.org/problem?id=3468) update:segment increase/decrease query:interval sum  
- [POJ 2528 Mayor’s posters](http://poj.org/problem?id=2528) discretization + update:segment replacement query:simple hash  
- [POJ 3225 Help with Intervals](http://poj.org/problem?id=3225) update:segment replacement, interval XOR query:simple hash

### 4. Interval Merging

This type of problem asks for the longest continuous interval in an interval that satisfies certain conditions, so when PushUp is performed, the intervals of the left and right children need to be merged

- [POJ 3667 Hotel](http://poj.org/problem?id=3667) update:interval replacement query:ask for the leftmost endpoint that satisfies the condition

### 5. Scan Line

This type of problem requires sorting some operations, and then using a scan line to sweep from left to right. The most typical examples are union of rectangle areas, union of perimeters, and similar problems

- [HDU 1542 Atlantis](http://acm.hdu.edu.cn/showproblem.php?pid=1542) update:interval increase/decrease query:directly take the value of the root node  
- [HDU 1828 Picture](http://acm.hdu.edu.cn/showproblem.php?pid=1828) update:interval increase/decrease query:directly take the value of the root node


### 6. Counting Problems

There is also a type of problem on LeetCode involving counting. Problems like [315. Count of Smaller Numbers After Self](https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0315.Count-of-Smaller-Numbers-After-Self/), [327. Count of Range Sum](https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0327.Count-of-Range-Sum/), and [493. Reverse Pairs](https://books.halfrost.com/leetcode/ChapterFour/0400~0499/0493.Reverse-Pairs/) can be solved using the following pattern. Each node of the segment tree stores an interval count.



```go
// SegmentCountTree define
type SegmentCountTree struct {
	data, tree  []int
	left, right int
	merge       func(i, j int) int
}

// Init define
func (st *SegmentCountTree) Init(nums []int, oper func(i, j int) int) {
	st.merge = oper

	data, tree := make([]int, len(nums)), make([]int, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	st.data, st.tree = data, tree
}

// Create a segment tree for the [left....right] interval at the position treeIndex
func (st *SegmentCountTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	st.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentCountTree) leftChild(index int) int {
	return 2*index + 1
}

func (st *SegmentCountTree) rightChild(index int) int {
	return 2*index + 2
}

// Query the value in the [left....right] interval

// Query define
func (st *SegmentCountTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.queryInTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

// In the segment tree rooted at treeIndex, within the [left...right] range, search for the value of interval [queryLeft...queryRight]; the value is a count value
func (st *SegmentCountTree) queryInTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if queryRight < st.data[left] || queryLeft > st.data[right] {
		return 0
	}
	if queryLeft <= st.data[left] && queryRight >= st.data[right] || left == right {
		return st.tree[treeIndex]
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	return st.queryInTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight) +
		st.queryInTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight)
}

// Update the count

// UpdateCount define
func (st *SegmentCountTree) UpdateCount(val int) {
	if len(st.data) > 0 {
		st.updateCountInTree(0, 0, len(st.data)-1, val)
	}
}

// With treeIndex as the root, update the count within the [left...right] interval
func (st *SegmentCountTree) updateCountInTree(treeIndex, left, right, val int) {
	if val >= st.data[left] && val <= st.data[right] {
		st.tree[treeIndex]++
		if left == right {
			return
		}
		midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
		st.updateCountInTree(rightTreeIndex, midTreeIndex+1, right, val)
		st.updateCountInTree(leftTreeIndex, left, midTreeIndex, val)
	}
}

```
