---
title: 2.18 ✅ Segment Tree
type: docs
weight: 18
---

# Segment Tree

![](https://img.halfrost.com/Leetcode/Segment_Tree.png)

- Classic array-based implementation of a segment tree. The pushUp logic for merging two nodes is abstracted out, so any operation can be implemented (common operations include addition, taking max, min, and so on). Problem 218, Problem 303, Problem 307, Problem 699.
- Classic implementation of a counting segment tree. Problem 315, Problem 327, Problem 493.
- Tree-based implementation of a segment tree. Problem 715, Problem 732.
- Lazy range update. Problem 218, Problem 699.
- Discretization. For discretization, note a special case: if the three intervals are [1,10] [1,4] [6,10], after discretization x[1]=1,x[2]=4,x[3]=6,x[4]=10. The first interval is [1,4], the second interval is [1,2], and the third interval is [3,4]. In this way, interval one = interval two + interval three, which does not match the model before discretization. Before discretization, it is obvious that interval one > interval two + interval three. The correct approach is: add a number between numbers whose difference is greater than 1. For example, add 5 between 1 4 6 10 above, so x[1]=1,x[2]=4,x[3]=5,x[4]=6,x[5]=10. After handling it this way, interval one is 1-5, interval two is 1-2, and interval three is 4-5.
- Build segment trees flexibly. Segment tree nodes can store multiple pieces of information, and the pushUp operation for merging two nodes can also be diverse. Problem 850, Problem 1157.


Segment tree [problem types](https://blog.csdn.net/xuechelingxiao/article/details/38313105) from simple to difficult:

1. Point update:  
	[HDU 1166 Enemy Troop Deployment](http://acm.hdu.edu.cn/showproblem.php?pid=1166) update: point increment/decrement query: range sum  
	[HDU 1754 I Hate It](http://acm.hdu.edu.cn/showproblem.php?pid=1754) update: point replacement query: range extremum  
	[HDU 1394 Minimum Inversion Number](http://acm.hdu.edu.cn/showproblem.php?pid=1394) update: point increment/decrement query: range sum  
	[HDU 2795 Billboard](http://acm.hdu.edu.cn/showproblem.php?pid=2795) query: position of the maximum value in a range (the update operation is done directly inside query)
2. Range update:  
	[HDU 1698 Just a Hook](http://acm.hdu.edu.cn/showproblem.php?pid=1698) update: range replacement (because the total range is queried only once, the information of node 1 can be output directly)  
	[POJ 3468 A Simple Problem with Integers](http://poj.org/problem?id=3468) update: range increment/decrement query: range sum  
	[POJ 2528 Mayor’s posters](http://poj.org/problem?id=2528) discretization + update: range replacement query: simple hash  
	[POJ 3225 Help with Intervals](http://poj.org/problem?id=3225) update: range replacement, interval XOR query: simple hash
3. Range merging (this type of problem asks for the longest contiguous interval within a range that satisfies a condition, so during PushUp, the intervals of the left and right children need to be merged):  
	[POJ 3667 Hotel](http://poj.org/problem?id=3667) update: range replacement query: query the leftmost endpoint that satisfies the condition
4. Sweep line (this type of problem requires sorting some operations, then sweeping from left to right with a scan line; the most typical examples are problems such as union area and union perimeter of rectangles):  
	[HDU 1542 Atlantis](http://acm.hdu.edu.cn/showproblem.php?pid=1542) update: range increment/decrement query: directly take the value of the root node  
	[HDU 1828 Picture](http://acm.hdu.edu.cn/showproblem.php?pid=1828) update: range increment/decrement query: directly take the value of the root node


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0218|The Skyline Problem|[Go]({{< relref "/ChapterFour/0200~0299/0218.The-Skyline-Problem.md" >}})|Hard| O(n log n)| O(n)|❤️|41.9%|
|0307|Range Sum Query - Mutable|[Go]({{< relref "/ChapterFour/0300~0399/0307.Range-Sum-Query-Mutable.md" >}})|Medium| O(1)| O(n)||40.7%|
|0315|Count of Smaller Numbers After Self|[Go]({{< relref "/ChapterFour/0300~0399/0315.Count-of-Smaller-Numbers-After-Self.md" >}})|Hard| O(n log n)| O(n)||42.6%|
|0327|Count of Range Sum|[Go]({{< relref "/ChapterFour/0300~0399/0327.Count-of-Range-Sum.md" >}})|Hard| O(n log n)| O(n)|❤️|35.8%|
|0493|Reverse Pairs|[Go]({{< relref "/ChapterFour/0400~0499/0493.Reverse-Pairs.md" >}})|Hard| O(n log n)| O(n)||30.9%|
|0699|Falling Squares|[Go]({{< relref "/ChapterFour/0600~0699/0699.Falling-Squares.md" >}})|Hard| O(n log n)| O(n)|❤️|44.7%|
|0715|Range Module|[Go]({{< relref "/ChapterFour/0700~0799/0715.Range-Module.md" >}})|Hard| O(log n)| O(n)|❤️|44.6%|
|0729|My Calendar I|[Go]({{< relref "/ChapterFour/0700~0799/0729.My-Calendar-I.md" >}})|Medium||||56.8%|
|0732|My Calendar III|[Go]({{< relref "/ChapterFour/0700~0799/0732.My-Calendar-III.md" >}})|Hard| O(log n)| O(n)|❤️|71.5%|
|0850|Rectangle Area II|[Go]({{< relref "/ChapterFour/0800~0899/0850.Rectangle-Area-II.md" >}})|Hard| O(n log n)| O(n)|❤️|53.9%|
|1157|Online Majority Element In Subarray|[Go]({{< relref "/ChapterFour/1100~1199/1157.Online-Majority-Element-In-Subarray.md" >}})|Hard| O(log n)| O(n)|❤️|41.8%|
|1649|Create Sorted Array through Instructions|[Go]({{< relref "/ChapterFour/1600~1699/1649.Create-Sorted-Array-through-Instructions.md" >}})|Hard||||37.5%|
