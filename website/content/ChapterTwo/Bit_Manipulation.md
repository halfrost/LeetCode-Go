---
title: 2.15 ✅ Bit Manipulation
type: docs
weight: 15
---

# Bit Manipulation

![](https://img.halfrost.com/Leetcode/Bit_Manipulation.png)

- 异或的特性。第 136 题，第 268 题，第 389 题，第 421 题，

```go
x ^ 0 = x
x ^ 11111……1111 = ~x
x ^ (~x) = 11111……1111
x ^ x = 0
a ^ b = c  => a ^ c = b  => b ^ c = a (交换律)
a ^ b ^ c = a ^ (b ^ c) = (a ^ b）^ c (结合律)
```

- 构造特殊 Mask，将特殊位置放 0 或 1。

```go
将 x 最右边的 n 位清零， x & ( ~0 << n )
获取 x 的第 n 位值(0 或者 1)，(x >> n) & 1
获取 x 的第 n 位的幂值，x & (1 << (n - 1))
仅将第 n 位置为 1，x | (1 << n)
仅将第 n 位置为 0，x & (~(1 << n))
将 x 最高位至第 n 位(含)清零，x & ((1 << n) - 1)
将第 n 位至第 0 位(含)清零，x & (~((1 << (n + 1)) - 1)）
```

- 有特殊意义的 & 位操作运算。第 260 题，第 201 题，第 318 题，第 371 题，第 397 题，第 461 题，第 693 题，

```go
X & 1 == 1 判断是否是奇数(偶数)
X & = (X - 1) 将最低位(LSB)的 1 清零
X & -X 得到最低位(LSB)的 1
X & ~X = 0
```



| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0029|Divide Two Integers|[Go]({{< relref "/ChapterFour/0001~0099/0029.Divide-Two-Integers.md" >}})|Medium||||17.3%|
|0067|Add Binary|[Go]({{< relref "/ChapterFour/0001~0099/0067.Add-Binary.md" >}})|Easy||||51.3%|
|0078|Subsets|[Go]({{< relref "/ChapterFour/0001~0099/0078.Subsets.md" >}})|Medium| O(n^2)| O(n)|❤️|73.8%|
|0089|Gray Code|[Go]({{< relref "/ChapterFour/0001~0099/0089.Gray-Code.md" >}})|Medium||||56.4%|
|0090|Subsets II|[Go]({{< relref "/ChapterFour/0001~0099/0090.Subsets-II.md" >}})|Medium||||55.2%|
|0136|Single Number|[Go]({{< relref "/ChapterFour/0100~0199/0136.Single-Number.md" >}})|Easy| O(n)| O(1)||70.0%|
|0137|Single Number II|[Go]({{< relref "/ChapterFour/0100~0199/0137.Single-Number-II.md" >}})|Medium| O(n)| O(1)|❤️|57.7%|
|0187|Repeated DNA Sequences|[Go]({{< relref "/ChapterFour/0100~0199/0187.Repeated-DNA-Sequences.md" >}})|Medium| O(n)| O(1)||46.1%|
|0190|Reverse Bits|[Go]({{< relref "/ChapterFour/0100~0199/0190.Reverse-Bits.md" >}})|Easy| O(n)| O(1)|❤️|52.0%|
|0191|Number of 1 Bits|[Go]({{< relref "/ChapterFour/0100~0199/0191.Number-of-1-Bits.md" >}})|Easy| O(n)| O(1)||64.6%|
|0201|Bitwise AND of Numbers Range|[Go]({{< relref "/ChapterFour/0200~0299/0201.Bitwise-AND-of-Numbers-Range.md" >}})|Medium| O(n)| O(1)|❤️|42.2%|
|0231|Power of Two|[Go]({{< relref "/ChapterFour/0200~0299/0231.Power-of-Two.md" >}})|Easy| O(1)| O(1)||45.7%|
|0260|Single Number III|[Go]({{< relref "/ChapterFour/0200~0299/0260.Single-Number-III.md" >}})|Medium| O(n)| O(1)|❤️|67.4%|
|0268|Missing Number|[Go]({{< relref "/ChapterFour/0200~0299/0268.Missing-Number.md" >}})|Easy| O(n)| O(1)||61.5%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium||||59.1%|
|0318|Maximum Product of Word Lengths|[Go]({{< relref "/ChapterFour/0300~0399/0318.Maximum-Product-of-Word-Lengths.md" >}})|Medium| O(n)| O(1)||60.1%|
|0338|Counting Bits|[Go]({{< relref "/ChapterFour/0300~0399/0338.Counting-Bits.md" >}})|Easy| O(n)| O(n)||75.2%|
|0342|Power of Four|[Go]({{< relref "/ChapterFour/0300~0399/0342.Power-of-Four.md" >}})|Easy| O(n)| O(1)||45.6%|
|0371|Sum of Two Integers|[Go]({{< relref "/ChapterFour/0300~0399/0371.Sum-of-Two-Integers.md" >}})|Medium| O(n)| O(1)||50.7%|
|0389|Find the Difference|[Go]({{< relref "/ChapterFour/0300~0399/0389.Find-the-Difference.md" >}})|Easy| O(n)| O(1)||60.4%|
|0393|UTF-8 Validation|[Go]({{< relref "/ChapterFour/0300~0399/0393.UTF-8-Validation.md" >}})|Medium| O(n)| O(1)||45.2%|
|0397|Integer Replacement|[Go]({{< relref "/ChapterFour/0300~0399/0397.Integer-Replacement.md" >}})|Medium| O(n)| O(1)||35.1%|
|0401|Binary Watch|[Go]({{< relref "/ChapterFour/0400~0499/0401.Binary-Watch.md" >}})|Easy| O(1)| O(1)||51.5%|
|0405|Convert a Number to Hexadecimal|[Go]({{< relref "/ChapterFour/0400~0499/0405.Convert-a-Number-to-Hexadecimal.md" >}})|Easy| O(n)| O(1)||46.1%|
|0421|Maximum XOR of Two Numbers in an Array|[Go]({{< relref "/ChapterFour/0400~0499/0421.Maximum-XOR-of-Two-Numbers-in-an-Array.md" >}})|Medium| O(n)| O(1)|❤️|54.6%|
|0461|Hamming Distance|[Go]({{< relref "/ChapterFour/0400~0499/0461.Hamming-Distance.md" >}})|Easy| O(n)| O(1)||74.8%|
|0473|Matchsticks to Square|[Go]({{< relref "/ChapterFour/0400~0499/0473.Matchsticks-to-Square.md" >}})|Medium||||40.4%|
|0476|Number Complement|[Go]({{< relref "/ChapterFour/0400~0499/0476.Number-Complement.md" >}})|Easy| O(n)| O(1)||67.1%|
|0477|Total Hamming Distance|[Go]({{< relref "/ChapterFour/0400~0499/0477.Total-Hamming-Distance.md" >}})|Medium| O(n)| O(1)||52.2%|
|0491|Increasing Subsequences|[Go]({{< relref "/ChapterFour/0400~0499/0491.Increasing-Subsequences.md" >}})|Medium||||52.0%|
|0526|Beautiful Arrangement|[Go]({{< relref "/ChapterFour/0500~0599/0526.Beautiful-Arrangement.md" >}})|Medium||||64.6%|
|0638|Shopping Offers|[Go]({{< relref "/ChapterFour/0600~0699/0638.Shopping-Offers.md" >}})|Medium||||54.2%|
|0645|Set Mismatch|[Go]({{< relref "/ChapterFour/0600~0699/0645.Set-Mismatch.md" >}})|Easy||||43.0%|
|0693|Binary Number with Alternating Bits|[Go]({{< relref "/ChapterFour/0600~0699/0693.Binary-Number-with-Alternating-Bits.md" >}})|Easy| O(n)| O(1)|❤️|61.3%|
|0756|Pyramid Transition Matrix|[Go]({{< relref "/ChapterFour/0700~0799/0756.Pyramid-Transition-Matrix.md" >}})|Medium| O(n log n)| O(n)||53.3%|
|0762|Prime Number of Set Bits in Binary Representation|[Go]({{< relref "/ChapterFour/0700~0799/0762.Prime-Number-of-Set-Bits-in-Binary-Representation.md" >}})|Easy| O(n)| O(1)||67.7%|
|0784|Letter Case Permutation|[Go]({{< relref "/ChapterFour/0700~0799/0784.Letter-Case-Permutation.md" >}})|Medium| O(n)| O(1)||73.4%|
|0810|Chalkboard XOR Game|[Go]({{< relref "/ChapterFour/0800~0899/0810.Chalkboard-XOR-Game.md" >}})|Hard||||55.1%|
|0864|Shortest Path to Get All Keys|[Go]({{< relref "/ChapterFour/0800~0899/0864.Shortest-Path-to-Get-All-Keys.md" >}})|Hard||||45.4%|
|0898|Bitwise ORs of Subarrays|[Go]({{< relref "/ChapterFour/0800~0899/0898.Bitwise-ORs-of-Subarrays.md" >}})|Medium| O(n)| O(1)||36.8%|
|0980|Unique Paths III|[Go]({{< relref "/ChapterFour/0900~0999/0980.Unique-Paths-III.md" >}})|Hard||||79.6%|
|0995|Minimum Number of K Consecutive Bit Flips|[Go]({{< relref "/ChapterFour/0900~0999/0995.Minimum-Number-of-K-Consecutive-Bit-Flips.md" >}})|Hard||||51.1%|
|0996|Number of Squareful Arrays|[Go]({{< relref "/ChapterFour/0900~0999/0996.Number-of-Squareful-Arrays.md" >}})|Hard||||49.3%|
|1009|Complement of Base 10 Integer|[Go]({{< relref "/ChapterFour/1000~1099/1009.Complement-of-Base-10-Integer.md" >}})|Easy||||62.0%|
|1178|Number of Valid Words for Each Puzzle|[Go]({{< relref "/ChapterFour/1100~1199/1178.Number-of-Valid-Words-for-Each-Puzzle.md" >}})|Hard||||46.6%|
|1239|Maximum Length of a Concatenated String with Unique Characters|[Go]({{< relref "/ChapterFour/1200~1299/1239.Maximum-Length-of-a-Concatenated-String-with-Unique-Characters.md" >}})|Medium||||52.2%|
|1310|XOR Queries of a Subarray|[Go]({{< relref "/ChapterFour/1300~1399/1310.XOR-Queries-of-a-Subarray.md" >}})|Medium||||72.2%|
|1442|Count Triplets That Can Form Two Arrays of Equal XOR|[Go]({{< relref "/ChapterFour/1400~1499/1442.Count-Triplets-That-Can-Form-Two-Arrays-of-Equal-XOR.md" >}})|Medium||||75.6%|
|1461|Check If a String Contains All Binary Codes of Size K|[Go]({{< relref "/ChapterFour/1400~1499/1461.Check-If-a-String-Contains-All-Binary-Codes-of-Size-K.md" >}})|Medium||||56.8%|
|1486|XOR Operation in an Array|[Go]({{< relref "/ChapterFour/1400~1499/1486.XOR-Operation-in-an-Array.md" >}})|Easy||||84.2%|
|1655|Distribute Repeating Integers|[Go]({{< relref "/ChapterFour/1600~1699/1655.Distribute-Repeating-Integers.md" >}})|Hard||||39.6%|
|1659|Maximize Grid Happiness|[Go]({{< relref "/ChapterFour/1600~1699/1659.Maximize-Grid-Happiness.md" >}})|Hard||||38.4%|
|1680|Concatenation of Consecutive Binary Numbers|[Go]({{< relref "/ChapterFour/1600~1699/1680.Concatenation-of-Consecutive-Binary-Numbers.md" >}})|Medium||||57.0%|
|1681|Minimum Incompatibility|[Go]({{< relref "/ChapterFour/1600~1699/1681.Minimum-Incompatibility.md" >}})|Hard||||37.4%|
|1684|Count the Number of Consistent Strings|[Go]({{< relref "/ChapterFour/1600~1699/1684.Count-the-Number-of-Consistent-Strings.md" >}})|Easy||||81.8%|
|1720|Decode XORed Array|[Go]({{< relref "/ChapterFour/1700~1799/1720.Decode-XORed-Array.md" >}})|Easy||||86.0%|
|1734|Decode XORed Permutation|[Go]({{< relref "/ChapterFour/1700~1799/1734.Decode-XORed-Permutation.md" >}})|Medium||||62.3%|
|1738|Find Kth Largest XOR Coordinate Value|[Go]({{< relref "/ChapterFour/1700~1799/1738.Find-Kth-Largest-XOR-Coordinate-Value.md" >}})|Medium||||61.5%|
|1763|Longest Nice Substring|[Go]({{< relref "/ChapterFour/1700~1799/1763.Longest-Nice-Substring.md" >}})|Easy||||61.7%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Sorting/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/">下一页➡️</a></p>
</div>
