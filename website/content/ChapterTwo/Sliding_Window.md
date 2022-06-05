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
|0003|Longest Substring Without Repeating Characters|[Go]({{< relref "/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters.md" >}})|Medium| O(n)| O(1)|❤️|33.1%|
|0030|Substring with Concatenation of All Words|[Go]({{< relref "/ChapterFour/0001~0099/0030.Substring-with-Concatenation-of-All-Words.md" >}})|Hard||||28.3%|
|0076|Minimum Window Substring|[Go]({{< relref "/ChapterFour/0001~0099/0076.Minimum-Window-Substring.md" >}})|Hard| O(n)| O(n)|❤️|39.2%|
|0187|Repeated DNA Sequences|[Go]({{< relref "/ChapterFour/0100~0199/0187.Repeated-DNA-Sequences.md" >}})|Medium||||45.0%|
|0209|Minimum Size Subarray Sum|[Go]({{< relref "/ChapterFour/0200~0299/0209.Minimum-Size-Subarray-Sum.md" >}})|Medium||||43.6%|
|0219|Contains Duplicate II|[Go]({{< relref "/ChapterFour/0200~0299/0219.Contains-Duplicate-II.md" >}})|Easy||||40.7%|
|0220|Contains Duplicate III|[Go]({{< relref "/ChapterFour/0200~0299/0220.Contains-Duplicate-III.md" >}})|Medium||||21.7%|
|0239|Sliding Window Maximum|[Go]({{< relref "/ChapterFour/0200~0299/0239.Sliding-Window-Maximum.md" >}})|Hard| O(n * k)| O(n)|❤️|46.3%|
|0395|Longest Substring with At Least K Repeating Characters|[Go]({{< relref "/ChapterFour/0300~0399/0395.Longest-Substring-with-At-Least-K-Repeating-Characters.md" >}})|Medium||||44.6%|
|0424|Longest Repeating Character Replacement|[Go]({{< relref "/ChapterFour/0400~0499/0424.Longest-Repeating-Character-Replacement.md" >}})|Medium| O(n)| O(1) ||50.9%|
|0438|Find All Anagrams in a String|[Go]({{< relref "/ChapterFour/0400~0499/0438.Find-All-Anagrams-in-a-String.md" >}})|Medium||||48.4%|
|0480|Sliding Window Median|[Go]({{< relref "/ChapterFour/0400~0499/0480.Sliding-Window-Median.md" >}})|Hard| O(n * log k)| O(k)|❤️|41.1%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.5%|
|0632|Smallest Range Covering Elements from K Lists|[Go]({{< relref "/ChapterFour/0600~0699/0632.Smallest-Range-Covering-Elements-from-K-Lists.md" >}})|Hard||||58.8%|
|0643|Maximum Average Subarray I|[Go]({{< relref "/ChapterFour/0600~0699/0643.Maximum-Average-Subarray-I.md" >}})|Easy||||43.6%|
|0713|Subarray Product Less Than K|[Go]({{< relref "/ChapterFour/0700~0799/0713.Subarray-Product-Less-Than-K.md" >}})|Medium||||44.0%|
|0718|Maximum Length of Repeated Subarray|[Go]({{< relref "/ChapterFour/0700~0799/0718.Maximum-Length-of-Repeated-Subarray.md" >}})|Medium||||51.4%|
|0862|Shortest Subarray with Sum at Least K|[Go]({{< relref "/ChapterFour/0800~0899/0862.Shortest-Subarray-with-Sum-at-Least-K.md" >}})|Hard||||26.2%|
|0904|Fruit Into Baskets|[Go]({{< relref "/ChapterFour/0900~0999/0904.Fruit-Into-Baskets.md" >}})|Medium||||42.7%|
|0930|Binary Subarrays With Sum|[Go]({{< relref "/ChapterFour/0900~0999/0930.Binary-Subarrays-With-Sum.md" >}})|Medium||||49.1%|
|0978|Longest Turbulent Subarray|[Go]({{< relref "/ChapterFour/0900~0999/0978.Longest-Turbulent-Subarray.md" >}})|Medium| O(n)| O(1)|❤️|47.5%|
|0992|Subarrays with K Different Integers|[Go]({{< relref "/ChapterFour/0900~0999/0992.Subarrays-with-K-Different-Integers.md" >}})|Hard| O(n)| O(n)|❤️|53.5%|
|0995|Minimum Number of K Consecutive Bit Flips|[Go]({{< relref "/ChapterFour/0900~0999/0995.Minimum-Number-of-K-Consecutive-Bit-Flips.md" >}})|Hard| O(n)| O(1)|❤️|50.9%|
|1004|Max Consecutive Ones III|[Go]({{< relref "/ChapterFour/1000~1099/1004.Max-Consecutive-Ones-III.md" >}})|Medium| O(n)| O(1) ||63.1%|
|1052|Grumpy Bookstore Owner|[Go]({{< relref "/ChapterFour/1000~1099/1052.Grumpy-Bookstore-Owner.md" >}})|Medium| O(n log n)| O(1) ||56.8%|
|1208|Get Equal Substrings Within Budget|[Go]({{< relref "/ChapterFour/1200~1299/1208.Get-Equal-Substrings-Within-Budget.md" >}})|Medium||||46.7%|
|1234|Replace the Substring for Balanced String|[Go]({{< relref "/ChapterFour/1200~1299/1234.Replace-the-Substring-for-Balanced-String.md" >}})|Medium||||36.3%|
|1423|Maximum Points You Can Obtain from Cards|[Go]({{< relref "/ChapterFour/1400~1499/1423.Maximum-Points-You-Can-Obtain-from-Cards.md" >}})|Medium||||50.1%|
|1438|Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit|[Go]({{< relref "/ChapterFour/1400~1499/1438.Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit.md" >}})|Medium||||46.8%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1600~1699/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||33.6%|
|1695|Maximum Erasure Value|[Go]({{< relref "/ChapterFour/1600~1699/1695.Maximum-Erasure-Value.md" >}})|Medium||||52.5%|
|1696|Jump Game VI|[Go]({{< relref "/ChapterFour/1600~1699/1696.Jump-Game-VI.md" >}})|Medium||||42.8%|
|1763|Longest Nice Substring|[Go]({{< relref "/ChapterFour/1700~1799/1763.Longest-Nice-Substring.md" >}})|Easy||||61.9%|
|1984|Minimum Difference Between Highest and Lowest of K Scores|[Go]({{< relref "/ChapterFour/1900~1999/1984.Minimum-Difference-Between-Highest-and-Lowest-of-K-Scores.md" >}})|Easy||||53.4%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Segment_Tree/">下一页➡️</a></p>
</div>
