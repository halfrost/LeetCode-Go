---
title: 2.11 Binary Search
type: docs
weight: 11
---

# Binary Search

- 二分搜索的经典写法。需要注意的三点：
	1. 循环退出条件，注意是 low <= high，而不是 low < high。
	2. mid 的取值，mid := low + (high-low)>>1
	3. low 和 high 的更新。low = mid + 1，high = mid - 1。

```go
func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
```

- 二分搜索的变种写法。有 4 个基本变种:
	1. 查找第一个与 target 相等的元素，时间复杂度 O(logn)
	2. 查找最后一个与 target 相等的元素，时间复杂度 O(logn)
	3. 查找第一个大于等于 target 的元素，时间复杂度 O(logn)
	4. 查找最后一个小于等于 target 的元素，时间复杂度 O(logn)

```go
// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后一个小于等于 target 的元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

- 在基本有序的数组中用二分搜索。经典解法可以解，变种写法也可以写，常见的题型，在山峰数组中找山峰，在旋转有序数组中找分界点。第 33 题，第 81 题，第 153 题，第 154 题，第 162 题，第 852 题

```go
func peakIndexInMountainArray(A []int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := low + (high-low)>>1
		// 如果 mid 较大，则左侧存在峰值，high = m，如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
		if A[mid] > A[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
```

- max-min 最大值最小化问题。求在最小满足条件的情况下的最大值。第 410 题，第 875 题，第 1011 题，第 1283 题。


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0004|Median of Two Sorted Arrays|[Go]({{< relref "/ChapterFour/0001~0099/0004.Median-of-Two-Sorted-Arrays.md" >}})|Hard||||35.2%|
|0033|Search in Rotated Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0033.Search-in-Rotated-Sorted-Array.md" >}})|Medium||||38.6%|
|0034|Find First and Last Position of Element in Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0034.Find-First-and-Last-Position-of-Element-in-Sorted-Array.md" >}})|Medium||||41.5%|
|0035|Search Insert Position|[Go]({{< relref "/ChapterFour/0001~0099/0035.Search-Insert-Position.md" >}})|Easy||||42.0%|
|0069|Sqrt(x)|[Go]({{< relref "/ChapterFour/0001~0099/0069.Sqrtx.md" >}})|Easy| O(log n)| O(1)||37.0%|
|0074|Search a 2D Matrix|[Go]({{< relref "/ChapterFour/0001~0099/0074.Search-a-2D-Matrix.md" >}})|Medium||||46.8%|
|0081|Search in Rotated Sorted Array II|[Go]({{< relref "/ChapterFour/0001~0099/0081.Search-in-Rotated-Sorted-Array-II.md" >}})|Medium||||35.7%|
|0153|Find Minimum in Rotated Sorted Array|[Go]({{< relref "/ChapterFour/0100~0199/0153.Find-Minimum-in-Rotated-Sorted-Array.md" >}})|Medium||||48.5%|
|0154|Find Minimum in Rotated Sorted Array II|[Go]({{< relref "/ChapterFour/0100~0199/0154.Find-Minimum-in-Rotated-Sorted-Array-II.md" >}})|Hard||||43.4%|
|0162|Find Peak Element|[Go]({{< relref "/ChapterFour/0100~0199/0162.Find-Peak-Element.md" >}})|Medium||||46.2%|
|0167|Two Sum II - Input Array Is Sorted|[Go]({{< relref "/ChapterFour/0100~0199/0167.Two-Sum-II-Input-Array-Is-Sorted.md" >}})|Medium| O(n)| O(1)||60.0%|
|0209|Minimum Size Subarray Sum|[Go]({{< relref "/ChapterFour/0200~0299/0209.Minimum-Size-Subarray-Sum.md" >}})|Medium| O(n)| O(1)||44.4%|
|0222|Count Complete Tree Nodes|[Go]({{< relref "/ChapterFour/0200~0299/0222.Count-Complete-Tree-Nodes.md" >}})|Medium| O(n)| O(1)||57.4%|
|0240|Search a 2D Matrix II|[Go]({{< relref "/ChapterFour/0200~0299/0240.Search-a-2D-Matrix-II.md" >}})|Medium||||50.5%|
|0268|Missing Number|[Go]({{< relref "/ChapterFour/0200~0299/0268.Missing-Number.md" >}})|Easy||||61.5%|
|0275|H-Index II|[Go]({{< relref "/ChapterFour/0200~0299/0275.H-Index-II.md" >}})|Medium||||37.4%|
|0278|First Bad Version|[Go]({{< relref "/ChapterFour/0200~0299/0278.First-Bad-Version.md" >}})|Easy||||42.9%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium| O(n)| O(1)|❤️|59.1%|
|0300|Longest Increasing Subsequence|[Go]({{< relref "/ChapterFour/0300~0399/0300.Longest-Increasing-Subsequence.md" >}})|Medium| O(n log n)| O(n)||51.5%|
|0315|Count of Smaller Numbers After Self|[Go]({{< relref "/ChapterFour/0300~0399/0315.Count-of-Smaller-Numbers-After-Self.md" >}})|Hard||||42.8%|
|0327|Count of Range Sum|[Go]({{< relref "/ChapterFour/0300~0399/0327.Count-of-Range-Sum.md" >}})|Hard||||36.0%|
|0349|Intersection of Two Arrays|[Go]({{< relref "/ChapterFour/0300~0399/0349.Intersection-of-Two-Arrays.md" >}})|Easy| O(n)| O(n) ||70.2%|
|0350|Intersection of Two Arrays II|[Go]({{< relref "/ChapterFour/0300~0399/0350.Intersection-of-Two-Arrays-II.md" >}})|Easy| O(n)| O(n) ||55.5%|
|0352|Data Stream as Disjoint Intervals|[Go]({{< relref "/ChapterFour/0300~0399/0352.Data-Stream-as-Disjoint-Intervals.md" >}})|Hard||||51.5%|
|0354|Russian Doll Envelopes|[Go]({{< relref "/ChapterFour/0300~0399/0354.Russian-Doll-Envelopes.md" >}})|Hard||||38.3%|
|0367|Valid Perfect Square|[Go]({{< relref "/ChapterFour/0300~0399/0367.Valid-Perfect-Square.md" >}})|Easy||||43.3%|
|0374|Guess Number Higher or Lower|[Go]({{< relref "/ChapterFour/0300~0399/0374.Guess-Number-Higher-or-Lower.md" >}})|Easy||||50.4%|
|0378|Kth Smallest Element in a Sorted Matrix|[Go]({{< relref "/ChapterFour/0300~0399/0378.Kth-Smallest-Element-in-a-Sorted-Matrix.md" >}})|Medium||||61.6%|
|0400|Nth Digit|[Go]({{< relref "/ChapterFour/0400~0499/0400.Nth-Digit.md" >}})|Medium||||34.0%|
|0410|Split Array Largest Sum|[Go]({{< relref "/ChapterFour/0400~0499/0410.Split-Array-Largest-Sum.md" >}})|Hard||||53.2%|
|0436|Find Right Interval|[Go]({{< relref "/ChapterFour/0400~0499/0436.Find-Right-Interval.md" >}})|Medium||||50.3%|
|0441|Arranging Coins|[Go]({{< relref "/ChapterFour/0400~0499/0441.Arranging-Coins.md" >}})|Easy||||46.1%|
|0456|132 Pattern|[Go]({{< relref "/ChapterFour/0400~0499/0456.132-Pattern.md" >}})|Medium||||32.4%|
|0475|Heaters|[Go]({{< relref "/ChapterFour/0400~0499/0475.Heaters.md" >}})|Medium||||36.1%|
|0483|Smallest Good Base|[Go]({{< relref "/ChapterFour/0400~0499/0483.Smallest-Good-Base.md" >}})|Hard||||38.4%|
|0493|Reverse Pairs|[Go]({{< relref "/ChapterFour/0400~0499/0493.Reverse-Pairs.md" >}})|Hard||||30.8%|
|0497|Random Point in Non-overlapping Rectangles|[Go]({{< relref "/ChapterFour/0400~0499/0497.Random-Point-in-Non-overlapping-Rectangles.md" >}})|Medium||||39.3%|
|0528|Random Pick with Weight|[Go]({{< relref "/ChapterFour/0500~0599/0528.Random-Pick-with-Weight.md" >}})|Medium||||46.1%|
|0532|K-diff Pairs in an Array|[Go]({{< relref "/ChapterFour/0500~0599/0532.K-diff-Pairs-in-an-Array.md" >}})|Medium||||40.7%|
|0540|Single Element in a Sorted Array|[Go]({{< relref "/ChapterFour/0500~0599/0540.Single-Element-in-a-Sorted-Array.md" >}})|Medium||||58.5%|
|0611|Valid Triangle Number|[Go]({{< relref "/ChapterFour/0600~0699/0611.Valid-Triangle-Number.md" >}})|Medium||||50.4%|
|0633|Sum of Square Numbers|[Go]({{< relref "/ChapterFour/0600~0699/0633.Sum-of-Square-Numbers.md" >}})|Medium||||34.6%|
|0658|Find K Closest Elements|[Go]({{< relref "/ChapterFour/0600~0699/0658.Find-K-Closest-Elements.md" >}})|Medium||||46.7%|
|0668|Kth Smallest Number in Multiplication Table|[Go]({{< relref "/ChapterFour/0600~0699/0668.Kth-Smallest-Number-in-Multiplication-Table.md" >}})|Hard||||51.5%|
|0704|Binary Search|[Go]({{< relref "/ChapterFour/0700~0799/0704.Binary-Search.md" >}})|Easy||||55.1%|
|0710|Random Pick with Blacklist|[Go]({{< relref "/ChapterFour/0700~0799/0710.Random-Pick-with-Blacklist.md" >}})|Hard| O(n)| O(n)  ||33.6%|
|0718|Maximum Length of Repeated Subarray|[Go]({{< relref "/ChapterFour/0700~0799/0718.Maximum-Length-of-Repeated-Subarray.md" >}})|Medium||||51.6%|
|0719|Find K-th Smallest Pair Distance|[Go]({{< relref "/ChapterFour/0700~0799/0719.Find-K-th-Smallest-Pair-Distance.md" >}})|Hard||||36.3%|
|0729|My Calendar I|[Go]({{< relref "/ChapterFour/0700~0799/0729.My-Calendar-I.md" >}})|Medium||||57.2%|
|0732|My Calendar III|[Go]({{< relref "/ChapterFour/0700~0799/0732.My-Calendar-III.md" >}})|Hard||||71.6%|
|0744|Find Smallest Letter Greater Than Target|[Go]({{< relref "/ChapterFour/0700~0799/0744.Find-Smallest-Letter-Greater-Than-Target.md" >}})|Easy||||44.6%|
|0778|Swim in Rising Water|[Go]({{< relref "/ChapterFour/0700~0799/0778.Swim-in-Rising-Water.md" >}})|Hard||||59.6%|
|0786|K-th Smallest Prime Fraction|[Go]({{< relref "/ChapterFour/0700~0799/0786.K-th-Smallest-Prime-Fraction.md" >}})|Medium||||50.7%|
|0793|Preimage Size of Factorial Zeroes Function|[Go]({{< relref "/ChapterFour/0700~0799/0793.Preimage-Size-of-Factorial-Zeroes-Function.md" >}})|Hard||||42.7%|
|0825|Friends Of Appropriate Ages|[Go]({{< relref "/ChapterFour/0800~0899/0825.Friends-Of-Appropriate-Ages.md" >}})|Medium||||46.4%|
|0826|Most Profit Assigning Work|[Go]({{< relref "/ChapterFour/0800~0899/0826.Most-Profit-Assigning-Work.md" >}})|Medium||||44.4%|
|0852|Peak Index in a Mountain Array|[Go]({{< relref "/ChapterFour/0800~0899/0852.Peak-Index-in-a-Mountain-Array.md" >}})|Medium||||69.5%|
|0862|Shortest Subarray with Sum at Least K|[Go]({{< relref "/ChapterFour/0800~0899/0862.Shortest-Subarray-with-Sum-at-Least-K.md" >}})|Hard||||26.1%|
|0875|Koko Eating Bananas|[Go]({{< relref "/ChapterFour/0800~0899/0875.Koko-Eating-Bananas.md" >}})|Medium||||52.4%|
|0878|Nth Magical Number|[Go]({{< relref "/ChapterFour/0800~0899/0878.Nth-Magical-Number.md" >}})|Hard||||35.7%|
|0887|Super Egg Drop|[Go]({{< relref "/ChapterFour/0800~0899/0887.Super-Egg-Drop.md" >}})|Hard||||27.2%|
|0888|Fair Candy Swap|[Go]({{< relref "/ChapterFour/0800~0899/0888.Fair-Candy-Swap.md" >}})|Easy||||60.5%|
|0911|Online Election|[Go]({{< relref "/ChapterFour/0900~0999/0911.Online-Election.md" >}})|Medium||||52.0%|
|0981|Time Based Key-Value Store|[Go]({{< relref "/ChapterFour/0900~0999/0981.Time-Based-Key-Value-Store.md" >}})|Medium||||53.6%|
|1004|Max Consecutive Ones III|[Go]({{< relref "/ChapterFour/1000~1099/1004.Max-Consecutive-Ones-III.md" >}})|Medium||||63.5%|
|1011|Capacity To Ship Packages Within D Days|[Go]({{< relref "/ChapterFour/1000~1099/1011.Capacity-To-Ship-Packages-Within-D-Days.md" >}})|Medium||||64.5%|
|1157|Online Majority Element In Subarray|[Go]({{< relref "/ChapterFour/1100~1199/1157.Online-Majority-Element-In-Subarray.md" >}})|Hard||||42.0%|
|1170|Compare Strings by Frequency of the Smallest Character|[Go]({{< relref "/ChapterFour/1100~1199/1170.Compare-Strings-by-Frequency-of-the-Smallest-Character.md" >}})|Medium||||61.3%|
|1201|Ugly Number III|[Go]({{< relref "/ChapterFour/1200~1299/1201.Ugly-Number-III.md" >}})|Medium||||28.5%|
|1208|Get Equal Substrings Within Budget|[Go]({{< relref "/ChapterFour/1200~1299/1208.Get-Equal-Substrings-Within-Budget.md" >}})|Medium||||47.6%|
|1235|Maximum Profit in Job Scheduling|[Go]({{< relref "/ChapterFour/1200~1299/1235.Maximum-Profit-in-Job-Scheduling.md" >}})|Hard||||51.1%|
|1283|Find the Smallest Divisor Given a Threshold|[Go]({{< relref "/ChapterFour/1200~1299/1283.Find-the-Smallest-Divisor-Given-a-Threshold.md" >}})|Medium||||55.3%|
|1300|Sum of Mutated Array Closest to Target|[Go]({{< relref "/ChapterFour/1300~1399/1300.Sum-of-Mutated-Array-Closest-to-Target.md" >}})|Medium||||43.1%|
|1337|The K Weakest Rows in a Matrix|[Go]({{< relref "/ChapterFour/1300~1399/1337.The-K-Weakest-Rows-in-a-Matrix.md" >}})|Easy||||73.0%|
|1385|Find the Distance Value Between Two Arrays|[Go]({{< relref "/ChapterFour/1300~1399/1385.Find-the-Distance-Value-Between-Two-Arrays.md" >}})|Easy||||65.4%|
|1439|Find the Kth Smallest Sum of a Matrix With Sorted Rows|[Go]({{< relref "/ChapterFour/1400~1499/1439.Find-the-Kth-Smallest-Sum-of-a-Matrix-With-Sorted-Rows.md" >}})|Hard||||61.4%|
|1482|Minimum Number of Days to Make m Bouquets|[Go]({{< relref "/ChapterFour/1400~1499/1482.Minimum-Number-of-Days-to-Make-m-Bouquets.md" >}})|Medium||||56.5%|
|1539|Kth Missing Positive Number|[Go]({{< relref "/ChapterFour/1500~1599/1539.Kth-Missing-Positive-Number.md" >}})|Easy||||55.9%|
|1608|Special Array With X Elements Greater Than or Equal X|[Go]({{< relref "/ChapterFour/1600~1699/1608.Special-Array-With-X-Elements-Greater-Than-or-Equal-X.md" >}})|Easy||||60.0%|
|1631|Path With Minimum Effort|[Go]({{< relref "/ChapterFour/1600~1699/1631.Path-With-Minimum-Effort.md" >}})|Medium||||55.3%|
|1648|Sell Diminishing-Valued Colored Balls|[Go]({{< relref "/ChapterFour/1600~1699/1648.Sell-Diminishing-Valued-Colored-Balls.md" >}})|Medium||||30.5%|
|1649|Create Sorted Array through Instructions|[Go]({{< relref "/ChapterFour/1600~1699/1649.Create-Sorted-Array-through-Instructions.md" >}})|Hard||||37.2%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1600~1699/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||37.6%|
|1818|Minimum Absolute Sum Difference|[Go]({{< relref "/ChapterFour/1800~1899/1818.Minimum-Absolute-Sum-Difference.md" >}})|Medium||||30.1%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Breadth_First_Search/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Math/">下一页➡️</a></p>
</div>
