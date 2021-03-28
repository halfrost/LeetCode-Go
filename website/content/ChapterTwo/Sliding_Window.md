---
title: 2.17 ✅ Sliding Window
type: docs
weight: 17
---

# Sliding Window

![](https://img.halfrost.com/Leetcode/Sliding_Window.png)

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
- 滑动窗口经典题。第 239 题，第 480 题。


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0003|Longest Substring Without Repeating Characters|[Go]({{< relref "/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters.md" >}})|Medium| O(n)| O(1)|❤️|31.5%|
|0076|Minimum Window Substring|[Go]({{< relref "/ChapterFour/0001~0099/0076.Minimum-Window-Substring.md" >}})|Hard| O(n)| O(n)|❤️|36.1%|
|0239|Sliding Window Maximum|[Go]({{< relref "/ChapterFour/0200~0299/0239.Sliding-Window-Maximum.md" >}})|Hard| O(n * k)| O(n)|❤️|44.7%|
|0395|Longest Substring with At Least K Repeating Characters|[Go]({{< relref "/ChapterFour/0300~0399/0395.Longest-Substring-with-At-Least-K-Repeating-Characters.md" >}})|Medium||||43.6%|
|0424|Longest Repeating Character Replacement|[Go]({{< relref "/ChapterFour/0400~0499/0424.Longest-Repeating-Character-Replacement.md" >}})|Medium| O(n)| O(1) ||48.3%|
|0480|Sliding Window Median|[Go]({{< relref "/ChapterFour/0400~0499/0480.Sliding-Window-Median.md" >}})|Hard| O(n * log k)| O(k)|❤️|38.8%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.6%|
|0978|Longest Turbulent Subarray|[Go]({{< relref "/ChapterFour/0900~0999/0978.Longest-Turbulent-Subarray.md" >}})|Medium| O(n)| O(1)|❤️|46.7%|
|0992|Subarrays with K Different Integers|[Go]({{< relref "/ChapterFour/0900~0999/0992.Subarrays-with-K-Different-Integers.md" >}})|Hard| O(n)| O(n)|❤️|50.8%|
|0995|Minimum Number of K Consecutive Bit Flips|[Go]({{< relref "/ChapterFour/0900~0999/0995.Minimum-Number-of-K-Consecutive-Bit-Flips.md" >}})|Hard| O(n)| O(1)|❤️|50.0%|
|1004|Max Consecutive Ones III|[Go]({{< relref "/ChapterFour/1000~1099/1004.Max-Consecutive-Ones-III.md" >}})|Medium| O(n)| O(1) ||60.9%|
|1040|Moving Stones Until Consecutive II|[Go]({{< relref "/ChapterFour/1000~1099/1040.Moving-Stones-Until-Consecutive-II.md" >}})|Medium| O(n log n)| O(1) |❤️|54.2%|
|1052|Grumpy Bookstore Owner|[Go]({{< relref "/ChapterFour/1000~1099/1052.Grumpy-Bookstore-Owner.md" >}})|Medium| O(n log n)| O(1) ||55.9%|
|1074|Number of Submatrices That Sum to Target|[Go]({{< relref "/ChapterFour/1000~1099/1074.Number-of-Submatrices-That-Sum-to-Target.md" >}})|Hard| O(n^3)| O(n) |❤️|62.0%|
|1208|Get Equal Substrings Within Budget|[Go]({{< relref "/ChapterFour/1200~1299/1208.Get-Equal-Substrings-Within-Budget.md" >}})|Medium||||44.2%|
|1423|Maximum Points You Can Obtain from Cards|[Go]({{< relref "/ChapterFour/1400~1499/1423.Maximum-Points-You-Can-Obtain-from-Cards.md" >}})|Medium||||46.8%|
|1438|Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit|[Go]({{< relref "/ChapterFour/1400~1499/1438.Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit.md" >}})|Medium||||44.4%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1600~1699/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||33.3%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Segment_Tree/">下一页➡️</a></p>
</div>
