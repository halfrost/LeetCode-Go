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
|0003|Longest Substring Without Repeating Characters|[Go]({{< relref "/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters.md" >}})|Medium| O(n)| O(1)|❤️|31.4%|
|0011|Container With Most Water|[Go]({{< relref "/ChapterFour/0001~0099/0011.Container-With-Most-Water.md" >}})|Medium| O(n)| O(1)||52.4%|
|0015|3Sum|[Go]({{< relref "/ChapterFour/0001~0099/0015.3Sum.md" >}})|Medium| O(n^2)| O(n)|❤️|27.9%|
|0016|3Sum Closest|[Go]({{< relref "/ChapterFour/0001~0099/0016.3Sum-Closest.md" >}})|Medium| O(n^2)| O(1)|❤️|46.3%|
|0018|4Sum|[Go]({{< relref "/ChapterFour/0001~0099/0018.4Sum.md" >}})|Medium| O(n^3)| O(n^2)|❤️|34.8%|
|0019|Remove Nth Node From End of List|[Go]({{< relref "/ChapterFour/0001~0099/0019.Remove-Nth-Node-From-End-of-List.md" >}})|Medium| O(n)| O(1)||35.7%|
|0026|Remove Duplicates from Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array.md" >}})|Easy| O(n)| O(1)||46.5%|
|0027|Remove Element|[Go]({{< relref "/ChapterFour/0001~0099/0027.Remove-Element.md" >}})|Easy| O(n)| O(1)||49.2%|
|0028|Implement strStr()|[Go]({{< relref "/ChapterFour/0001~0099/0028.Implement-strStr.md" >}})|Easy| O(n)| O(1)||35.1%|
|0030|Substring with Concatenation of All Words|[Go]({{< relref "/ChapterFour/0001~0099/0030.Substring-with-Concatenation-of-All-Words.md" >}})|Hard| O(n)| O(n)|❤️|26.1%|
|0042|Trapping Rain Water|[Go]({{< relref "/ChapterFour/0001~0099/0042.Trapping-Rain-Water.md" >}})|Hard| O(n)| O(1)|❤️|51.0%|
|0061|Rotate List|[Go]({{< relref "/ChapterFour/0001~0099/0061.Rotate-List.md" >}})|Medium| O(n)| O(1)||31.7%|
|0075|Sort Colors|[Go]({{< relref "/ChapterFour/0001~0099/0075.Sort-Colors.md" >}})|Medium| O(n)| O(1)|❤️|49.2%|
|0076|Minimum Window Substring|[Go]({{< relref "/ChapterFour/0001~0099/0076.Minimum-Window-Substring.md" >}})|Hard| O(n)| O(n)|❤️|35.9%|
|0080|Remove Duplicates from Sorted Array II|[Go]({{< relref "/ChapterFour/0001~0099/0080.Remove-Duplicates-from-Sorted-Array-II.md" >}})|Medium| O(n)| O(1||46.0%|
|0086|Partition List|[Go]({{< relref "/ChapterFour/0001~0099/0086.Partition-List.md" >}})|Medium| O(n)| O(1)|❤️|43.3%|
|0088|Merge Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0088.Merge-Sorted-Array.md" >}})|Easy| O(n)| O(1)|❤️|40.7%|
|0125|Valid Palindrome|[Go]({{< relref "/ChapterFour/0100~0199/0125.Valid-Palindrome.md" >}})|Easy| O(n)| O(1)||38.1%|
|0141|Linked List Cycle|[Go]({{< relref "/ChapterFour/0100~0199/0141.Linked-List-Cycle.md" >}})|Easy| O(n)| O(1)|❤️|42.8%|
|0142|Linked List Cycle II|[Go]({{< relref "/ChapterFour/0100~0199/0142.Linked-List-Cycle-II.md" >}})|Medium| O(n)| O(1)|❤️|39.6%|
|0167|Two Sum II - Input array is sorted|[Go]({{< relref "/ChapterFour/0100~0199/0167.Two-Sum-II---Input-array-is-sorted.md" >}})|Easy| O(n)| O(1)||55.6%|
|0209|Minimum Size Subarray Sum|[Go]({{< relref "/ChapterFour/0200~0299/0209.Minimum-Size-Subarray-Sum.md" >}})|Medium| O(n)| O(1)||39.4%|
|0234|Palindrome Linked List|[Go]({{< relref "/ChapterFour/0200~0299/0234.Palindrome-Linked-List.md" >}})|Easy| O(n)| O(1)||40.4%|
|0283|Move Zeroes|[Go]({{< relref "/ChapterFour/0200~0299/0283.Move-Zeroes.md" >}})|Easy| O(n)| O(1)||58.5%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium| O(n)| O(1)|❤️|57.4%|
|0344|Reverse String|[Go]({{< relref "/ChapterFour/0300~0399/0344.Reverse-String.md" >}})|Easy| O(n)| O(1)||70.2%|
|0345|Reverse Vowels of a String|[Go]({{< relref "/ChapterFour/0300~0399/0345.Reverse-Vowels-of-a-String.md" >}})|Easy| O(n)| O(1)||45.0%|
|0349|Intersection of Two Arrays|[Go]({{< relref "/ChapterFour/0300~0399/0349.Intersection-of-Two-Arrays.md" >}})|Easy| O(n)| O(n) ||64.7%|
|0350|Intersection of Two Arrays II|[Go]({{< relref "/ChapterFour/0300~0399/0350.Intersection-of-Two-Arrays-II.md" >}})|Easy| O(n)| O(n) ||52.0%|
|0424|Longest Repeating Character Replacement|[Go]({{< relref "/ChapterFour/0400~0499/0424.Longest-Repeating-Character-Replacement.md" >}})|Medium| O(n)| O(1) ||48.2%|
|0457|Circular Array Loop|[Go]({{< relref "/ChapterFour/0400~0499/0457.Circular-Array-Loop.md" >}})|Medium||||30.0%|
|0524|Longest Word in Dictionary through Deleting|[Go]({{< relref "/ChapterFour/0500~0599/0524.Longest-Word-in-Dictionary-through-Deleting.md" >}})|Medium| O(n)| O(1) ||49.0%|
|0532|K-diff Pairs in an Array|[Go]({{< relref "/ChapterFour/0500~0599/0532.K-diff-Pairs-in-an-Array.md" >}})|Medium| O(n)| O(n)||35.1%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.6%|
|0632|Smallest Range Covering Elements from K Lists|[Go]({{< relref "/ChapterFour/0600~0699/0632.Smallest-Range-Covering-Elements-from-K-Lists.md" >}})|Hard||||54.1%|
|0713|Subarray Product Less Than K|[Go]({{< relref "/ChapterFour/0700~0799/0713.Subarray-Product-Less-Than-K.md" >}})|Medium| O(n)| O(1)||40.5%|
|0763|Partition Labels|[Go]({{< relref "/ChapterFour/0700~0799/0763.Partition-Labels.md" >}})|Medium| O(n)| O(1)|❤️|78.0%|
|0826|Most Profit Assigning Work|[Go]({{< relref "/ChapterFour/0800~0899/0826.Most-Profit-Assigning-Work.md" >}})|Medium| O(n log n)| O(n)||39.0%|
|0828|Count Unique Characters of All Substrings of a Given String|[Go]({{< relref "/ChapterFour/0800~0899/0828.Count-Unique-Characters-of-All-Substrings-of-a-Given-String.md" >}})|Hard| O(n)| O(1)|❤️|46.8%|
|0838|Push Dominoes|[Go]({{< relref "/ChapterFour/0800~0899/0838.Push-Dominoes.md" >}})|Medium| O(n)| O(n)||49.8%|
|0844|Backspace String Compare|[Go]({{< relref "/ChapterFour/0800~0899/0844.Backspace-String-Compare.md" >}})|Easy| O(n)| O(n) ||46.9%|
|0845|Longest Mountain in Array|[Go]({{< relref "/ChapterFour/0800~0899/0845.Longest-Mountain-in-Array.md" >}})|Medium| O(n)| O(1) ||38.5%|
|0881|Boats to Save People|[Go]({{< relref "/ChapterFour/0800~0899/0881.Boats-to-Save-People.md" >}})|Medium| O(n log n)| O(1) ||48.8%|
|0904|Fruit Into Baskets|[Go]({{< relref "/ChapterFour/0900~0999/0904.Fruit-Into-Baskets.md" >}})|Medium| O(n log n)| O(1) ||43.0%|
|0923|3Sum With Multiplicity|[Go]({{< relref "/ChapterFour/0900~0999/0923.3Sum-With-Multiplicity.md" >}})|Medium| O(n^2)| O(n) ||36.2%|
|0925|Long Pressed Name|[Go]({{< relref "/ChapterFour/0900~0999/0925.Long-Pressed-Name.md" >}})|Easy| O(n)| O(1)||38.1%|
|0930|Binary Subarrays With Sum|[Go]({{< relref "/ChapterFour/0900~0999/0930.Binary-Subarrays-With-Sum.md" >}})|Medium| O(n)| O(n)  |❤️|44.5%|
|0977|Squares of a Sorted Array|[Go]({{< relref "/ChapterFour/0900~0999/0977.Squares-of-a-Sorted-Array.md" >}})|Easy| O(n)| O(1)||72.1%|
|0986|Interval List Intersections|[Go]({{< relref "/ChapterFour/0900~0999/0986.Interval-List-Intersections.md" >}})|Medium| O(n)| O(1)||68.3%|
|0992|Subarrays with K Different Integers|[Go]({{< relref "/ChapterFour/0900~0999/0992.Subarrays-with-K-Different-Integers.md" >}})|Hard| O(n)| O(n)|❤️|50.6%|
|1004|Max Consecutive Ones III|[Go]({{< relref "/ChapterFour/1000~1099/1004.Max-Consecutive-Ones-III.md" >}})|Medium| O(n)| O(1) ||60.7%|
|1093|Statistics from a Large Sample|[Go]({{< relref "/ChapterFour/1000~1099/1093.Statistics-from-a-Large-Sample.md" >}})|Medium| O(n)| O(1) ||49.3%|
|1234|Replace the Substring for Balanced String|[Go]({{< relref "/ChapterFour/1200~1299/1234.Replace-the-Substring-for-Balanced-String.md" >}})|Medium||||34.3%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1600~1699/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||33.4%|
|1695|Maximum Erasure Value|[Go]({{< relref "/ChapterFour/1600~1699/1695.Maximum-Erasure-Value.md" >}})|Medium||||49.6%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/String/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/">下一页➡️</a></p>
</div>
