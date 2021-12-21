---
title: 2.03 ✅ Two Pointers
type: docs
weight: 3
---

# Two Pointers

![](https://img.halfrost.com/Leetcode/Two_pointers.png)

- 双指针滑动窗口的经典写法。右指针不断往右移，移动到不能往右移动为止(具体条件根据题目而定)。当右指针到最右边以后，开始挪动左指针，释放窗口左边界。第 3 题，第 76 题，第 209 题，第 424 题，第 438 题，第 567 题，第 713 题，第 763 题，第 845 题，第 881 题，第 904 题，第 978 题，第 992 题，第 1004 题，第 1040 题，第 1052 题。

```c
	left, right := 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
```

- 快慢指针可以查找重复数字，时间复杂度 O(n)，第 287 题。
- 替换字母以后，相同字母能出现连续最长的长度。第 424 题。
- SUM 问题集。第 1 题，第 15 题，第 16 题，第 18 题，第 167 题，第 923 题，第 1074 题。


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0011|Container With Most Water|[Go]({{< relref "/ChapterFour/0001~0099/0011.Container-With-Most-Water.md" >}})|Medium| O(n)| O(1)||53.2%|
|0015|3Sum|[Go]({{< relref "/ChapterFour/0001~0099/0015.3Sum.md" >}})|Medium| O(n^2)| O(n)|❤️|30.2%|
|0016|3Sum Closest|[Go]({{< relref "/ChapterFour/0001~0099/0016.3Sum-Closest.md" >}})|Medium| O(n^2)| O(1)|❤️|46.9%|
|0018|4Sum|[Go]({{< relref "/ChapterFour/0001~0099/0018.4Sum.md" >}})|Medium| O(n^3)| O(n^2)|❤️|36.8%|
|0019|Remove Nth Node From End of List|[Go]({{< relref "/ChapterFour/0001~0099/0019.Remove-Nth-Node-From-End-of-List.md" >}})|Medium| O(n)| O(1)||37.4%|
|0026|Remove Duplicates from Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array.md" >}})|Easy| O(n)| O(1)||47.9%|
|0027|Remove Element|[Go]({{< relref "/ChapterFour/0001~0099/0027.Remove-Element.md" >}})|Easy| O(n)| O(1)||50.6%|
|0028|Implement strStr()|[Go]({{< relref "/ChapterFour/0001~0099/0028.Implement-strStr.md" >}})|Easy| O(n)| O(1)||35.7%|
|0031|Next Permutation|[Go]({{< relref "/ChapterFour/0001~0099/0031.Next-Permutation.md" >}})|Medium||||34.9%|
|0042|Trapping Rain Water|[Go]({{< relref "/ChapterFour/0001~0099/0042.Trapping-Rain-Water.md" >}})|Hard| O(n)| O(1)|❤️|54.9%|
|0061|Rotate List|[Go]({{< relref "/ChapterFour/0001~0099/0061.Rotate-List.md" >}})|Medium| O(n)| O(1)||33.3%|
|0075|Sort Colors|[Go]({{< relref "/ChapterFour/0001~0099/0075.Sort-Colors.md" >}})|Medium| O(n)| O(1)|❤️|53.3%|
|0080|Remove Duplicates from Sorted Array II|[Go]({{< relref "/ChapterFour/0001~0099/0080.Remove-Duplicates-from-Sorted-Array-II.md" >}})|Medium| O(n)| O(1||48.1%|
|0082|Remove Duplicates from Sorted List II|[Go]({{< relref "/ChapterFour/0001~0099/0082.Remove-Duplicates-from-Sorted-List-II.md" >}})|Medium||||41.7%|
|0086|Partition List|[Go]({{< relref "/ChapterFour/0001~0099/0086.Partition-List.md" >}})|Medium| O(n)| O(1)|❤️|46.9%|
|0088|Merge Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0088.Merge-Sorted-Array.md" >}})|Easy| O(n)| O(1)|❤️|42.6%|
|0125|Valid Palindrome|[Go]({{< relref "/ChapterFour/0100~0199/0125.Valid-Palindrome.md" >}})|Easy| O(n)| O(1)||40.3%|
|0141|Linked List Cycle|[Go]({{< relref "/ChapterFour/0100~0199/0141.Linked-List-Cycle.md" >}})|Easy| O(n)| O(1)|❤️|44.6%|
|0142|Linked List Cycle II|[Go]({{< relref "/ChapterFour/0100~0199/0142.Linked-List-Cycle-II.md" >}})|Medium| O(n)| O(1)|❤️|42.2%|
|0143|Reorder List|[Go]({{< relref "/ChapterFour/0100~0199/0143.Reorder-List.md" >}})|Medium||||45.7%|
|0148|Sort List|[Go]({{< relref "/ChapterFour/0100~0199/0148.Sort-List.md" >}})|Medium||||49.4%|
|0151|Reverse Words in a String|[Go]({{< relref "/ChapterFour/0100~0199/0151.Reverse-Words-in-a-String.md" >}})|Medium||||27.2%|
|0160|Intersection of Two Linked Lists|[Go]({{< relref "/ChapterFour/0100~0199/0160.Intersection-of-Two-Linked-Lists.md" >}})|Easy||||48.1%|
|0167|Two Sum II - Input Array Is Sorted|[Go]({{< relref "/ChapterFour/0100~0199/0167.Two-Sum-II-Input-Array-Is-Sorted.md" >}})|Easy| O(n)| O(1)||57.7%|
|0189|Rotate Array|[Go]({{< relref "/ChapterFour/0100~0199/0189.Rotate-Array.md" >}})|Medium||||37.4%|
|0202|Happy Number|[Go]({{< relref "/ChapterFour/0200~0299/0202.Happy-Number.md" >}})|Easy||||52.4%|
|0234|Palindrome Linked List|[Go]({{< relref "/ChapterFour/0200~0299/0234.Palindrome-Linked-List.md" >}})|Easy| O(n)| O(1)||45.1%|
|0283|Move Zeroes|[Go]({{< relref "/ChapterFour/0200~0299/0283.Move-Zeroes.md" >}})|Easy| O(n)| O(1)||59.9%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium| O(n)| O(1)|❤️|58.3%|
|0344|Reverse String|[Go]({{< relref "/ChapterFour/0300~0399/0344.Reverse-String.md" >}})|Easy| O(n)| O(1)||73.0%|
|0345|Reverse Vowels of a String|[Go]({{< relref "/ChapterFour/0300~0399/0345.Reverse-Vowels-of-a-String.md" >}})|Easy| O(n)| O(1)||46.4%|
|0349|Intersection of Two Arrays|[Go]({{< relref "/ChapterFour/0300~0399/0349.Intersection-of-Two-Arrays.md" >}})|Easy| O(n)| O(n) ||67.8%|
|0350|Intersection of Two Arrays II|[Go]({{< relref "/ChapterFour/0300~0399/0350.Intersection-of-Two-Arrays-II.md" >}})|Easy| O(n)| O(n) ||54.1%|
|0392|Is Subsequence|[Go]({{< relref "/ChapterFour/0300~0399/0392.Is-Subsequence.md" >}})|Easy||||50.0%|
|0457|Circular Array Loop|[Go]({{< relref "/ChapterFour/0400~0499/0457.Circular-Array-Loop.md" >}})|Medium||||31.2%|
|0475|Heaters|[Go]({{< relref "/ChapterFour/0400~0499/0475.Heaters.md" >}})|Medium||||34.7%|
|0524|Longest Word in Dictionary through Deleting|[Go]({{< relref "/ChapterFour/0500~0599/0524.Longest-Word-in-Dictionary-through-Deleting.md" >}})|Medium| O(n)| O(1) ||50.8%|
|0532|K-diff Pairs in an Array|[Go]({{< relref "/ChapterFour/0500~0599/0532.K-diff-Pairs-in-an-Array.md" >}})|Medium| O(n)| O(n)||37.2%|
|0541|Reverse String II|[Go]({{< relref "/ChapterFour/0500~0599/0541.Reverse-String-II.md" >}})|Easy||||49.9%|
|0557|Reverse Words in a String III|[Go]({{< relref "/ChapterFour/0500~0599/0557.Reverse-Words-in-a-String-III.md" >}})|Easy||||76.4%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.2%|
|0581|Shortest Unsorted Continuous Subarray|[Go]({{< relref "/ChapterFour/0500~0599/0581.Shortest-Unsorted-Continuous-Subarray.md" >}})|Medium||||33.7%|
|0611|Valid Triangle Number|[Go]({{< relref "/ChapterFour/0600~0699/0611.Valid-Triangle-Number.md" >}})|Medium||||49.4%|
|0633|Sum of Square Numbers|[Go]({{< relref "/ChapterFour/0600~0699/0633.Sum-of-Square-Numbers.md" >}})|Medium||||34.6%|
|0653|Two Sum IV - Input is a BST|[Go]({{< relref "/ChapterFour/0600~0699/0653.Two-Sum-IV-Input-is-a-BST.md" >}})|Easy||||58.2%|
|0658|Find K Closest Elements|[Go]({{< relref "/ChapterFour/0600~0699/0658.Find-K-Closest-Elements.md" >}})|Medium||||43.9%|
|0696|Count Binary Substrings|[Go]({{< relref "/ChapterFour/0600~0699/0696.Count-Binary-Substrings.md" >}})|Easy||||63.6%|
|0719|Find K-th Smallest Pair Distance|[Go]({{< relref "/ChapterFour/0700~0799/0719.Find-K-th-Smallest-Pair-Distance.md" >}})|Hard||||34.1%|
|0763|Partition Labels|[Go]({{< relref "/ChapterFour/0700~0799/0763.Partition-Labels.md" >}})|Medium| O(n)| O(1)|❤️|78.5%|
|0795|Number of Subarrays with Bounded Maximum|[Go]({{< relref "/ChapterFour/0700~0799/0795.Number-of-Subarrays-with-Bounded-Maximum.md" >}})|Medium||||52.2%|
|0821|Shortest Distance to a Character|[Go]({{< relref "/ChapterFour/0800~0899/0821.Shortest-Distance-to-a-Character.md" >}})|Easy||||70.7%|
|0826|Most Profit Assigning Work|[Go]({{< relref "/ChapterFour/0800~0899/0826.Most-Profit-Assigning-Work.md" >}})|Medium| O(n log n)| O(n)||40.6%|
|0832|Flipping an Image|[Go]({{< relref "/ChapterFour/0800~0899/0832.Flipping-an-Image.md" >}})|Easy||||79.2%|
|0838|Push Dominoes|[Go]({{< relref "/ChapterFour/0800~0899/0838.Push-Dominoes.md" >}})|Medium| O(n)| O(n)||52.0%|
|0844|Backspace String Compare|[Go]({{< relref "/ChapterFour/0800~0899/0844.Backspace-String-Compare.md" >}})|Easy| O(n)| O(n) ||47.5%|
|0845|Longest Mountain in Array|[Go]({{< relref "/ChapterFour/0800~0899/0845.Longest-Mountain-in-Array.md" >}})|Medium| O(n)| O(1) ||39.4%|
|0876|Middle of the Linked List|[Go]({{< relref "/ChapterFour/0800~0899/0876.Middle-of-the-Linked-List.md" >}})|Easy||||71.0%|
|0881|Boats to Save People|[Go]({{< relref "/ChapterFour/0800~0899/0881.Boats-to-Save-People.md" >}})|Medium| O(n log n)| O(1) ||49.6%|
|0922|Sort Array By Parity II|[Go]({{< relref "/ChapterFour/0900~0999/0922.Sort-Array-By-Parity-II.md" >}})|Easy||||70.5%|
|0923|3Sum With Multiplicity|[Go]({{< relref "/ChapterFour/0900~0999/0923.3Sum-With-Multiplicity.md" >}})|Medium| O(n^2)| O(n) ||41.3%|
|0925|Long Pressed Name|[Go]({{< relref "/ChapterFour/0900~0999/0925.Long-Pressed-Name.md" >}})|Easy| O(n)| O(1)||35.3%|
|0942|DI String Match|[Go]({{< relref "/ChapterFour/0900~0999/0942.DI-String-Match.md" >}})|Easy||||75.2%|
|0969|Pancake Sorting|[Go]({{< relref "/ChapterFour/0900~0999/0969.Pancake-Sorting.md" >}})|Medium||||69.4%|
|0977|Squares of a Sorted Array|[Go]({{< relref "/ChapterFour/0900~0999/0977.Squares-of-a-Sorted-Array.md" >}})|Easy| O(n)| O(1)||71.4%|
|0986|Interval List Intersections|[Go]({{< relref "/ChapterFour/0900~0999/0986.Interval-List-Intersections.md" >}})|Medium| O(n)| O(1)||70.5%|
|1040|Moving Stones Until Consecutive II|[Go]({{< relref "/ChapterFour/1000~1099/1040.Moving-Stones-Until-Consecutive-II.md" >}})|Medium||||55.1%|
|1048|Longest String Chain|[Go]({{< relref "/ChapterFour/1000~1099/1048.Longest-String-Chain.md" >}})|Medium||||57.3%|
|1089|Duplicate Zeros|[Go]({{< relref "/ChapterFour/1000~1099/1089.Duplicate-Zeros.md" >}})|Easy||||51.1%|
|1093|Statistics from a Large Sample|[Go]({{< relref "/ChapterFour/1000~1099/1093.Statistics-from-a-Large-Sample.md" >}})|Medium| O(n)| O(1) ||46.4%|
|1332|Remove Palindromic Subsequences|[Go]({{< relref "/ChapterFour/1300~1399/1332.Remove-Palindromic-Subsequences.md" >}})|Easy||||69.2%|
|1385|Find the Distance Value Between Two Arrays|[Go]({{< relref "/ChapterFour/1300~1399/1385.Find-the-Distance-Value-Between-Two-Arrays.md" >}})|Easy||||66.1%|
|1679|Max Number of K-Sum Pairs|[Go]({{< relref "/ChapterFour/1600~1699/1679.Max-Number-of-K-Sum-Pairs.md" >}})|Medium||||53.5%|
|1721|Swapping Nodes in a Linked List|[Go]({{< relref "/ChapterFour/1700~1799/1721.Swapping-Nodes-in-a-Linked-List.md" >}})|Medium||||65.8%|
|1877|Minimize Maximum Pair Sum in Array|[Go]({{< relref "/ChapterFour/1800~1899/1877.Minimize-Maximum-Pair-Sum-in-Array.md" >}})|Medium||||80.2%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/String/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/">下一页➡️</a></p>
</div>
