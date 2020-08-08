---
title: Bit Manipulation
type: docs
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


| Title | Solution | Difficulty | Time | Space | 收藏 |
| ----- | :--------: | :----------: | :----: | :-----: |:-----: |
|78. Subsets | [Go]({{< relref "/ChapterFour/0078.Subsets.md" >}})| Medium | O(n^2)| O(n)|❤️|
|136. Single Number | [Go]({{< relref "/ChapterFour/0136.Single-Number.md" >}})| Easy | O(n)| O(1)||
|137. Single Number II  | [Go]({{< relref "/ChapterFour/0137.Single-Number-II.md" >}})| Medium | O(n)| O(1)|❤️|
|169. Majority Element | [Go]({{< relref "/ChapterFour/0169.Majority-Element.md" >}})| Easy | O(n)| O(1)|❤️|
|187. Repeated DNA Sequences  | [Go]({{< relref "/ChapterFour/0187.Repeated-DNA-Sequences.md" >}})| Medium | O(n)| O(1)||
|190. Reverse Bits | [Go]({{< relref "/ChapterFour/0190.Reverse-Bits.md" >}})| Easy | O(n)| O(1)|❤️|
|191. Number of 1 Bits  | [Go]({{< relref "/ChapterFour/0191.Number-of-1-Bits.md" >}})| Easy | O(n)| O(1)||
|201. Bitwise AND of Numbers Range  | [Go]({{< relref "/ChapterFour/0201.Bitwise-AND-of-Numbers-Range.md" >}})| Medium | O(n)| O(1)|❤️|
|231. Power of Two  | [Go]({{< relref "/ChapterFour/0231.Power-of-Two.md" >}})| Easy | O(1)| O(1)||
|260. Single Number III  | [Go]({{< relref "/ChapterFour/0260.Single-Number-III.md" >}})| Medium | O(n)| O(1)|❤️|
|268. Missing Number  | [Go]({{< relref "/ChapterFour/0268.Missing-Number.md" >}})| Easy | O(n)| O(1)||
|318. Maximum Product of Word Lengths | [Go]({{< relref "/ChapterFour/0318.Maximum-Product-of-Word-Lengths.md" >}})| Medium | O(n)| O(1)||
|338. Counting Bits   | [Go]({{< relref "/ChapterFour/0338.Counting-Bits.md" >}})| Medium | O(n)| O(n)||
|342. Power of Four  | [Go]({{< relref "/ChapterFour/0342.Power-of-Four.md" >}})| Easy | O(n)| O(1)||
|371. Sum of Two Integers  | [Go]({{< relref "/ChapterFour/0371.Sum-of-Two-Integers.md" >}})| Easy | O(n)| O(1)||
|389. Find the Difference  | [Go]({{< relref "/ChapterFour/0389.Find-the-Difference.md" >}})| Easy | O(n)| O(1)||
|393. UTF-8 Validation   | [Go]({{< relref "/ChapterFour/0393.UTF-8-Validation.md" >}})| Medium | O(n)| O(1)||
|397. Integer Replacement   | [Go]({{< relref "/ChapterFour/0397.Integer-Replacement.md" >}})| Medium | O(n)| O(1)||
|401. Binary Watch  | [Go]({{< relref "/ChapterFour/0401.Binary-Watch.md" >}})| Easy | O(1)| O(1)||
|405. Convert a Number to Hexadecimal | [Go]({{< relref "/ChapterFour/0405.Convert-a-Number-to-Hexadecimal.md" >}})| Easy | O(n)| O(1)||
|421. Maximum XOR of Two Numbers in an Array  | [Go]({{< relref "/ChapterFour/0421.Maximum-XOR-of-Two-Numbers-in-an-Array.md" >}})| Medium | O(n)| O(1)|❤️|
|461. Hamming Distance  | [Go]({{< relref "/ChapterFour/0461.Hamming-Distance.md" >}})| Easy | O(n)| O(1)||
|476. Number Complement  | [Go]({{< relref "/ChapterFour/0476.Number-Complement.md" >}})| Easy | O(n)| O(1)||
|477. Total Hamming Distance  | [Go]({{< relref "/ChapterFour/0477.Total-Hamming-Distance.md" >}})| Medium | O(n)| O(1)||
|693. Binary Number with Alternating Bits | [Go]({{< relref "/ChapterFour/0693.Binary-Number-with-Alternating-Bits.md" >}})| Easy | O(n)| O(1)|❤️|
|756. Pyramid Transition Matrix  | [Go]({{< relref "/ChapterFour/0756.Pyramid-Transition-Matrix.md" >}})| Medium | O(n log n)| O(n)||
|762. Prime Number of Set Bits in Binary Representation | [Go]({{< relref "/ChapterFour/0762.Prime-Number-of-Set-Bits-in-Binary-Representation.md" >}})| Easy | O(n)| O(1)||
|784. Letter Case Permutation | [Go]({{< relref "/ChapterFour/0784.Letter-Case-Permutation.md" >}})| Easy | O(n)| O(1)||
|898. Bitwise ORs of Subarrays | [Go]({{< relref "/ChapterFour/0898.Bitwise-ORs-of-Subarrays.md" >}})| Medium | O(n)| O(1)||
|---------------------------------------|---------------------------------|--------------------------|-----------------------|-----------|--------|