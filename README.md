
# LeetCode by Go
[LeetCode Online Judge] (https://leetcode.com/) is a website containing many **algorithm questions**. Most of them are real interview questions of **Google, Facebook, LinkedIn, Apple**, etc. This repo shows my solutions by Go with the code style strictly follows the [Google Golang Style Guide](https://github.com/golang/go/wiki/CodeReviewComments). Please feel free to reference and **STAR** to support this repo, thank you!


<p align='center'>
<img src='./logo.png'>
</p>


## Data Structures

> 标识了 (已全部做完) 的专题是完成所有题目了的，没有标识的是还没有做完所有题目的

* [Array](#array)
* [String](#string)
* [Linked List](#linked-list-已全部做完)
* [Stack](#stack-已全部做完)
* [Tree](#tree)
* [Dynamic programming](#dynamic-programming)
* [Depth-first search](#depth-first-search)
* [Breadth-first Search](#breadth-first-search)
* [Binary Search](#binary-search)
* [Math](#math)
* [Hash Table](#hash-table)
* [Sort](#sort-已全部做完)


| 数据结构 | 变种 | 相关题目 |
|:-------:|:-------|:------|
|顺序线性表：向量|||
|单链表|1.双链表<br>2.静态链表<br>3.对称矩阵<br>4.稀疏矩阵||
|栈|广义栈||
|队列|1.链表实现<br>2.循环数组实现<br>3.双端队列||
|字符串|1.KMP算法<br>2.有限状态自动机<br>3.模式匹配有限状态自动机<br>4.BM模式匹配算法<br>5.BM-KMP算法||
|树|1.二叉树<br>2.并查集<br>3.Huffman数||
|数组实现的堆|1.极大堆和极小堆<br>2.极大极小堆<br>3.双端堆<br>4.d叉堆||
|树实现的堆|1.左堆<br>2.扁堆<br>3.二项式堆<br>4.斐波那契堆<br>5.配对堆||
|查找|1.哈希表<br>2.跳跃表<br>3.排序二叉树<br>4.AVL树<br>5.B树<br>6.AA树<br>7.红黑树<br>8.排序二叉堆<br>9.Splay树<br>10.双链树<br>11.Trie树||



## 算法


## 一.分类

## Array

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[1. Two Sum](https://leetcode.com/problems/two-sum/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0001.%20Two%20Sum)| Easy | O(n)| O(n)||
|[11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0011.%20Container%20With%20Most%20Water)| Medium | O(n)| O(1)||
|[15. 3Sum](https://leetcode.com/problems/3sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0015.%203Sum)| Medium | O(n^2)| O(n)|❤️|
|[16. 3Sum Closest](https://leetcode.com/problems/3sum-closest)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0016.%203Sum%20Closest)| Medium | O(n^2)| O(1)||
|[18. 4Sum](https://leetcode.com/problems/4sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0018.%204Sum)| Medium | O(n^3)| O(n^2)|❤️|
|[26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0026.%20Remove%20Duplicates%20from%20Sorted%20Array)| Easy | O(n)| O(1)||
|[27. Remove Element](https://leetcode.com/problems/remove-element)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0027.%20Remove%20Element)| Easy | O(n)| O(1)||
|[39. Combination Sum](https://leetcode.com/problems/combination-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0039.%20Combination%20Sum)| Medium | O(n log n)| O(n)||
|[40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0040.%20Combination%20Sum%20II)| Medium | O(n log n)| O(n)||
|[41. First Missing Positive](https://leetcode.com/problems/first-missing-positive)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0041.%20First-Missing-Positive)| Hard | O(n)| O(n)||
|[42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0042.%20Trapping%20Rain%20Water)| Hard | O(n)| O(1)||
|[48. Rotate Image](https://leetcode.com/problems/rotate-image)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0048.%20Rotate%20Image)| Medium | O(n)| O(1)||
|[53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0053.%20Maximum%20Subarray)| Easy | O(n)| O(n)||
|[54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0054.%20Spiral%20Matrix)| Medium | O(n)| O(n^2)||
|[56. Merge Intervals](https://leetcode.com/problems/merge-intervals)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0056.%20Merge%20Intervals)| Medium | O(n log n)| O(1)||
|[57. Insert Interval](https://leetcode.com/problems/insert-interval)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0057.%20Insert%20Interval)| Hard | O(n)| O(1)||
|[59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0059.%20Spiral%20Matrix%20II)| Medium | O(n)| O(n^2)||
|[62. Unique Paths](https://leetcode.com/problems/unique-paths)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0062.%20Unique%20Paths)| Medium | O(n^2)| O(n^2)||
|[63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0063.%20Unique%20Paths%20II)| Medium | O(n^2)| O(n^2)||
|[64. Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0064.%20Minimum%20Path%20Sum)| Medium | O(n^2)| O(n^2)||
|[75. Sort Colors](https://leetcode.com/problems/sort-colors)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0075.%20Sort%20Colors)| Medium | O(n)| O(1)||
|[78. Subsets](https://leetcode.com/problems/subsets)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0078.%20Subsets)| Medium | O(n^2)| O(n)||
|[79. Word Search](https://leetcode.com/problems/word-search)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0079.%20Word%20Search)| Medium | O(n^2)| O(n^2)||
|[80. Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0080.%20Remove%20Duplicates%20from%20Sorted%20Array%20II)| Medium | O(n^2)| O(n^2)||
|[84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0084.%20Largest%20Rectangle%20in%20Histogram)| Medium | O(n)| O(n)||
|[88. Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0088.%20Merge-Sorted-Array)| Easy | O(n)| O(1)||
|[90. Subsets II](https://leetcode.com/problems/subsets-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0090.%20Subsets%20II)| Medium | O(n^2)| O(n)||
|[120. Triangle](https://leetcode.com/problems/triangle)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0120.%20Triangle)| Medium | O(n^2)| O(n)||
|[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock)| Easy | O(n)| O(1)||
|[122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0122.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20II)| Easy | O(n)| O(1)||
|[126. Word Ladder II](https://leetcode.com/problems/word-ladder-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0126.%20Word%20Ladder%20II)| Hard | O(n)| O(n^2)||
|[152. Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0152.%20Maximum%20Product%20Subarray)| Medium | O(n)| O(1)||
|[167. Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0167.%20Two%20Sum%20II%20-%20Input%20array%20is%20sorted)| Easy | O(n)| O(1)||
|[209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0209.%20Minimum%20Size%20Subarray%20Sum)| Medium | O(n)| O(1)||
|[216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0216.%20Combination%20Sum%20III)| Medium | O(n)| O(1)||
|[217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0217.%20Contains%20Duplicate)| Easy | O(n)| O(n)||
|[219. Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0219.%20Contains%20Duplicate%20II)| Easy | O(n)| O(n)||
|[283. Move Zeroes](https://leetcode.com/problems/move-zeroes)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0283.%20Move%20Zeroes)| Easy | O(n)| O(1)||
|[287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0287.%20Find%20the%20Duplicate%20Number)| Easy | O(n)| O(1)||
|[532. K-diff Pairs in an Array](https://leetcode.com/problems/k-diff-pairs-in-an-array)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0532.%20K-diff%20Pairs%20in%20an%20Array)| Easy | O(n)| O(n)||
|[566. Reshape the Matrix](https://leetcode.com/problems/reshape-the-matrix)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0566.%20Reshape%20the%20Matrix)| Easy | O(n^2)| O(n^2)||
|[628. Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0628.%20Maximum%20Product%20of%20Three%20Numbers)| Easy | O(n)| O(1)||
|[713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0713.%20Subarray%20Product%20Less%20Than%20K)| Medium | O(n)| O(1)||
|[714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0714.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20with%20Transaction%20Fee)| Medium | O(n)| O(1)||
|[746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0746.%20Min%20Cost%20Climbing%20Stairs)| Easy | O(n)| O(1)||
|[766. Toeplitz Matrix](https://leetcode.com/problems/toeplitz-matrix)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0766.%20Toeplitz%20Matrix)| Easy | O(n)| O(1)||
|[867. Transpose Matrix](https://leetcode.com/problems/transpose-matrix)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0867.%20Transpose%20Matrix)| Easy | O(n)| O(1)||
|[891. Sum of Subsequence Widths](https://leetcode.com/problems/sum-of-subsequence-widths)| [Go]()| Hard | O(n)| O(1)||
|[907. Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums)| [Go]()| Medium | O(n)| O(1)||
|[922. Sort Array By Parity II](https://leetcode.com/problems/sum-of-subarray-minimums)| [Go]()| Medium | O(n)| O(1)||
|[969. Pancake Sorting](https://leetcode.com/problems/pancake-sorting)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0969.%20Pancake%20Sorting)| Medium | O(n)| O(1)||
|[977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0977.%20Squares%20of%20a%20Sorted%20Array)| Easy | O(n)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|



## String

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters)| Medium | O(n)| O(1)||
|[17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0017.%20Letter%20Combinations%20of%20a%20Phone%20Number)| Medium | O(log n)| O(1)||
|[20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0020.%20Valid-Parentheses)| Easy | O(log n)| O(1)||
|[22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0022.%20Generate%20Parentheses)| Medium | O(log n)| O(1)||
|[28. Implement strStr()](https://leetcode.com/problems/implement-strstr)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0028.%20Implement%20strStr())| Easy | O(n)| O(1)||
|[30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0030.%20Substring%20with%20Concatenation%20of%20All%20Words)| Hard | O(n)| O(n)||
|[49. Group Anagrams](https://leetcode.com/problems/group-anagrams)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0049.%20Group%20Anagrams)| Medium | O(n log n)| O(n)||
|[71. Simplify Path](https://leetcode.com/problems/simplify-path)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0071.%20Simplify%20Path)| Medium | O(n)| O(n)||
|[76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0076.%20Minimum%20Window%20Substring)| Hard | O(n)| O(n)||
|[91. Decode Ways](https://leetcode.com/problems/decode-ways)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0091.%20Decode%20Ways)| Medium | O(n)| O(n)||
|[93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0093.%20Restore%20IP%20Addresses)| Medium | O(log n)| O(n)||
|[125. Valid Palindrome](https://leetcode.com/problems/valid-palindrome)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0125.%20Valid-Palindrome)| Easy | O(n)| O(1)||
|[126. Word Ladder II](https://leetcode.com/problems/word-ladder-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0126.%20Word%20Ladder%20II)| Hard | O(n)| O(n^2)||
|[344. Reverse String](https://leetcode.com/problems/reverse-string)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0344.%20Reverse%20String)| Easy | O(n)| O(1)||
|[345. Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0345.%20Reverse%20Vowels%20of%20a%20String)| Easy | O(n)| O(1)||
|[767. Reorganize String](https://leetcode.com/problems/reorganize-string/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0767.%20Reorganize%20String)| Medium | O(n log n)| O(log n)  |❤️|
|[842. Split Array into Fibonacci Sequence](https://leetcode.com/problems/split-array-into-fibonacci-sequence)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0842.%20Split%20Array%20into%20Fibonacci%20Sequence)| Medium | O(n^2)| O(1)||
|[856. Score of Parentheses](https://leetcode.com/problems/score-of-parentheses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0856.%20Score%20of%20Parentheses)| Medium | O(n)| O(n)||
|[925. Long Pressed Name](https://leetcode.com/problems/long-pressed-name)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0925.%20Long%20Pressed%20Name)| Easy | O(n)| O(1)||
|[1003. Check If Word Is Valid After Substitutions](https://leetcode.com/problems/check-if-word-is-valid-after-substitutions)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1003.%20Check%20If%20Word%20Is%20Valid%20After%20Substitutions)| Medium | O(n)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|




## Linked List (已全部做完)

- 巧妙的构造虚拟头结点。可以使遍历处理逻辑更加统一。
- 灵活使用递归。构造递归条件，使用递归可以巧妙的解题。不过需要注意有些题目不能使用递归，因为递归深度太深会导致超时和栈溢出。
- 链表区间逆序。第 92 题。
- 链表寻找中间节点。第 876 题。链表寻找倒数第 n 个节点。第 19 题。只需要一次遍历就可以得到答案。
- 合并 K 个有序链表。第 21 题，第 23 题。
- 链表归类。第 86 题，第 328 题。
- 链表排序，时间复杂度要求 O(n * log n)，空间复杂度 O(1)。只有一种做法，归并排序，至顶向下归并。第 148 题。
- 判断链表是否存在环，如果有环，输出环的交叉点的下标；判断 2 个链表是否有交叉点，如果有交叉点，输出交叉点。第 141 题，第 142 题，第 160 题。



| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0002.%20Add-Two-Number)| Medium | O(n)| O(1)||
|[19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0019.%20Remove%20Nth%20Node%20From%20End%20of%20List)| Medium | O(n)| O(1)||
|[21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0021.%20Merge%20Two%20Sorted%20Lists)| Easy | O(log n)| O(1)||
|[23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0023.%20Merge%20k%20Sorted%20Lists)| Hard | O(log n)| O(1)|❤️|
|[24. Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0024.%20Swap%20Nodes%20in%20Pairs)| Medium | O(n)| O(1)||
|[25. Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0025.%20Reverse%20Nodes%20in%20k%20Group)| Hard | O(log n)| O(1)|❤️|
|[61. Rotate List](https://leetcode.com/problems/rotate-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0061.%20Rotate%20List)| Medium | O(n)| O(1)||
|[82. Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0082.%20Remove%20Duplicates%20from%20Sorted%20List%20II)| Medium | O(n)| O(1)||
|[83. Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0083.%20Remove%20Duplicates%20from%20Sorted%20List)| Easy | O(n)| O(1)||
|[86. Partition List](https://leetcode.com/problems/partition-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0086.%20Partition%20List)| Medium | O(n)| O(1)|❤️|
|[92. Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0092.%20Reverse%20Linked%20List%20II)| Medium | O(n)| O(1)|❤️|
|[109. Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0109.%20Convert%20Sorted%20List%20to%20Binary%20Search%20Tree)| Medium | O(log n)| O(n)||
|[141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0141.%20Linked%20List%20Cycle)| Easy | O(n)| O(1)|❤️|
|[142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0142.%20Linked%20List%20Cycle%20II)| Medium | O(n)| O(1)|❤️|
|[143. Reorder List](https://leetcode.com/problems/reorder-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0143.%20Reorder%20List)| Medium | O(n)| O(1)|❤️|
|[147. Insertion Sort List](https://leetcode.com/problems/insertion-sort-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0147.%20Insertion%20Sort%20List)| Medium | O(n)| O(1)||
|[148. Sort List](https://leetcode.com/problems/sort-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0148.%20Sort%20List)| Medium | O(log n)| O(n)|❤️|
|[160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0160.%20Intersection%20of%20Two%20Linked%20Lists)| Easy | O(n)| O(1)|❤️|
|[203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0203.%20Remove%20Linked%20List%20Elements)| Easy | O(n)| O(1)||
|[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0206.%20Reverse-Linked-List)| Easy | O(n)| O(1)||
|[234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0234.%20Palindrome%20Linked%20List)| Easy | O(n)| O(1)||
|[237. Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0237.%20Delete%20Node%20in%20a%20Linked%20List)| Easy | O(n)| O(1)||
|[328. Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0328.%20Odd%20Even%20Linked%20List)| Medium | O(n)| O(1)||
|[445. Add Two Numbers II](https://leetcode.com/problems/add-two-numbers-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0445.%20Add%20Two%20Numbers%20II)| Medium | O(n)| O(n)||
|[725. Split Linked List in Parts](https://leetcode.com/problems/split-linked-list-in-parts/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0725.%20Split%20Linked%20List%20in%20Parts)| Medium | O(n)| O(1)||
|[817. Linked List Components](https://leetcode.com/problems/linked-list-components/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0817.%20Linked%20List%20Components)| Medium | O(n)| O(1)||
|[707. Design Linked List](https://leetcode.com/problems/design-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0707.%20Design%20Linked%20List)| Easy | O(n)| O(1)||
|[876. Middle of the Linked List](https://leetcode.com/problems/middle-of-the-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0876.%20Middle%20of%20the%20Linked%20List)| Easy | O(n)| O(1)|❤️|
|[1019. Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1019.%20Next%20Greater%20Node%20In%20Linked%20List)| Medium | O(n)| O(1)||
|----------------------------------------------------------------------------|-------------|-------------| -------------| -------------|-------------|



## Stack (已全部做完)

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0020.%20Valid-Parentheses)| Easy | O(log n)| O(1)||
|[42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0042.%20Trapping%20Rain%20Water)| Hard | O(n)| O(1)||
|[71. Simplify Path](https://leetcode.com/problems/simplify-path)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0071.%20Simplify%20Path)| Medium | O(n)| O(n)||
|[84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0084.%20Largest%20Rectangle%20in%20Histogram)| Medium | O(n)| O(n)||
|[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0094.%20Binary%20Tree%20Inorder%20Traversal)| Medium | O(n)| O(1)||
|[103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0103.%20Binary%20Tree%20Zigzag%20Level%20Order%20Traversal)| Medium | O(n)| O(n)||
|[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0144.%20Binary%20Tree%20Preorder%20Traversal)| Medium | O(n)| O(1)||
|[145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0145.%20Binary%20Tree%20Postorder%20Traversal)| Hard | O(n)| O(1)||
|[150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0150.%20Evaluate%20Reverse%20Polish%20Notation)| Medium | O(n)| O(1)||
|[155. Min Stack](https://leetcode.com/problems/min-stack)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0155.%20Min%20Stack)| Easy | O(n)| O(n)||
|[173. Binary Search Tree Iterator](https://leetcode.com/problems/binary-search-tree-iterator)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0173.%20Binary%20Search%20Tree%20Iterator)| Medium | O(n)| O(1)||
|[224. Basic Calculator](https://leetcode.com/problems/basic-calculator)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0224.%20Basic%20Calculator)| Hard | O(n)| O(n)||
|[225. Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0225.%20Implement%20Stack%20using%20Queues)| Easy | O(n)| O(n)||
|[232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0232.%20Implement%20Queue%20using%20Stacks)| Easy | O(n)| O(n)||
|[331. Verify Preorder Serialization of a Binary Tree](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0331.%20Verify%20Preorder%20Serialization%20of%20a%20Binary%20Tree)| Medium | O(n)| O(1)||
|[394. Decode String](https://leetcode.com/problems/decode-string)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0394.%20Decode%20String)| Medium | O(n)| O(n)||
|[402. Remove K Digits](https://leetcode.com/problems/remove-k-digits)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0402.%20Remove%20K%20Digits)| Medium | O(n)| O(1)||
|[456. 132 Pattern](https://leetcode.com/problems/132-pattern)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0456.%20132%20Pattern)| Medium | O(n)| O(n)||
|[496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0496.%20Next%20Greater%20Element%20I)| Easy | O(n)| O(n)||
|[503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0503.%20Next%20Greater%20Element%20II)| Medium | O(n)| O(n)||
|[636. Exclusive Time of Functions](https://leetcode.com/problems/exclusive-time-of-functions)| [Go](https://leetcode.com/problems/exclusive-time-of-functions)| Medium | O(n)| O(n)||
|[682. Baseball Game](https://leetcode.com/problems/baseball-game)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0682.%20Baseball%20Game)| Easy | O(n)| O(n)||
|[726. Number of Atoms](https://leetcode.com/problems/number-of-atoms)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0726.%20Number%20of%20Atoms)| Hard | O(n)| O(n) ||
|[735. Asteroid Collision](https://leetcode.com/problems/asteroid-collision)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0735.%20Asteroid%20Collision)| Medium | O(n)| O(n) ||
|[739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0739.%20Daily%20Temperatures)| Medium | O(n)| O(n) ||
|[844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0844.%20Backspace%20String%20Compare)| Easy | O(n)| O(n) ||
|[856. Score of Parentheses](https://leetcode.com/problems/score-of-parentheses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0856.%20Score%20of%20Parentheses)| Medium | O(n)| O(n)||
|[880. Decoded String at Index](https://leetcode.com/problems/decoded-string-at-index)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0880.%20Decoded%20String%20at%20Index)| Medium | O(n)| O(n)||
|[895. Maximum Frequency Stack](https://leetcode.com/problems/maximum-frequency-stack)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0895.%20Maximum%20Frequency%20Stack)| Hard | O(n)| O(n)  ||
|[901. Online Stock Span](https://leetcode.com/problems/online-stock-span)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0901.%20Online%20Stock%20Span)| Medium | O(n)| O(n)  ||
|[907. Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums)| [Go]()| Medium | O(n)| O(1)||
|[921. Minimum Add to Make Parentheses Valid](https://leetcode.com/problems/minimum-add-to-make-parentheses-valid)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0921.%20Minimum%20Add%20to%20Make%20Parentheses%20Valid)| Medium | O(n)| O(n)||
|[946. Validate Stack Sequences](https://leetcode.com/problems/validate-stack-sequences)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0946.%20Validate%20Stack%20Sequences)| Medium | O(n)| O(n)||
|[1003. Check If Word Is Valid After Substitutions](https://leetcode.com/problems/check-if-word-is-valid-after-substitutions)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1003.%20Check%20If%20Word%20Is%20Valid%20After%20Substitutions)| Medium | O(n)| O(1)||
|[1019. Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1019.%20Next%20Greater%20Node%20In%20Linked%20List)| Medium | O(n)| O(1)||
|[1021. Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1021.%20Remove%20Outermost%20Parentheses)| Medium | O(n)| O(1)||
|[1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1047.%20Remove%20All%20Adjacent%20Duplicates%20In%20String)| Medium | O(n)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|



## Tree

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0094.%20Binary%20Tree%20Inorder%20Traversal)| Medium | O(n)| O(1)||
|[96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0096.%20Unique%20Binary%20Search%20Trees)| Medium | O(n^2)| O(n)||
|[98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0098.%20Validate%20Binary%20Search%20Tree)| Medium | O(n)| O(1)||
|[99. Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0099.%20Recover%20Binary%20Search%20Tree)| Hard | O(n)| O(1)||
|[100. Same Tree](https://leetcode.com/problems/same-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0100.%20Same%20Tree)| Easy | O(n)| O(1)||
|[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0101.%20Symmetric%20Tree)| Easy | O(n)| O(1)||
|[102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0102.%20Binary%20Tree%20Level%20Order%20Traversal)| Medium | O(n)| O(1)||
|[103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0103.%20Binary%20Tree%20Zigzag%20Level%20Order%20Traversal)| Medium | O(n)| O(n)||
|[104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0104.%20Maximum%20Depth%20of%20Binary%20Tree)| Easy | O(n)| O(1)||
|[107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0107.%20Binary%20Tree%20Level%20Order%20Traversal%20II)| Easy | O(n)| O(1)||
|[108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0108.%20Convert%20Sorted%20Array%20to%20Binary%20Search%20Tree)| Easy | O(n)| O(1)||
|[110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0110.%20Balanced%20Binary%20Tree)| Easy | O(n)| O(1)||
|[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0111.%20Minimum%20Depth%20of%20Binary%20Tree)| Easy | O(n)| O(1)||
|[112. Path Sum](https://leetcode.com/problems/path-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0112.%20Path%20Sum)| Easy | O(n)| O(1)||
|[113. Path Sum II](https://leetcode.com/problems/path-sum-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0113.%20Path%20Sum%20II)| Medium | O(n)| O(1)||
|[114. Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0114.%20Flatten%20Binary%20Tree%20to%20Linked%20List)| Medium | O(n)| O(1)||
|[124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0124.%20Binary%20Tree%20Maximum%20Path%20Sum)| Hard | O(n)| O(1)||
|[129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0129.%20Sum%20Root%20to%20Leaf%20Numbers)| Medium | O(n)| O(1)||
|[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0144.%20Binary%20Tree%20Preorder%20Traversal)| Medium | O(n)| O(1)||
|[145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0145.%20Binary%20Tree%20Postorder%20Traversal)| Hard | O(n)| O(1)||
|[173. Binary Search Tree Iterator](https://leetcode.com/problems/binary-search-tree-iterator)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0173.%20Binary%20Search%20Tree%20Iterator)| Medium | O(n)| O(1)||
|[199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0199.%20Binary%20Tree%20Right%20Side%20Views)| Medium | O(n)| O(1)||
|[222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0222.%20Count%20Complete%20Tree%20Nodes)| Medium | O(n)| O(1)||
|[226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0226.%20Invert%20Binary%20Tree)| Easy | O(n)| O(1)||
|[230. Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0230.%20Kth%20Smallest%20Element%20in%20a%20BST)| Medium | O(n)| O(1)||
|[235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0235.%20Lowest%20Common%20Ancestor%20of%20a%20Binary%20Search%20Tree)| Easy | O(n)| O(1)||
|[236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0236.%20Lowest%20Common%20Ancestor%20of%20a%20Binary%20Tree)| Medium | O(n)| O(1)||
|[257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0257.%20Binary%20Tree%20Paths)| Easy | O(n)| O(1)||
|[404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0404.%20Sum%20of%20Left%20Leaves)| Easy | O(n)| O(1)||
|[437. Path Sum III](https://leetcode.com/problems/path-sum-iii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0437.%20Path%20Sum%20III)| Easy | O(n)| O(1)||
|[515. Find Largest Value in Each Tree Row](https://leetcode.com/problems/find-largest-value-in-each-tree-row)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0515.%20Find%20Largest%20Value%20in%20Each%20Tree%20Row)| Medium | O(n)| O(n)||
|[637. Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0637.%20Average%20of%20Levels%20in%20Binary%20Tree)| Easy | O(n)| O(n)||
|[993. Cousins in Binary Tree](https://leetcode.com/problems/cousins-in-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0993.%20Cousins%20in%20Binary%20Tree)| Easy | O(n)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|






## Dynamic Programming

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0053.%20Maximum%20Subarray)| Easy | O(n)| O(n)||
|[62. Unique Paths](https://leetcode.com/problems/unique-paths)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0062.%20Unique%20Paths)| Medium | O(n^2)| O(n^2)||
|[63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0063.%20Unique%20Paths%20II)| Medium | O(n^2)| O(n^2)||
|[64. Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0064.%20Minimum%20Path%20Sum)| Medium | O(n^2)| O(n^2)||
|[70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0070.%20Climbing%20Stairs)| Easy | O(n)| O(n)||
|[91. Decode Ways](https://leetcode.com/problems/decode-ways)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0091.%20Decode%20Ways)| Medium | O(n)| O(n)||
|[96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0096.%20Unique%20Binary%20Search%20Trees)| Medium | O(n)| O(n)||
|[120. Triangle](https://leetcode.com/problems/triangle)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0120.%20Triangle)| Medium | O(n^2)| O(n)||
|[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock)| Easy | O(n)| O(1)||
|[152. Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0152.%20Maximum%20Product%20Subarray)| Medium | O(n)| O(1)||
|[198. House Robber](https://leetcode.com/problems/house-robber)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0198.%20House%20Robber)| Easy | O(n)| O(n)||
|[213. House Robber II](https://leetcode.com/problems/house-robber-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0213.%20House%20Robber%20II)| Medium | O(n)| O(n)||
|[300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0300.%20Longest%20Increasing%20Subsequence)| Medium | O(n log n)| O(n)||
|[309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0309.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20with%20Cooldown)| Medium | O(n)| O(n)||
|[322. Coin Change](https://leetcode.com/problems/coin-change)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0322.%20Coin%20Change)| Medium | O(n)| O(n)||
|[338. Counting Bits](https://leetcode.com/problems/counting-bits)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0338.%20Counting%20Bits)| Medium | O(n)| O(n)||
|[343. Integer Break](https://leetcode.com/problems/integer-break)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0343.%20Integer%20Break)| Medium | O(n^2)| O(n)||
|[357. Count Numbers with Unique Digits](https://leetcode.com/problems/count-numbers-with-unique-digits)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0357.%20Count%20Numbers%20with%20Unique%20Digits)| Medium | O(1)| O(1)||
|[392. Is Subsequence](https://leetcode.com/problems/is-subsequence)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0392.%20Is%20Subsequence)| Medium | O(n)| O(1)||
|[416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0416.%20Partition%20Equal%20Subset%20Sum)| Medium | O(n^2)| O(n)||
|[714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0714.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20with%20Transaction%20Fee)| Medium | O(n)| O(1)||
|[746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0746.%20Min%20Cost%20Climbing%20Stairs)| Easy | O(n)| O(1)||
|[838. Push Dominoes](https://leetcode.com/problems/push-dominoes)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0838.%20Push%20Dominoes)| Medium | O(n)| O(1)||
|[1025. Divisor Game](https://leetcode.com/problems/divisor-game)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1025.%20Divisor%20Game)| Easy | O(1)| O(1)||
|[891. Sum of Subsequence Widths](https://leetcode.com/problems/sum-of-subsequence-widths)| [Go]()| Hard | O(n)| O(1)||
|[942. DI String Match](https://leetcode.com/problems/di-string-match)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0942.%20DI%20String%20Match)| Easy | O(n)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|





## Depth-first Search


| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0098.%20Validate%20Binary%20Search%20Tree)| Medium | O(n)| O(1)||
|[99. Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0099.%20Recover%20Binary%20Search%20Tree)| Hard | O(n)| O(1)||
|[100. Same Tree](https://leetcode.com/problems/same-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0100.%20Same%20Tree)| Easy | O(n)| O(1)||
|[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0101.%20Symmetric%20Tree)| Easy | O(n)| O(1)||
|[104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0104.%20Maximum%20Depth%20of%20Binary%20Tree)| Easy | O(n)| O(1)||
|[108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0108.%20Convert%20Sorted%20Array%20to%20Binary%20Search%20Tree)| Easy | O(n)| O(1)||
|[109. Convert Sorted List to Binary Search Tree](https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0109.%20Convert%20Sorted%20List%20to%20Binary%20Search%20Tree)| Medium | O(log n)| O(n)||
|[110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0110.%20Balanced%20Binary%20Tree)| Easy | O(n)| O(1)||
|[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0111.%20Minimum%20Depth%20of%20Binary%20Tree)| Easy | O(n)| O(1)||
|[112. Path Sum](https://leetcode.com/problems/path-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0112.%20Path%20Sum)| Easy | O(n)| O(1)||
|[113. Path Sum II](https://leetcode.com/problems/path-sum-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0113.%20Path%20Sum%20II)| Medium | O(n)| O(1)||
|[114. Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0114.%20Flatten%20Binary%20Tree%20to%20Linked%20List)| Medium | O(n)| O(1)||
|[124. Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0124.%20Binary%20Tree%20Maximum%20Path%20Sum)| Hard | O(n)| O(1)||
|[129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0129.%20Sum%20Root%20to%20Leaf%20Numbers)| Medium | O(n)| O(1)||
|[199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0199.%20Binary%20Tree%20Right%20Side%20Views)| Medium | O(n)| O(1)||
|[200. Number of Islands](https://leetcode.com/problems/number-of-islands)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0200.%20Number%20of%20Islands)| Medium | O(n^2)| O(n^2)||
|[207. Course Schedule](https://leetcode.com/problems/course-schedule)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0207.%20Course%20Schedule)| Medium | O(n^2)| O(n^2)||
|[210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0210.%20Course%20Schedule%20II)| Medium | O(n^2)| O(n^2)||
|[257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0257.%20Binary%20Tree%20Paths)| Easy | O(n)| O(1)||
|[394. Decode String](https://leetcode.com/problems/decode-string)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0394.%20Decode%20String)| Medium | O(n)| O(n)||
|[515. Find Largest Value in Each Tree Row](https://leetcode.com/problems/find-largest-value-in-each-tree-row)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0515.%20Find%20Largest%20Value%20in%20Each%20Tree%20Row)| Medium | O(n)| O(n)||
|[542. 01 Matrix](https://leetcode.com/problems/01-matrix)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0542.%2001%20Matrix)| Medium | O(n)| O(1)||
|[980. Unique Paths III](https://leetcode.com/problems/unique-paths-iii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0980.%20Unique%20Paths%20III)| Medium | O(n log n)| O(n)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|






## Breadth-first Search


| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0101.%20Symmetric%20Tree)| Easy | O(n)| O(1)||
|[102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0102.%20Binary%20Tree%20Level%20Order%20Traversal)| Medium | O(n)| O(1)||
|[103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0103.%20Binary%20Tree%20Zigzag%20Level%20Order%20Traversal)| Medium | O(n)| O(n)||
|[107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0107.%20Binary%20Tree%20Level%20Order%20Traversal%20II)| Easy | O(n)| O(1)||
|[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0111.%20Minimum%20Depth%20of%20Binary%20Tree)| Easy | O(n)| O(1)||
|[126. Word Ladder II](https://leetcode.com/problems/word-ladder-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0126.%20Word%20Ladder%20II)| Hard | O(n)| O(n^2)||
|[127. Word Ladder](https://leetcode.com/problems/word-ladder)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0127.%20Word%20Ladder)| Medium | O(n)| O(n)||
|[199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0199.%20Binary%20Tree%20Right%20Side%20Views)| Medium | O(n)| O(1)||
|[200. Number of Islands](https://leetcode.com/problems/number-of-islands)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0200.%20Number%20of%20Islands)| Medium | O(n^2)| O(n^2)||
|[207. Course Schedule](https://leetcode.com/problems/course-schedule)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0207.%20Course%20Schedule)| Medium | O(n^2)| O(n^2)||
|[210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0210.%20Course%20Schedule%20II)| Medium | O(n^2)| O(n^2)||
|[515. Find Largest Value in Each Tree Row](https://leetcode.com/problems/find-largest-value-in-each-tree-row)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0515.%20Find%20Largest%20Value%20in%20Each%20Tree%20Row)| Medium | O(n)| O(n)||
|[542. 01 Matrix](https://leetcode.com/problems/01-matrix)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0542.%2001%20Matrix)| Medium | O(n)| O(1)||
|[993. Cousins in Binary Tree](https://leetcode.com/problems/cousins-in-binary-tree)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0993.%20Cousins%20in%20Binary%20Tree)| Easy | O(n)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|






## Binary Search

| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[50. Pow(x, n)](https://leetcode.com/problems/powx-n)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0050.%20Pow(x%2C%20n))| Medium | O(log n)| O(1)||
|[69. Sqrt(x)](https://leetcode.com/problems/sqrtx)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0069.%20Sqrt(x))| Easy | O(log n)| O(1)||
|[167. Two Sum II - Input array is sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0167.%20Two%20Sum%20II%20-%20Input%20array%20is%20sorted)| Easy | O(n)| O(1)||
|[209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0209.%20Minimum%20Size%20Subarray%20Sum)| Medium | O(n)| O(1)||
|[222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0222.%20Count%20Complete%20Tree%20Nodes)| Medium | O(n)| O(1)||
|[230. Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0230.%20Kth%20Smallest%20Element%20in%20a%20BST)| Medium | O(n)| O(1)||
|[287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0287.%20Find%20the%20Duplicate%20Number)| Easy | O(n)| O(1)||
|[300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0300.%20Longest%20Increasing%20Subsequence)| Medium | O(n log n)| O(n)||
|[349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0349.%20Intersection%20of%20Two%20Arrays)| Easy | O(n)| O(n) ||
|[350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0350.%20Intersection%20of%20Two%20Arrays%20II)| Easy | O(n)| O(n) ||
|[392. Is Subsequence](https://leetcode.com/problems/is-subsequence)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0392.%20Is%20Subsequence)| Medium | O(n)| O(1)||
|[454. 4Sum II](https://leetcode.com/problems/4sum-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0454.%204Sum%20II)| Medium | O(n^2)| O(n) ||
|[710. Random Pick with Blacklist](https://leetcode.com/problems/random-pick-with-blacklist/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0710.%20Random%20Pick%20with%20Blacklist)| Hard | O(n)| O(n)  ||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|













## Math


| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0002.%20Add-Two-Number)| Medium | O(n)| O(1)||
|[50. Pow(x, n)](https://leetcode.com/problems/powx-n)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0050.%20Pow(x%2C%20n))| Medium | O(log n)| O(1)||
|[60. Permutation Sequence](https://leetcode.com/problems/permutation-sequence)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0060.%20Permutation%20Sequence)| Medium | O(n log n)| O(1)||
|[69. Sqrt(x)](https://leetcode.com/problems/sqrtx)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0069.%20Sqrt(x))| Easy | O(log n)| O(1)||
|[202. Happy Number](https://leetcode.com/problems/happy-number)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0202.%20Happy%20Number)| Easy | O(log n)| O(1)||
|[224. Basic Calculator](https://leetcode.com/problems/basic-calculator)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0224.%20Basic%20Calculator)| Hard | O(n)| O(n)||
|[231. Power of Two](https://leetcode.com/problems/power-of-twor)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0231.%20Power%20of%20Two)| Easy | O(1)| O(1)||
|[263. Ugly Number](https://leetcode.com/problems/ugly-number)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0263.%20Ugly%20Number)| Easy | O(log n)| O(1)||
|[326. Power of Three](https://leetcode.com/problems/power-of-three)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0326.%20Power%20of%20Three)| Easy | O(1)| O(1)||
|[343. Integer Break](https://leetcode.com/problems/integer-break)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0343.%20Integer%20Break)| Medium | O(n^2)| O(n)||
|[357. Count Numbers with Unique Digits](https://leetcode.com/problems/count-numbers-with-unique-digits)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0357.%20Count%20Numbers%20with%20Unique%20Digits)| Medium | O(1)| O(1)||
|[628. Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0628.%20Maximum%20Product%20of%20Three%20Numbers)| Easy | O(n)| O(1)||
|[885. Spiral Matrix III](https://leetcode.com/problems/spiral-matrix-iii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0885.%20Spiral%20Matrix%20III)| Medium | O(n^2)| O(1)||
|[891. Sum of Subsequence Widths](https://leetcode.com/problems/sum-of-subsequence-widths)| [Go]()| Hard | O(n)| O(1)||
|[942. DI String Match](https://leetcode.com/problems/di-string-match)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0942.%20DI%20String%20Match)| Easy | O(n)| O(1)||
|[976. Largest Perimeter Triangle](https://leetcode.com/problems/largest-perimeter-triangle/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0976.%20Largest%20Perimeter%20Triangle)| Easy | O(n log n)| O(log n) ||
|[996. Number of Squareful Arrays](https://leetcode.com/problems/number-of-squareful-arrays)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0996.%20Number%20of%20Squareful%20Arrays)| Hard | O(n log n)| O(n) ||
|[1025. Divisor Game](https://leetcode.com/problems/divisor-game)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1025.%20Divisor%20Game)| Easy | O(1)| O(1)||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|



## Hash Table


| Title | Solution | Difficulty | Time | Space |收藏| 
| ----- | :--------: | :----------: | :----: | :-----: | :-----: |
|[1. Two Sum](https://leetcode.com/problems/two-sum/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0001.%20Two%20Sum)| Easy | O(n)| O(n)||
|[3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0003.%20Longest%20Substring%20Without%20Repeating%20Characters)| Medium | O(n)| O(1)||
|[18. 4Sum](https://leetcode.com/problems/4sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0018.%204Sum)| Medium | O(n^3)| O(n^2)|❤️|
|[30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0030.%20Substring%20with%20Concatenation%20of%20All%20Words)| Hard | O(n)| O(n)||
|[36. Valid Sudoku](https://leetcode.com/problems/valid-sudoku)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0036.%20Valid%20Sudoku)| Medium | O(n^2)| O(n^2)||
|[37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0037.%20Sudoku%20Solver)| Hard | O(n^2)| O(n^2)||
|[49. Group Anagrams](https://leetcode.com/problems/group-anagrams)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0049.%20Group%20Anagrams)| Medium | O(n log n)| O(n)||
|[76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0076.%20Minimum%20Window%20Substring)| Hard | O(n)| O(n)||
|[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0094.%20Binary%20Tree%20Inorder%20Traversal)| Medium | O(n)| O(1)||
|[138. Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer)| [Go]()| Medium | O(n)| O(1)||
|[202. Happy Number](https://leetcode.com/problems/happy-number)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0202.%20Happy%20Number)| Easy | O(log n)| O(1)||
|[205. Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0205.%20Isomorphic%20Strings)| Easy | O(log n)| O(n)||
|[217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0217.%20Contains%20Duplicate)| Easy | O(n)| O(n)||
|[219. Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0219.%20Contains%20Duplicate%20II)| Easy | O(n)| O(n)||
|[242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0242.%20Valid%20Anagram)| Easy | O(n)| O(n) ||
|[274. H-Index](https://leetcode.com/problems/h-index/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0274.%20H-Index)| Medium |  O(n)| O(n) ||
|[290. Word Pattern](https://leetcode.com/problems/word-pattern)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0290.%20Word%20Pattern)| Easy |  O(n)| O(n) ||
|[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0347.%20Top%20K%20Frequent%20Elements)| Medium |  O(n)| O(n) ||
|[349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0349.%20Intersection%20of%20Two%20Arrays)| Easy | O(n)| O(n) ||
|[350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0350.%20Intersection%20of%20Two%20Arrays%20II)| Easy | O(n)| O(n) ||
|[438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0438.%20Find%20All%20Anagrams%20in%20a%20String)| Easy | O(n)| O(1) ||
|[447. Number of Boomerangs](https://leetcode.com/problems/number-of-boomerangs)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0447.%20Number%20of%20Boomerangs)| Easy | O(n)| O(1) ||
|[451. Sort Characters By Frequency](https://leetcode.com/problems/sort-characters-by-frequency)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0451.%20Sort%20Characters%20By%20Frequency)| Medium | O(n log n)| O(1) ||
|[454. 4Sum II](https://leetcode.com/problems/4sum-ii)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0454.%204Sum%20II)| Medium | O(n^2)| O(n) ||
|[648. Replace Words](https://leetcode.com/problems/replace-words)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0648.%20Replace%20Words)| Medium | O(n)| O(n) ||
|[676. Implement Magic Dictionary](https://leetcode.com/problems/implement-magic-dictionary)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0676.%20Implement%20Magic%20Dictionary)| Medium | O(n)| O(n) ||
|[720. Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0720.%20Longest%20Word%20in%20Dictionary)| Easy | O(n)| O(n) ||
|[726. Number of Atoms](https://leetcode.com/problems/number-of-atoms)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0726.%20Number%20of%20Atoms)| Hard | O(n)| O(n) ||
|[739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0739.%20Daily%20Temperatures)| Medium | O(n)| O(n) ||
|[710. Random Pick with Blacklist](https://leetcode.com/problems/random-pick-with-blacklist/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0710.%20Random%20Pick%20with%20Blacklist)| Hard | O(n)| O(n)  ||
|[895. Maximum Frequency Stack](https://leetcode.com/problems/maximum-frequency-stack)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0895.%20Maximum%20Frequency%20Stack)| Hard | O(n)| O(n)  ||
|[930. Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0930.%20Binary%20Subarrays%20With%20Sum)| Medium | O(n)| O(n)  ||
|[992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0992.%20Subarrays%20with%20K%20Different%20Integers)| Hard | O(n)| O(n)  ||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|



## Sort (已全部做完)

- 深刻的理解多路快排。第 75 题。
- 链表的排序，插入排序(第 147 题)和归并排序(第 148 题)
- 桶排序和基数排序。第 164 题。
- "摆动排序"。第 324 题。
- 两两不相邻的排序。第 767 题，第 1054 题。
- "饼子排序"。第 969 题。

| Title | Solution | Difficulty | Time | Space | 收藏 |
| ----- | :--------: | :----------: | :----: | :-----: |:-----: |
|[56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0056.%20Merge%20Intervals)| Medium | O(n log n)| O(log n)||
|[57. Insert Interval](https://leetcode.com/problems/insert-interval/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0057.%20Insert%20Interval)| Hard | O(n)| O(1)||
|[75. Sort Colors](https://leetcode.com/problems/sort-colors/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0075.%20Sort%20Colors)| Medium| O(n)| O(1)|❤️|
|[147. Insertion Sort List](https://leetcode.com/problems/insertion-sort-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0147.%20Insertion%20Sort%20List)| Medium | O(n)| O(1) |❤️|
|[148. Sort List](https://leetcode.com/problems/sort-list/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0148.%20Sort%20List)| Medium |O(n log n)| O(log n)|❤️|
|[164. Maximum Gap](https://leetcode.com/problems/maximum-gap/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0164.%20Maximum%20Gap)| Hard | O(n log n)| O(log n) |❤️|
|[179. Largest Number](https://leetcode.com/problems/largest-number/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0179.%20Largest%20Number)| Medium | O(n log n)| O(log n) |❤️|
|[220. Contains Duplicate III](https://leetcode.com/problems/contains-duplicate-iii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0220.%20Contains%20Duplicate%20III)| Medium | O(n^2)| O(1) ||
|[242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0242.%20Valid%20Anagram)| Easy | O(n)| O(n) ||
|[274. H-Index](https://leetcode.com/problems/h-index/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0274.%20H-Index)| Medium |  O(n)| O(n) ||
|[324. Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0324.%20Wiggle%20Sort%20II)| Medium| O(n)| O(n)|❤️|
|[349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0349.%20Intersection%20of%20Two%20Arrays)| Easy | O(n)| O(n) ||
|[350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0350.%20Intersection%20of%20Two%20Arrays%20II)| Easy | O(n)| O(n) ||
|[524. Longest Word in Dictionary through Deleting](https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0524.%20Longest%20Word%20in%20Dictionary%20through%20Deleting)| Medium | O(n)| O(1) ||
|[767. Reorganize String](https://leetcode.com/problems/reorganize-string/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0767.%20Reorganize%20String)| Medium | O(n log n)| O(log n)  |❤️|
|[853. Car Fleet](https://leetcode.com/problems/car-fleet/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0853.%20Car%20Fleet)| Medium | O(n log n)| O(log n)  ||
|[710. Random Pick with Blacklist](https://leetcode.com/problems/random-pick-with-blacklist/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0710.%20Random%20Pick%20with%20Blacklist)| Hard | O(n)| O(n)  ||
|[922. Sort Array By Parity II](https://leetcode.com/problems/sort-array-by-parity-ii/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0922.%20Sort%20Array%20By%20Parity%20II)| Easy | O(n)| O(1) ||
|[969. Pancake Sorting](https://leetcode.com/problems/pancake-sorting/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0969.%20Pancake%20Sorting)| Medium | O(n log n)| O(log n) |❤️|
|[973. K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0973.%20K%20Closest%20Points%20to%20Origin)| Medium | O(n log n)| O(log n) ||
|[976. Largest Perimeter Triangle](https://leetcode.com/problems/largest-perimeter-triangle/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0976.%20Largest%20Perimeter%20Triangle)| Easy | O(n log n)| O(log n) ||
|[1030. Matrix Cells in Distance Order](https://leetcode.com/problems/matrix-cells-in-distance-order/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1030.%20Matrix%20Cells%20in%20Distance%20Order)| Easy | O(n^2)| O(1) ||
|[1054. Distant Barcodes](https://leetcode.com/problems/distant-barcodes/)| [Go](https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/1054.%20Distant%20Barcodes)| Medium | O(n log n)| O(log n) ||
|-----------------------------------------------------------------|-------------|-------------| --------------------------| --------------------------|-------------|





----------------------------------------------------------------------------------------
