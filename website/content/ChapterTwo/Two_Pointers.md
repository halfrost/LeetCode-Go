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
|0011|Container With Most Water|[Go]({{< relref "/ChapterFour/0001~0099/0011.Container-With-Most-Water.md" >}})|Medium| O(n)| O(1)||53.1%|
|0015|3Sum|[Go]({{< relref "/ChapterFour/0001~0099/0015.3Sum.md" >}})|Medium| O(n^2)| O(n)|❤️|28.9%|
|0016|3Sum Closest|[Go]({{< relref "/ChapterFour/0001~0099/0016.3Sum-Closest.md" >}})|Medium| O(n^2)| O(1)|❤️|46.5%|
|0018|4Sum|[Go]({{< relref "/ChapterFour/0001~0099/0018.4Sum.md" >}})|Medium| O(n^3)| O(n^2)|❤️|35.8%|
|0019|Remove Nth Node From End of List|[Go]({{< relref "/ChapterFour/0001~0099/0019.Remove-Nth-Node-From-End-of-List.md" >}})|Medium| O(n)| O(1)||36.5%|
|0026|Remove Duplicates from Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array.md" >}})|Easy| O(n)| O(1)||47.3%|
|0027|Remove Element|[Go]({{< relref "/ChapterFour/0001~0099/0027.Remove-Element.md" >}})|Easy| O(n)| O(1)||49.9%|
|0028|Implement strStr()|[Go]({{< relref "/ChapterFour/0001~0099/0028.Implement-strStr.md" >}})|Easy| O(n)| O(1)||35.7%|
|0031|Next Permutation|[Go]({{< relref "/ChapterFour/0001~0099/0031.Next-Permutation.md" >}})|Medium||||34.3%|
|0042|Trapping Rain Water|[Go]({{< relref "/ChapterFour/0001~0099/0042.Trapping-Rain-Water.md" >}})|Hard| O(n)| O(1)|❤️|52.5%|
|0061|Rotate List|[Go]({{< relref "/ChapterFour/0001~0099/0061.Rotate-List.md" >}})|Medium| O(n)| O(1)||32.4%|
|0075|Sort Colors|[Go]({{< relref "/ChapterFour/0001~0099/0075.Sort-Colors.md" >}})|Medium| O(n)| O(1)|❤️|50.9%|
|0080|Remove Duplicates from Sorted Array II|[Go]({{< relref "/ChapterFour/0001~0099/0080.Remove-Duplicates-from-Sorted-Array-II.md" >}})|Medium| O(n)| O(1||47.0%|
|0082|Remove Duplicates from Sorted List II|[Go]({{< relref "/ChapterFour/0001~0099/0082.Remove-Duplicates-from-Sorted-List-II.md" >}})|Medium||||40.1%|
|0086|Partition List|[Go]({{< relref "/ChapterFour/0001~0099/0086.Partition-List.md" >}})|Medium| O(n)| O(1)|❤️|45.6%|
|0088|Merge Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0088.Merge-Sorted-Array.md" >}})|Easy| O(n)| O(1)|❤️|41.3%|
|0125|Valid Palindrome|[Go]({{< relref "/ChapterFour/0100~0199/0125.Valid-Palindrome.md" >}})|Easy| O(n)| O(1)||39.2%|
|0141|Linked List Cycle|[Go]({{< relref "/ChapterFour/0100~0199/0141.Linked-List-Cycle.md" >}})|Easy| O(n)| O(1)|❤️|43.6%|
|0142|Linked List Cycle II|[Go]({{< relref "/ChapterFour/0100~0199/0142.Linked-List-Cycle-II.md" >}})|Medium| O(n)| O(1)|❤️|40.8%|
|0143|Reorder List|[Go]({{< relref "/ChapterFour/0100~0199/0143.Reorder-List.md" >}})|Medium||||42.3%|
|0148|Sort List|[Go]({{< relref "/ChapterFour/0100~0199/0148.Sort-List.md" >}})|Medium||||47.7%|
|0151|Reverse Words in a String|[Go]({{< relref "/ChapterFour/0100~0199/0151.Reverse-Words-in-a-String.md" >}})|Medium||||25.0%|
|0160|Intersection of Two Linked Lists|[Go]({{< relref "/ChapterFour/0100~0199/0160.Intersection-of-Two-Linked-Lists.md" >}})|Easy||||45.9%|
|0167|Two Sum II - Input array is sorted|[Go]({{< relref "/ChapterFour/0100~0199/0167.Two-Sum-II-Input-array-is-sorted.md" >}})|Easy| O(n)| O(1)||56.3%|
|0189|Rotate Array|[Go]({{< relref "/ChapterFour/0100~0199/0189.Rotate-Array.md" >}})|Medium||||36.9%|
|0202|Happy Number|[Go]({{< relref "/ChapterFour/0200~0299/0202.Happy-Number.md" >}})|Easy||||51.7%|
|0234|Palindrome Linked List|[Go]({{< relref "/ChapterFour/0200~0299/0234.Palindrome-Linked-List.md" >}})|Easy| O(n)| O(1)||43.3%|
|0283|Move Zeroes|[Go]({{< relref "/ChapterFour/0200~0299/0283.Move-Zeroes.md" >}})|Easy| O(n)| O(1)||59.0%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium| O(n)| O(1)|❤️|58.1%|
|0344|Reverse String|[Go]({{< relref "/ChapterFour/0300~0399/0344.Reverse-String.md" >}})|Easy| O(n)| O(1)||71.4%|
|0345|Reverse Vowels of a String|[Go]({{< relref "/ChapterFour/0300~0399/0345.Reverse-Vowels-of-a-String.md" >}})|Easy| O(n)| O(1)||45.7%|
|0349|Intersection of Two Arrays|[Go]({{< relref "/ChapterFour/0300~0399/0349.Intersection-of-Two-Arrays.md" >}})|Easy| O(n)| O(n) ||66.3%|
|0350|Intersection of Two Arrays II|[Go]({{< relref "/ChapterFour/0300~0399/0350.Intersection-of-Two-Arrays-II.md" >}})|Easy| O(n)| O(n) ||52.6%|
|0392|Is Subsequence|[Go]({{< relref "/ChapterFour/0300~0399/0392.Is-Subsequence.md" >}})|Easy||||49.8%|
|0457|Circular Array Loop|[Go]({{< relref "/ChapterFour/0400~0499/0457.Circular-Array-Loop.md" >}})|Medium||||30.6%|
|0524|Longest Word in Dictionary through Deleting|[Go]({{< relref "/ChapterFour/0500~0599/0524.Longest-Word-in-Dictionary-through-Deleting.md" >}})|Medium| O(n)| O(1) ||50.3%|
|0532|K-diff Pairs in an Array|[Go]({{< relref "/ChapterFour/0500~0599/0532.K-diff-Pairs-in-an-Array.md" >}})|Medium| O(n)| O(n)||36.1%|
|0541|Reverse String II|[Go]({{< relref "/ChapterFour/0500~0599/0541.Reverse-String-II.md" >}})|Easy||||49.8%|
|0557|Reverse Words in a String III|[Go]({{< relref "/ChapterFour/0500~0599/0557.Reverse-Words-in-a-String-III.md" >}})|Easy||||73.2%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.6%|
|0581|Shortest Unsorted Continuous Subarray|[Go]({{< relref "/ChapterFour/0500~0599/0581.Shortest-Unsorted-Continuous-Subarray.md" >}})|Medium||||33.2%|
|0633|Sum of Square Numbers|[Go]({{< relref "/ChapterFour/0600~0699/0633.Sum-of-Square-Numbers.md" >}})|Medium||||32.9%|
|0653|Two Sum IV - Input is a BST|[Go]({{< relref "/ChapterFour/0600~0699/0653.Two-Sum-IV-Input-is-a-BST.md" >}})|Easy||||56.7%|
|0658|Find K Closest Elements|[Go]({{< relref "/ChapterFour/0600~0699/0658.Find-K-Closest-Elements.md" >}})|Medium||||43.4%|
|0696|Count Binary Substrings|[Go]({{< relref "/ChapterFour/0600~0699/0696.Count-Binary-Substrings.md" >}})|Easy||||61.6%|
|0719|Find K-th Smallest Pair Distance|[Go]({{< relref "/ChapterFour/0700~0799/0719.Find-K-th-Smallest-Pair-Distance.md" >}})|Hard||||33.0%|
|0763|Partition Labels|[Go]({{< relref "/ChapterFour/0700~0799/0763.Partition-Labels.md" >}})|Medium| O(n)| O(1)|❤️|78.3%|
|0795|Number of Subarrays with Bounded Maximum|[Go]({{< relref "/ChapterFour/0700~0799/0795.Number-of-Subarrays-with-Bounded-Maximum.md" >}})|Medium||||51.7%|
|0821|Shortest Distance to a Character|[Go]({{< relref "/ChapterFour/0800~0899/0821.Shortest-Distance-to-a-Character.md" >}})|Easy||||70.3%|
|0826|Most Profit Assigning Work|[Go]({{< relref "/ChapterFour/0800~0899/0826.Most-Profit-Assigning-Work.md" >}})|Medium| O(n log n)| O(n)||39.6%|
|0832|Flipping an Image|[Go]({{< relref "/ChapterFour/0800~0899/0832.Flipping-an-Image.md" >}})|Easy||||78.7%|
|0838|Push Dominoes|[Go]({{< relref "/ChapterFour/0800~0899/0838.Push-Dominoes.md" >}})|Medium| O(n)| O(n)||50.6%|
|0844|Backspace String Compare|[Go]({{< relref "/ChapterFour/0800~0899/0844.Backspace-String-Compare.md" >}})|Easy| O(n)| O(n) ||47.2%|
|0845|Longest Mountain in Array|[Go]({{< relref "/ChapterFour/0800~0899/0845.Longest-Mountain-in-Array.md" >}})|Medium| O(n)| O(1) ||38.9%|
|0876|Middle of the Linked List|[Go]({{< relref "/ChapterFour/0800~0899/0876.Middle-of-the-Linked-List.md" >}})|Easy||||69.5%|
|0881|Boats to Save People|[Go]({{< relref "/ChapterFour/0800~0899/0881.Boats-to-Save-People.md" >}})|Medium| O(n log n)| O(1) ||49.3%|
|0922|Sort Array By Parity II|[Go]({{< relref "/ChapterFour/0900~0999/0922.Sort-Array-By-Parity-II.md" >}})|Easy||||70.7%|
|0923|3Sum With Multiplicity|[Go]({{< relref "/ChapterFour/0900~0999/0923.3Sum-With-Multiplicity.md" >}})|Medium| O(n^2)| O(n) ||41.0%|
|0925|Long Pressed Name|[Go]({{< relref "/ChapterFour/0900~0999/0925.Long-Pressed-Name.md" >}})|Easy| O(n)| O(1)||36.2%|
|0942|DI String Match|[Go]({{< relref "/ChapterFour/0900~0999/0942.DI-String-Match.md" >}})|Easy||||74.3%|
|0969|Pancake Sorting|[Go]({{< relref "/ChapterFour/0900~0999/0969.Pancake-Sorting.md" >}})|Medium||||69.0%|
|0977|Squares of a Sorted Array|[Go]({{< relref "/ChapterFour/0900~0999/0977.Squares-of-a-Sorted-Array.md" >}})|Easy| O(n)| O(1)||71.5%|
|0986|Interval List Intersections|[Go]({{< relref "/ChapterFour/0900~0999/0986.Interval-List-Intersections.md" >}})|Medium| O(n)| O(1)||69.0%|
|1040|Moving Stones Until Consecutive II|[Go]({{< relref "/ChapterFour/1000~1099/1040.Moving-Stones-Until-Consecutive-II.md" >}})|Medium||||54.8%|
|1048|Longest String Chain|[Go]({{< relref "/ChapterFour/1000~1099/1048.Longest-String-Chain.md" >}})|Medium||||56.3%|
|1089|Duplicate Zeros|[Go]({{< relref "/ChapterFour/1000~1099/1089.Duplicate-Zeros.md" >}})|Easy||||51.5%|
|1093|Statistics from a Large Sample|[Go]({{< relref "/ChapterFour/1000~1099/1093.Statistics-from-a-Large-Sample.md" >}})|Medium| O(n)| O(1) ||47.6%|
|1332|Remove Palindromic Subsequences|[Go]({{< relref "/ChapterFour/1300~1399/1332.Remove-Palindromic-Subsequences.md" >}})|Easy||||68.7%|
|1385|Find the Distance Value Between Two Arrays|[Go]({{< relref "/ChapterFour/1300~1399/1385.Find-the-Distance-Value-Between-Two-Arrays.md" >}})|Easy||||66.5%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1600~1699/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||33.3%|
|1679|Max Number of K-Sum Pairs|[Go]({{< relref "/ChapterFour/1600~1699/1679.Max-Number-of-K-Sum-Pairs.md" >}})|Medium||||53.6%|
|1721|Swapping Nodes in a Linked List|[Go]({{< relref "/ChapterFour/1700~1799/1721.Swapping-Nodes-in-a-Linked-List.md" >}})|Medium||||66.5%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/String/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/">下一页➡️</a></p>
</div>
