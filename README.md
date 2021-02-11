
# LeetCode in Go
[LeetCode Online Judge](https://leetcode.com/) is a website containing many **algorithm questions**. Most of them are real interview questions of **Google, Facebook, LinkedIn, Apple**, etc. and it always help to sharp our algorithm Skills. Level up your coding skills and quickly land a job. This is the best place to expand your knowledge and get prepared for your next interview. This repo shows my solutions in Go with the code style strictly follows the [Google Golang Style Guide](https://github.com/golang/go/wiki/CodeReviewComments). Please feel free to reference and **STAR** to support this repo, thank you!


<p align='center'>
<img src='./logo.png'>
</p>

![](./website/static/wechat-qr-code.png)

<p align='center'>
<a href="https://github.com/halfrost/LeetCode-Go/releases/" rel="nofollow"><img alt="GitHub All Releases" src="https://img.shields.io/github/downloads/halfrost/LeetCode-Go/total?label=PDF%20downloads"></a>
<img src="https://img.shields.io/badge/Total%20Word%20Count-738884-success">
<a href="https://github.com/halfrost/LeetCode-Go/actions" rel="nofollow"><img src="https://github.com/halfrost/LeetCode-Go/workflows/Deploy%20leetcode-cookbook/badge.svg?branch=master"></a>
<a href="https://travis-ci.org/github/halfrost/LeetCode-Go" rel="nofollow"><img src="https://travis-ci.org/halfrost/LeetCode-Go.svg?branch=master"></a>
<a href="https://goreportcard.com/report/github.com/halfrost/LeetCode-Go" rel="nofollow"><img src="https://goreportcard.com/badge/github.com/halfrost/LeetCode-Go"></a>
<img src="https://img.shields.io/badge/runtime%20beats-100%25-success">
<!--<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/halfrost/LeetCode-Go?color=26C2F0">-->
<img alt="Support Go version" src="https://img.shields.io/badge/Go-v1.15-26C2F0">
<img src="https://visitor-badge.laobi.icu/badge?page_id=halfrost.LeetCode-Go">
</p>

<p align='center'>
<a href="https://github.com/halfrost/LeetCode-Go/blob/master/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/halfrost/LeetCode-Go?label=License"></a>
<img src="https://img.shields.io/badge/License-CC-000000.svg">
<a href="https://leetcode.com/halfrost/"><img src="https://img.shields.io/badge/@halfrost-8751-yellow.svg">
<img src="https://img.shields.io/badge/language-Golang-26C2F0.svg">
<a href="https://halfrost.com"><img src="https://img.shields.io/badge/Blog-Halfrost--Field-80d4f9.svg?style=flat"></a>
<a href="http://weibo.com/halfrost"><img src="https://img.shields.io/badge/weibo-@halfrost-f974ce.svg?style=flat&colorA=f4292e"></a>
<a href="https://twitter.com/halffrost"><img src="https://img.shields.io/badge/twitter-@halffrost-F8E81C.svg?style=flat&colorA=009df2"></a>
<a href="https://www.zhihu.com/people/halfrost/activities"><img src="https://img.shields.io/badge/%E7%9F%A5%E4%B9%8E-@halfrost-fd6f32.svg?style=flat&colorA=0083ea"></a>
<img src="https://img.shields.io/badge/made%20with-=1-blue.svg">
<a href="https://github.com/halfrost/LeetCode-Go/pulls"><img src="https://img.shields.io/badge/PR-Welcome-brightgreen.svg"></a>
</p>

支持 Progressive Web Apps 和 Dark Mode 的题解电子书《LeetCode Cookbook》 <a href="https://books.halfrost.com/leetcode/" rel="nofollow">Online Reading</a>

<p align='center'>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_Safari_0.png"></a>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_Chrome_PWA.png"></a>
</p>

离线版本的电子书《LeetCode Cookbook》PDF <a href="https://github.com/halfrost/LeetCode-Go/releases/" rel="nofollow">Download here</a>

<p align='center'>
<a href="https://github.com/halfrost/LeetCode-Go/releases/"><img src="https://img.halfrost.com/Leetcode/Cookbook.png"></a>
</p>

通过 iOS / Android 浏览器安装 PWA 版《LeetCode Cookbook》至设备桌面随时学习

<p align='center'>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_PWA_iPad.png"></a>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_PWA_iPad_example1__.png"></a>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_PWA_iPad_example2__.png"></a>
</p>


## Data Structures

> 标识了 ✅ 的专题是完成所有题目了的，没有标识的是还没有做完所有题目的

<a href="https://books.halfrost.com/leetcode/"><img src="./website/static/logo.png" alt="logo" height="550" align="right" /></a>

* [Array](#array)
* [String](#string)
* [✅ Two Pointers](#two-pointers)
* [✅ Linked List](#linked-list)
* [✅ Stack](#stack)
* [Tree](#tree)
* [Dynamic programming](#dynamic-programming)
* [✅ Backtracking](#backtracking)
* [Depth First Search](#depth-first-search)
* [Breadth First Search](#breadth-first-search)
* [Binary Search](#binary-search)
* [Math](#math)
* [Hash Table](#hash-table)
* [✅ Sort](#sort)
* [✅ Bit Manipulation](#bit-manipulation)
* [✅ Union Find](#union-find)
* [✅ Sliding Window](#sliding-window)
* [✅ Segment Tree](#segment-tree)
* [✅ Binary Indexed Tree](#binary-indexed-tree)

<br>
<br>

| 数据结构 | 变种 | 相关题目 | 讲解文章 | 
|:-------:|:-------|:------|:------|
|顺序线性表：向量||||
|单链表|1. 双向链表<br>2. 静态链表<br>3. 对称矩阵<br>4. 稀疏矩阵|||
|哈希表|1. 散列函数<br>2. 解决碰撞/填充因子<br>|||
|栈和队列|1. 广义栈<br>2. 双端队列<br>|||
|队列|1. 链表实现<br>2. 循环数组实现<br>3. 双端队列|||
|字符串|1. KMP算法<br>2. 有限状态自动机<br>3. 模式匹配有限状态自动机<br>4. BM 模式匹配算法<br>5. BM-KMP 算法<br>6. BF 算法|||
|树|1. 二叉树<br>2. 并查集<br>3. Huffman 树|||
|数组实现的堆|1. 极大堆和极小堆<br>2. 极大极小堆<br>3. 双端堆<br>4. d 叉堆|||
|树实现的堆|1. 左堆<br>2. 扁堆<br>3. 二项式堆<br>4. 斐波那契堆<br>5. 配对堆|||
|查找|1. 哈希表<br>2. 跳跃表<br>3. 排序二叉树<br>4. AVL 树<br>5. B 树 / B+ 树 / B* 树<br>6. AA 树<br>7. 红黑树<br>8. 排序二叉堆<br>9. Splay 树<br>10. 双链树<br>11. Trie 树<br>12. R 树|||
|--------------------------------------------|--------------------------------------------------------------------------------------------|---------------------------|-----------------------------------|


## Algorithm


| 算法 | 具体类型 | 相关题目 | 讲解文章 | 
|:-------:|:-------|:------|:------|
|排序算法|1. 冒泡排序<br>2. 插入排序<br>3. 选择排序<br>4. 希尔 Shell 排序<br>5. 快速排序<br>6. 归并排序<br>7. 堆排序<br>8. 线性排序算法<br>9. 自省排序<br>10. 间接排序<br>11. 计数排序<br>12. 基数排序<br>13. 桶排序<br>14. 外部排序 - k 路归并败者树<br>15. 外部排序 - 最佳归并树|||
|递归与分治||1. 二分搜索/查找<br>2. 大整数的乘法<br>3. Strassen 矩阵乘法<br>4. 棋盘覆盖<br>5. 合并排序<br>6. 快速排序<br>7. 线性时间选择<br>8. 最接近点对问题<br>9. 循环赛日程表<br>||
|动态规划||1. 矩阵连乘问题<br>2. 最长公共子序列<br>3. 最大子段和<br>4. 凸多边形最优三角剖分<br>5. 多边形游戏<br>6. 图像压缩<br>7. 电路布线<br>8. 流水作业调度<br>9. 0-1 背包问题/背包九讲<br>10. 最优二叉搜索树<br>11. 动态规划加速原理<br>12. 树型 DP<br>||
|贪心||1. 活动安排问题<br>2. 最优装载<br>3. 哈夫曼编码<br>4. 单源最短路径<br>5. 最小生成树<br>6. 多机调度问题<br>||
|回溯法||1. 装载问题<br>2. 批处理作业调度<br>3. 符号三角形问题<br>4. n 后问题<br>5. 0-1 背包问题<br>6. 最大团问题<br>7. 图的 m 着色问题<br>8. 旅行售货员问题<br>9. 圆排列问题<br>10. 电路板排列问题<br>11. 连续邮资问题<br>||
|搜索|1. 枚举<br>2. DFS<br>3. BFS<br>4. 启发式搜索<br>|||
|随机化|1. 随机数<br>2. 数值随机化算法<br>3. Sherwood 舍伍德算法<br>4. Las Vegas 拉斯维加斯算法<br>5. Monte Carlo 蒙特卡罗算法<br>|1. 计算 π 值<br>2. 计算定积分<br>3. 解非线性方程组<br>4. 线性时间选择算法<br>5. 跳跃表<br>6. n 后问题<br>7. 整数因子分解<br>8. 主元素问题<br>9. 素数测试<br>||
|图论|1. 遍历 DFS / BFS<br>2. AOV / AOE 网络<br>3. Kruskal 算法(最小生成树)<br>4. Prim 算法(最小生成树)<br>5. Boruvka 算法(最小生成树)<br>6. Dijkstra 算法(单源最短路径)<br>7. Bellman-Ford 算法(单源最短路径)<br>8. SPFA 算法(单源最短路径)<br>9. Floyd 算法(多源最短路径)<br>10. Johnson 算法(多源最短路径)<br>11. Fleury 算法(欧拉回路)<br>12. Ford-Fulkerson 算法(最大网络流增广路)<br>13. Edmonds-Karp 算法(最大网络流)<br>14. Dinic 算法(最大网络流)<br>15. 一般预流推进算法<br>16. 最高标号预流推进 HLPP 算法<br>17. Primal-Dual 原始对偶算法(最小费用流)18. Kosaraju 算法(有向图强连通分量)<br>19. Tarjan 算法(有向图强连通分量)<br>20. Gabow 算法(有向图强连通分量)<br>21. 匈牙利算法(二分图匹配)<br>22. Hopcroft－Karp 算法(二分图匹配)<br>23. kuhn munkras 算法(二分图最佳匹配)<br>24. Edmonds’ Blossom-Contraction 算法(一般图匹配)<br>|1. 图遍历<br>2. 有向图和无向图的强弱连通性<br>3. 割点/割边<br>3. AOV 网络和拓扑排序<br>4. AOE 网络和关键路径<br>5. 最小代价生成树/次小生成树<br>6. 最短路径问题/第 K 短路问题<br>7. 最大网络流问题<br>8. 最小费用流问题<br>9. 图着色问题<br>10. 差分约束系统<br>11. 欧拉回路<br>12. 中国邮递员问题<br>13. 汉密尔顿回路<br>14. 最佳边割集/最佳点割集/最小边割集/最小点割集/最小路径覆盖/最小点集覆盖 <br>15. 边覆盖集<br>16. 二分图完美匹配和最大匹配问题<br>17. 仙人掌图<br>18. 弦图<br>19. 稳定婚姻问题<br>20. 最大团问题<br>||
|数论||1. 最大公约数<br> 2. 最小公倍数<br>3. 分解质因数<br>4. 素数判定<br>5. 进制转换<br>6. 高精度计算<br>7. 整除问题<br>8. 同余问题<br>9. 欧拉函数<br>10. 扩展欧几里得<br>11. 置换群<br>12. 母函数<br>13. 离散变换<br>14. 康托展开<br>15. 矩阵<br>16. 向量<br>17. 线性方程组<br>18. 线性规划<br> ||
|几何||1. 凸包 - Gift wrapping<br>2. 凸包 - Graham scan<br>3. 线段问题<br> 4. 多边形和多面体相关问题<br>||
|NP 完全|1. 计算模型<br>2. P 类与 NP 类问题<br>3. NP 完全问题<br>4. NP 完全问题的近似算法<br>|1. 随机存取机 RAM<br>2. 随机存取存储程序机 RASP<br>3. 图灵机<br>4. 非确定性图灵机<br>5. P 类与 NP 类语言<br>6. 多项式时间验证<br>7. 多项式时间变换<br>8. Cook定理<br>9. 合取范式的可满足性问题 CNF-SAT<br>10. 3 元合取范式的可满足性问题 3-SAT<br>11. 团问题 CLIQUE<br>12. 顶点覆盖问题 VERTEX-COVER<br>13. 子集和问题 SUBSET-SUM<br>14. 哈密顿回路问题 HAM-CYCLE<br>15. 旅行售货员问题 TSP<br>16. 顶点覆盖问题的近似算法<br>17. 旅行售货员问题近似算法<br>18. 具有三角不等式性质的旅行售货员问题<br>19. 一般的旅行售货员问题<br>20. 集合覆盖问题的近似算法<br>21. 子集和问题的近似算法<br>22. 子集和问题的指数时间算法<br>23. 子集和问题的多项式时间近似格式<br>||
|------------|------------------------------------------------------------------|-----------------------------------------------------------------|--------------------|


## LeetCode Problems

## 一. 个人数据

|    |  Easy  |  Medium  |  Hard |  Total |
|:--------:|:--------:|:--------:|:--------:|:--------:|
|Optimizing|38|48|19|105|
|Accepted|**253**|**328**|**99**|**680**|
|Total|461|922|373|1756|
|Perfection Rate|85.0%|85.4%|80.8%|84.6%|
|Completion Rate|54.9%|35.6%|26.5%|38.7%|
|------------|----------------------------|----------------------------|----------------------------|----------------------------|

## 二. 目录

以下已经收录了 575 道题的题解，还有 14 道题在尝试优化到 beats 100%

| No.    |  Title  |  Solution  |  Acceptance |  Difficulty |  Frequency |
|:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|
|0001|Two Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0001.Two-Sum)|46.4%|Easy||
|0002|Add Two Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0002.Add-Two-Numbers)|35.4%|Medium||
|0003|Longest Substring Without Repeating Characters|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0003.Longest-Substring-Without-Repeating-Characters)|31.4%|Medium||
|0004|Median of Two Sorted Arrays|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0004.Median-of-Two-Sorted-Arrays)|31.0%|Hard||
|0005|Longest Palindromic Substring||30.4%|Medium||
|0006|ZigZag Conversion||37.8%|Medium||
|0007|Reverse Integer|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0007.Reverse-Integer)|25.9%|Easy||
|0008|String to Integer (atoi)||15.6%|Medium||
|0009|Palindrome Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0009.Palindrome-Number)|49.7%|Easy||
|0010|Regular Expression Matching||27.3%|Hard||
|0011|Container With Most Water|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0011.Container-With-Most-Water)|52.4%|Medium||
|0012|Integer to Roman||56.1%|Medium||
|0013|Roman to Integer|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0013.Roman-to-Integer)|56.5%|Easy||
|0014|Longest Common Prefix||36.1%|Easy||
|0015|3Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0015.3Sum)|27.9%|Medium||
|0016|3Sum Closest|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0016.3Sum-Closest)|46.3%|Medium||
|0017|Letter Combinations of a Phone Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0017.Letter-Combinations-of-a-Phone-Number)|49.0%|Medium||
|0018|4Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0018.4Sum)|34.8%|Medium||
|0019|Remove Nth Node From End of List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0019.Remove-Nth-Node-From-End-of-List)|35.7%|Medium||
|0020|Valid Parentheses|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0020.Valid-Parentheses)|39.8%|Easy||
|0021|Merge Two Sorted Lists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0021.Merge-Two-Sorted-Lists)|55.9%|Easy||
|0022|Generate Parentheses|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0022.Generate-Parentheses)|65.1%|Medium||
|0023|Merge k Sorted Lists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0023.Merge-k-Sorted-Lists)|42.5%|Hard||
|0024|Swap Nodes in Pairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0024.Swap-Nodes-in-Pairs)|52.9%|Medium||
|0025|Reverse Nodes in k-Group|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0025.Reverse-Nodes-in-k-Group)|44.7%|Hard||
|0026|Remove Duplicates from Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0026.Remove-Duplicates-from-Sorted-Array)|46.5%|Easy||
|0027|Remove Element|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0027.Remove-Element)|49.2%|Easy||
|0028|Implement strStr()|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0028.Implement-strStr())|35.1%|Easy||
|0029|Divide Two Integers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0029.Divide-Two-Integers)|16.6%|Medium||
|0030|Substring with Concatenation of All Words|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0030.Substring-with-Concatenation-of-All-Words)|26.1%|Hard||
|0031|Next Permutation||33.8%|Medium||
|0032|Longest Valid Parentheses||29.2%|Hard||
|0033|Search in Rotated Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0033.Search-in-Rotated-Sorted-Array)|35.8%|Medium||
|0034|Find First and Last Position of Element in Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0034.Find-First-and-Last-Position-of-Element-in-Sorted-Array)|37.2%|Medium||
|0035|Search Insert Position|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0035.Search-Insert-Position)|42.8%|Easy||
|0036|Valid Sudoku|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0036.Valid-Sudoku)|50.4%|Medium||
|0037|Sudoku Solver|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0037.Sudoku-Solver)|46.3%|Hard||
|0038|Count and Say||46.0%|Easy||
|0039|Combination Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0039.Combination-Sum)|59.1%|Medium||
|0040|Combination Sum II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0040.Combination-Sum-II)|50.0%|Medium||
|0041|First Missing Positive|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0041.First-Missing-Positive)|33.6%|Hard||
|0042|Trapping Rain Water|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0042.Trapping-Rain-Water)|51.0%|Hard||
|0043|Multiply Strings||34.9%|Medium||
|0044|Wildcard Matching||25.4%|Hard||
|0045|Jump Game II||31.4%|Hard||
|0046|Permutations|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0046.Permutations)|66.4%|Medium||
|0047|Permutations II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0047.Permutations-II)|49.3%|Medium||
|0048|Rotate Image|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0048.Rotate-Image)|59.7%|Medium||
|0049|Group Anagrams|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0049.Group-Anagrams)|59.1%|Medium||
|0050|Pow(x, n)|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0050.Pow(x,-n))|30.9%|Medium||
|0051|N-Queens|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0051.N-Queens)|49.4%|Hard||
|0052|N-Queens II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0052.N-Queens-II)|60.0%|Hard||
|0053|Maximum Subarray|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0053.Maximum-Subarray)|47.7%|Easy||
|0054|Spiral Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0054.Spiral-Matrix)|35.8%|Medium||
|0055|Jump Game|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0055.Jump-Game)|35.1%|Medium||
|0056|Merge Intervals|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0056.Merge-Intervals)|40.9%|Medium||
|0057|Insert Interval|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0057.Insert-Interval)|35.0%|Medium||
|0058|Length of Last Word||33.5%|Easy||
|0059|Spiral Matrix II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0059.Spiral-Matrix-II)|57.7%|Medium||
|0060|Permutation Sequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0060.Permutation-Sequence)|39.3%|Hard||
|0061|Rotate List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0061.Rotate-List)|31.7%|Medium||
|0062|Unique Paths|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0062.Unique-Paths)|55.9%|Medium||
|0063|Unique Paths II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0063.Unique-Paths-II)|35.2%|Medium||
|0064|Minimum Path Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0064.Minimum-Path-Sum)|56.1%|Medium||
|0065|Valid Number||15.8%|Hard||
|0066|Plus One|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0066.Plus-One)|42.4%|Easy||
|0067|Add Binary|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0067.Add-Binary)|46.8%|Easy||
|0068|Text Justification||29.5%|Hard||
|0069|Sqrt(x)|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0069.Sqrt(x))|35.0%|Easy||
|0070|Climbing Stairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0070.Climbing-Stairs)|48.6%|Easy||
|0071|Simplify Path|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0071.Simplify-Path)|34.6%|Medium||
|0072|Edit Distance||46.6%|Hard||
|0073|Set Matrix Zeroes||44.2%|Medium||
|0074|Search a 2D Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0074.Search-a-2D-Matrix)|37.6%|Medium||
|0075|Sort Colors|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0075.Sort-Colors)|49.2%|Medium||
|0076|Minimum Window Substring|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0076.Minimum-Window-Substring)|35.9%|Hard||
|0077|Combinations|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0077.Combinations)|57.3%|Medium||
|0078|Subsets|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0078.Subsets)|64.9%|Medium||
|0079|Word Search|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0079.Word-Search)|36.7%|Medium||
|0080|Remove Duplicates from Sorted Array II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0080.Remove-Duplicates-from-Sorted-Array-II)|46.0%|Medium||
|0081|Search in Rotated Sorted Array II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0081.Search-in-Rotated-Sorted-Array-II)|33.5%|Medium||
|0082|Remove Duplicates from Sorted List II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0082.Remove-Duplicates-from-Sorted-List-II)|39.2%|Medium||
|0083|Remove Duplicates from Sorted List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0083.Remove-Duplicates-from-Sorted-List)|46.4%|Easy||
|0084|Largest Rectangle in Histogram|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0084.Largest-Rectangle-in-Histogram)|37.0%|Hard||
|0085|Maximal Rectangle||39.4%|Hard||
|0086|Partition List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0086.Partition-List)|43.3%|Medium||
|0087|Scramble String||34.6%|Hard||
|0088|Merge Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0088.Merge-Sorted-Array)|40.7%|Easy||
|0089|Gray Code|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0089.Gray-Code)|50.3%|Medium||
|0090|Subsets II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0090.Subsets-II)|48.7%|Medium||
|0091|Decode Ways|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0091.Decode-Ways)|26.4%|Medium||
|0092|Reverse Linked List II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0092.Reverse-Linked-List-II)|40.5%|Medium||
|0093|Restore IP Addresses|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0093.Restore-IP-Addresses)|37.4%|Medium||
|0094|Binary Tree Inorder Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0094.Binary-Tree-Inorder-Traversal)|65.7%|Medium||
|0095|Unique Binary Search Trees II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0095.Unique-Binary-Search-Trees-II)|42.4%|Medium||
|0096|Unique Binary Search Trees|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0096.Unique-Binary-Search-Trees)|54.3%|Medium||
|0097|Interleaving String||32.5%|Hard||
|0098|Validate Binary Search Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0098.Validate-Binary-Search-Tree)|28.7%|Medium||
|0099|Recover Binary Search Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0099.Recover-Binary-Search-Tree)|42.4%|Hard||
|0100|Same Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0100.Same-Tree)|54.1%|Easy||
|0101|Symmetric Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0101.Symmetric-Tree)|48.0%|Easy||
|0102|Binary Tree Level Order Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0102.Binary-Tree-Level-Order-Traversal)|56.5%|Medium||
|0103|Binary Tree Zigzag Level Order Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0103.Binary-Tree-Zigzag-Level-Order-Traversal)|50.0%|Medium||
|0104|Maximum Depth of Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0104.Maximum-Depth-of-Binary-Tree)|67.9%|Easy||
|0105|Construct Binary Tree from Preorder and Inorder Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0105.Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal)|51.6%|Medium||
|0106|Construct Binary Tree from Inorder and Postorder Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0106.Construct-Binary-Tree-from-Inorder-and-Postorder-Traversal)|49.5%|Medium||
|0107|Binary Tree Level Order Traversal II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0107.Binary-Tree-Level-Order-Traversal-II)|55.1%|Easy||
|0108|Convert Sorted Array to Binary Search Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0108.Convert-Sorted-Array-to-Binary-Search-Tree)|60.2%|Easy||
|0109|Convert Sorted List to Binary Search Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0109.Convert-Sorted-List-to-Binary-Search-Tree)|50.1%|Medium||
|0110|Balanced Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0110.Balanced-Binary-Tree)|44.7%|Easy||
|0111|Minimum Depth of Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0111.Minimum-Depth-of-Binary-Tree)|39.4%|Easy||
|0112|Path Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0112.Path-Sum)|42.3%|Easy||
|0113|Path Sum II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0113.Path-Sum-II)|48.9%|Medium||
|0114|Flatten Binary Tree to Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0114.Flatten-Binary-Tree-to-Linked-List)|51.7%|Medium||
|0115|Distinct Subsequences||39.6%|Hard||
|0116|Populating Next Right Pointers in Each Node||48.9%|Medium||
|0117|Populating Next Right Pointers in Each Node II||41.9%|Medium||
|0118|Pascal's Triangle|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0118.Pascal's-Triangle)|54.7%|Easy||
|0119|Pascal's Triangle II||52.1%|Easy||
|0120|Triangle|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0120.Triangle)|45.7%|Medium||
|0121|Best Time to Buy and Sell Stock|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0121.Best-Time-to-Buy-and-Sell-Stock)|51.5%|Easy||
|0122|Best Time to Buy and Sell Stock II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0122.Best-Time-to-Buy-and-Sell-Stock-II)|58.4%|Easy||
|0123|Best Time to Buy and Sell Stock III||39.8%|Hard||
|0124|Binary Tree Maximum Path Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0124.Binary-Tree-Maximum-Path-Sum)|35.4%|Hard||
|0125|Valid Palindrome|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0125.Valid-Palindrome)|38.1%|Easy||
|0126|Word Ladder II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0126.Word-Ladder-II)|23.6%|Hard||
|0127|Word Ladder|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0127.Word-Ladder)|31.7%|Hard||
|0128|Longest Consecutive Sequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0128.Longest-Consecutive-Sequence)|46.2%|Hard||
|0129|Sum Root to Leaf Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0129.Sum-Root-to-Leaf-Numbers)|50.8%|Medium||
|0130|Surrounded Regions|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0130.Surrounded-Regions)|29.3%|Medium||
|0131|Palindrome Partitioning|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0131.Palindrome-Partitioning)|51.8%|Medium||
|0132|Palindrome Partitioning II||31.2%|Hard||
|0133|Clone Graph||39.0%|Medium||
|0134|Gas Station||41.1%|Medium||
|0135|Candy||32.9%|Hard||
|0136|Single Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0136.Single-Number)|66.5%|Easy||
|0137|Single Number II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0137.Single-Number-II)|53.7%|Medium||
|0138|Copy List with Random Pointer|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0138.Copy-List-with-Random-Pointer)|40.5%|Medium||
|0139|Word Break||41.6%|Medium||
|0140|Word Break II||34.5%|Hard||
|0141|Linked List Cycle|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0141.Linked-List-Cycle)|42.8%|Easy||
|0142|Linked List Cycle II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0142.Linked-List-Cycle-II)|39.6%|Medium||
|0143|Reorder List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0143.Reorder-List)|40.5%|Medium||
|0144|Binary Tree Preorder Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0144.Binary-Tree-Preorder-Traversal)|57.4%|Medium||
|0145|Binary Tree Postorder Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0145.Binary-Tree-Postorder-Traversal)|57.4%|Medium||
|0146|LRU Cache|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0146.LRU-Cache)|35.6%|Medium||
|0147|Insertion Sort List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0147.Insertion-Sort-List)|44.3%|Medium||
|0148|Sort List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0148.Sort-List)|46.1%|Medium||
|0149|Max Points on a Line||17.3%|Hard||
|0150|Evaluate Reverse Polish Notation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0150.Evaluate-Reverse-Polish-Notation)|37.8%|Medium||
|0151|Reverse Words in a String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0151.Reverse-Words-in-a-String)|23.6%|Medium||
|0152|Maximum Product Subarray|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0152.Maximum-Product-Subarray)|32.7%|Medium||
|0153|Find Minimum in Rotated Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0153.Find-Minimum-in-Rotated-Sorted-Array)|46.1%|Medium||
|0154|Find Minimum in Rotated Sorted Array II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0154.Find-Minimum-in-Rotated-Sorted-Array-II)|42.0%|Hard||
|0155|Min Stack|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0155.Min-Stack)|46.3%|Easy||
|0156|Binary Tree Upside Down||56.3%|Medium||
|0157|Read N Characters Given Read4||37.1%|Easy||
|0158|Read N Characters Given Read4 II - Call multiple times||36.7%|Hard||
|0159|Longest Substring with At Most Two Distinct Characters||50.4%|Medium||
|0160|Intersection of Two Linked Lists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0160.Intersection-of-Two-Linked-Lists)|43.0%|Easy||
|0161|One Edit Distance||33.0%|Medium||
|0162|Find Peak Element|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0162.Find-Peak-Element)|43.9%|Medium||
|0163|Missing Ranges||26.7%|Easy||
|0164|Maximum Gap|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0164.Maximum-Gap)|36.7%|Hard||
|0165|Compare Version Numbers||30.3%|Medium||
|0166|Fraction to Recurring Decimal||22.2%|Medium||
|0167|Two Sum II - Input array is sorted|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0167.Two-Sum-II---Input-array-is-sorted)|55.6%|Easy||
|0168|Excel Sheet Column Title|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0168.Excel-Sheet-Column-Title)|31.7%|Easy||
|0169|Majority Element|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0169.Majority-Element)|60.0%|Easy||
|0170|Two Sum III - Data structure design||34.8%|Easy||
|0171|Excel Sheet Column Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0171.Excel-Sheet-Column-Number)|56.8%|Easy||
|0172|Factorial Trailing Zeroes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0172.Factorial-Trailing-Zeroes)|38.5%|Easy||
|0173|Binary Search Tree Iterator|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0173.Binary-Search-Tree-Iterator)|59.9%|Medium||
|0174|Dungeon Game|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0174.Dungeon-Game)|33.2%|Hard||
|0175|Combine Two Tables||63.9%|Easy||
|0176|Second Highest Salary||33.1%|Easy||
|0177|Nth Highest Salary||33.2%|Medium||
|0178|Rank Scores||49.8%|Medium||
|0179|Largest Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0179.Largest-Number)|30.6%|Medium||
|0180|Consecutive Numbers||42.0%|Medium||
|0181|Employees Earning More Than Their Managers||60.1%|Easy||
|0182|Duplicate Emails||64.4%|Easy||
|0183|Customers Who Never Order||56.6%|Easy||
|0184|Department Highest Salary||39.8%|Medium||
|0185|Department Top Three Salaries||38.6%|Hard||
|0186|Reverse Words in a String II||45.4%|Medium||
|0187|Repeated DNA Sequences|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0187.Repeated-DNA-Sequences)|41.4%|Medium||
|0188|Best Time to Buy and Sell Stock IV||29.5%|Hard||
|0189|Rotate Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0189.Rotate-Array)|36.5%|Medium||
|0190|Reverse Bits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0190.Reverse-Bits)|41.9%|Easy||
|0191|Number of 1 Bits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0191.Number-of-1-Bits)|53.6%|Easy||
|0192|Word Frequency||25.7%|Medium||
|0193|Valid Phone Numbers||25.3%|Easy||
|0194|Transpose File||24.4%|Medium||
|0195|Tenth Line||32.9%|Easy||
|0196|Delete Duplicate Emails||44.9%|Easy||
|0197|Rising Temperature||39.8%|Easy||
|0198|House Robber|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0198.House-Robber)|42.8%|Medium||
|0199|Binary Tree Right Side View|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0199.Binary-Tree-Right-Side-View)|56.3%|Medium||
|0200|Number of Islands|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0200.Number-of-Islands)|48.9%|Medium||
|0201|Bitwise AND of Numbers Range|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0201.Bitwise-AND-of-Numbers-Range)|39.6%|Medium||
|0202|Happy Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0202.Happy-Number)|51.2%|Easy||
|0203|Remove Linked List Elements|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0203.Remove-Linked-List-Elements)|39.2%|Easy||
|0204|Count Primes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0204.Count-Primes)|32.2%|Easy||
|0205|Isomorphic Strings|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0205.Isomorphic-Strings)|40.4%|Easy||
|0206|Reverse Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0206.Reverse-Linked-List)|65.1%|Easy||
|0207|Course Schedule|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0207.Course-Schedule)|44.3%|Medium||
|0208|Implement Trie (Prefix Tree)|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0208.Implement-Trie-(Prefix-Tree))|51.9%|Medium||
|0209|Minimum Size Subarray Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0209.Minimum-Size-Subarray-Sum)|39.4%|Medium||
|0210|Course Schedule II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0210.Course-Schedule-II)|42.5%|Medium||
|0211|Design Add and Search Words Data Structure|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0211.Design-Add-and-Search-Words-Data-Structure)|40.0%|Medium||
|0212|Word Search II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0212.Word-Search-II)|36.9%|Hard||
|0213|House Robber II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0213.House-Robber-II)|37.5%|Medium||
|0214|Shortest Palindrome||30.6%|Hard||
|0215|Kth Largest Element in an Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0215.Kth-Largest-Element-in-an-Array)|58.3%|Medium||
|0216|Combination Sum III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0216.Combination-Sum-III)|60.2%|Medium||
|0217|Contains Duplicate|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0217.Contains-Duplicate)|56.6%|Easy||
|0218|The Skyline Problem|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0218.The-Skyline-Problem)|36.2%|Hard||
|0219|Contains Duplicate II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0219.Contains-Duplicate-II)|38.6%|Easy||
|0220|Contains Duplicate III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0220.Contains-Duplicate-III)|21.3%|Medium||
|0221|Maximal Square||39.0%|Medium||
|0222|Count Complete Tree Nodes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0222.Count-Complete-Tree-Nodes)|49.1%|Medium||
|0223|Rectangle Area|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0223.Rectangle-Area)|38.2%|Medium||
|0224|Basic Calculator|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0224.Basic-Calculator)|38.1%|Hard||
|0225|Implement Stack using Queues|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0225.Implement-Stack-using-Queues)|47.2%|Easy||
|0226|Invert Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0226.Invert-Binary-Tree)|66.8%|Easy||
|0227|Basic Calculator II||38.5%|Medium||
|0228|Summary Ranges|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0228.Summary-Ranges)|42.3%|Easy||
|0229|Majority Element II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0229.Majority-Element-II)|38.7%|Medium||
|0230|Kth Smallest Element in a BST|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0230.Kth-Smallest-Element-in-a-BST)|62.4%|Medium||
|0231|Power of Two|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0231.Power-of-Two)|43.8%|Easy||
|0232|Implement Queue using Stacks|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0232.Implement-Queue-using-Stacks)|52.0%|Easy||
|0233|Number of Digit One||31.7%|Hard||
|0234|Palindrome Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0234.Palindrome-Linked-List)|40.4%|Easy||
|0235|Lowest Common Ancestor of a Binary Search Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0235.Lowest-Common-Ancestor-of-a-Binary-Search-Tree)|51.6%|Easy||
|0236|Lowest Common Ancestor of a Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0236.Lowest-Common-Ancestor-of-a-Binary-Tree)|48.5%|Medium||
|0237|Delete Node in a Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0237.Delete-Node-in-a-Linked-List)|66.7%|Easy||
|0238|Product of Array Except Self||61.4%|Medium||
|0239|Sliding Window Maximum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0239.Sliding-Window-Maximum)|44.6%|Hard||
|0240|Search a 2D Matrix II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0240.Search-a-2D-Matrix-II)|44.4%|Medium||
|0241|Different Ways to Add Parentheses||57.2%|Medium||
|0242|Valid Anagram|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0242.Valid-Anagram)|58.1%|Easy||
|0243|Shortest Word Distance||61.9%|Easy||
|0244|Shortest Word Distance II||53.9%|Medium||
|0245|Shortest Word Distance III||55.9%|Medium||
|0246|Strobogrammatic Number||45.8%|Easy||
|0247|Strobogrammatic Number II||48.4%|Medium||
|0248|Strobogrammatic Number III||40.2%|Hard||
|0249|Group Shifted Strings||57.9%|Medium||
|0250|Count Univalue Subtrees||53.3%|Medium||
|0251|Flatten 2D Vector||46.3%|Medium||
|0252|Meeting Rooms||55.3%|Easy||
|0253|Meeting Rooms II||46.8%|Medium||
|0254|Factor Combinations||47.4%|Medium||
|0255|Verify Preorder Sequence in Binary Search Tree||46.1%|Medium||
|0256|Paint House||53.5%|Medium||
|0257|Binary Tree Paths|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0257.Binary-Tree-Paths)|53.4%|Easy||
|0258|Add Digits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0258.Add-Digits)|58.4%|Easy||
|0259|3Sum Smaller||48.9%|Medium||
|0260|Single Number III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0260.Single-Number-III)|65.3%|Medium||
|0261|Graph Valid Tree||43.1%|Medium||
|0262|Trips and Users||35.3%|Hard||
|0263|Ugly Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0263.Ugly-Number)|41.7%|Easy||
|0264|Ugly Number II||42.8%|Medium||
|0265|Paint House II||45.6%|Hard||
|0266|Palindrome Permutation||62.4%|Easy||
|0267|Palindrome Permutation II||37.4%|Medium||
|0268|Missing Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0268.Missing-Number)|53.7%|Easy||
|0269|Alien Dictionary||33.7%|Hard||
|0270|Closest Binary Search Tree Value||49.8%|Easy||
|0271|Encode and Decode Strings||32.6%|Medium||
|0272|Closest Binary Search Tree Value II||52.2%|Hard||
|0273|Integer to English Words||28.0%|Hard||
|0274|H-Index|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0274.H-Index)|36.3%|Medium||
|0275|H-Index II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0275.H-Index-II)|36.3%|Medium||
|0276|Paint Fence||39.0%|Easy||
|0277|Find the Celebrity||43.3%|Medium||
|0278|First Bad Version||37.3%|Easy||
|0279|Perfect Squares||48.8%|Medium||
|0280|Wiggle Sort||64.6%|Medium||
|0281|Zigzag Iterator||59.4%|Medium||
|0282|Expression Add Operators||36.7%|Hard||
|0283|Move Zeroes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0283.Move-Zeroes)|58.5%|Easy||
|0284|Peeking Iterator|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0284.Peeking-Iterator)|50.4%|Medium||
|0285|Inorder Successor in BST||42.7%|Medium||
|0286|Walls and Gates||56.3%|Medium||
|0287|Find the Duplicate Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0287.Find-the-Duplicate-Number)|57.4%|Medium||
|0288|Unique Word Abbreviation||23.0%|Medium||
|0289|Game of Life||58.0%|Medium||
|0290|Word Pattern|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0290.Word-Pattern)|38.3%|Easy||
|0291|Word Pattern II||44.1%|Medium||
|0292|Nim Game||55.0%|Easy||
|0293|Flip Game||61.3%|Easy||
|0294|Flip Game II||50.6%|Medium||
|0295|Find Median from Data Stream||46.7%|Hard||
|0296|Best Meeting Point||58.1%|Hard||
|0297|Serialize and Deserialize Binary Tree||49.6%|Hard||
|0298|Binary Tree Longest Consecutive Sequence||48.0%|Medium||
|0299|Bulls and Cows||44.4%|Medium||
|0300|Longest Increasing Subsequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0300.Longest-Increasing-Subsequence)|43.9%|Medium||
|0301|Remove Invalid Parentheses||44.6%|Hard||
|0302|Smallest Rectangle Enclosing Black Pixels||52.4%|Hard||
|0303|Range Sum Query - Immutable|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0303.Range-Sum-Query---Immutable)|47.4%|Easy||
|0304|Range Sum Query 2D - Immutable||40.4%|Medium||
|0305|Number of Islands II||39.7%|Hard||
|0306|Additive Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0306.Additive-Number)|29.6%|Medium||
|0307|Range Sum Query - Mutable|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0307.Range-Sum-Query---Mutable)|36.6%|Medium||
|0308|Range Sum Query 2D - Mutable||37.5%|Hard||
|0309|Best Time to Buy and Sell Stock with Cooldown|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0309.Best-Time-to-Buy-and-Sell-Stock-with-Cooldown)|48.1%|Medium||
|0310|Minimum Height Trees||34.6%|Medium||
|0311|Sparse Matrix Multiplication||63.8%|Medium||
|0312|Burst Balloons||53.7%|Hard||
|0313|Super Ugly Number||46.1%|Medium||
|0314|Binary Tree Vertical Order Traversal||46.9%|Medium||
|0315|Count of Smaller Numbers After Self|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0315.Count-of-Smaller-Numbers-After-Self)|42.5%|Hard||
|0316|Remove Duplicate Letters||39.0%|Medium||
|0317|Shortest Distance from All Buildings||42.7%|Hard||
|0318|Maximum Product of Word Lengths|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0318.Maximum-Product-of-Word-Lengths)|52.2%|Medium||
|0319|Bulb Switcher||45.3%|Medium||
|0320|Generalized Abbreviation||53.6%|Medium||
|0321|Create Maximum Number||27.5%|Hard||
|0322|Coin Change|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0322.Coin-Change)|37.1%|Medium||
|0323|Number of Connected Components in an Undirected Graph||57.6%|Medium||
|0324|Wiggle Sort II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0324.Wiggle-Sort-II)|30.6%|Medium||
|0325|Maximum Size Subarray Sum Equals k||47.3%|Medium||
|0326|Power of Three|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0326.Power-of-Three)|42.1%|Easy||
|0327|Count of Range Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0327.Count-of-Range-Sum)|36.0%|Hard||
|0328|Odd Even Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0328.Odd-Even-Linked-List)|57.0%|Medium||
|0329|Longest Increasing Path in a Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0329.Longest-Increasing-Path-in-a-Matrix)|44.7%|Hard||
|0330|Patching Array||34.9%|Hard||
|0331|Verify Preorder Serialization of a Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0331.Verify-Preorder-Serialization-of-a-Binary-Tree)|41.0%|Medium||
|0332|Reconstruct Itinerary||37.8%|Medium||
|0333|Largest BST Subtree||37.7%|Medium||
|0334|Increasing Triplet Subsequence||40.7%|Medium||
|0335|Self Crossing||28.6%|Hard||
|0336|Palindrome Pairs||34.5%|Hard||
|0337|House Robber III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0337.House-Robber-III)|51.8%|Medium||
|0338|Counting Bits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0338.Counting-Bits)|70.3%|Medium||
|0339|Nested List Weight Sum||76.3%|Medium||
|0340|Longest Substring with At Most K Distinct Characters||45.4%|Medium||
|0341|Flatten Nested List Iterator||54.4%|Medium||
|0342|Power of Four|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0342.Power-of-Four)|41.7%|Easy||
|0343|Integer Break|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0343.Integer-Break)|51.2%|Medium||
|0344|Reverse String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0344.Reverse-String)|70.2%|Easy||
|0345|Reverse Vowels of a String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0345.Reverse-Vowels-of-a-String)|45.0%|Easy||
|0346|Moving Average from Data Stream||73.1%|Easy||
|0347|Top K Frequent Elements|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0347.Top-K-Frequent-Elements)|62.3%|Medium||
|0348|Design Tic-Tac-Toe||55.5%|Medium||
|0349|Intersection of Two Arrays|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0349.Intersection-of-Two-Arrays)|64.7%|Easy||
|0350|Intersection of Two Arrays II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0350.Intersection-of-Two-Arrays-II)|52.0%|Easy||
|0351|Android Unlock Patterns||49.6%|Medium||
|0352|Data Stream as Disjoint Intervals||48.6%|Hard||
|0353|Design Snake Game||35.7%|Medium||
|0354|Russian Doll Envelopes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0354.Russian-Doll-Envelopes)|36.2%|Hard||
|0355|Design Twitter||31.3%|Medium||
|0356|Line Reflection||32.8%|Medium||
|0357|Count Numbers with Unique Digits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0357.Count-Numbers-with-Unique-Digits)|48.8%|Medium||
|0358|Rearrange String k Distance Apart||35.6%|Hard||
|0359|Logger Rate Limiter||72.1%|Easy||
|0360|Sort Transformed Array||49.6%|Medium||
|0361|Bomb Enemy||46.6%|Medium||
|0362|Design Hit Counter||65.1%|Medium||
|0363|Max Sum of Rectangle No Larger Than K||38.4%|Hard||
|0364|Nested List Weight Sum II||63.6%|Medium||
|0365|Water and Jug Problem||31.0%|Medium||
|0366|Find Leaves of Binary Tree||71.7%|Medium||
|0367|Valid Perfect Square|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0367.Valid-Perfect-Square)|42.0%|Easy||
|0368|Largest Divisible Subset||38.3%|Medium||
|0369|Plus One Linked List||59.5%|Medium||
|0370|Range Addition||63.5%|Medium||
|0371|Sum of Two Integers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0371.Sum-of-Two-Integers)|50.6%|Medium||
|0372|Super Pow|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0372.Super-Pow)|36.7%|Medium||
|0373|Find K Pairs with Smallest Sums|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0373.Find-K-Pairs-with-Smallest-Sums)|37.6%|Medium||
|0374|Guess Number Higher or Lower||44.6%|Easy||
|0375|Guess Number Higher or Lower II||42.1%|Medium||
|0376|Wiggle Subsequence||40.2%|Medium||
|0377|Combination Sum IV||46.0%|Medium||
|0378|Kth Smallest Element in a Sorted Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0378.Kth-Smallest-Element-in-a-Sorted-Matrix)|56.1%|Medium||
|0379|Design Phone Directory||47.9%|Medium||
|0380|Insert Delete GetRandom O(1)||48.8%|Medium||
|0381|Insert Delete GetRandom O(1) - Duplicates allowed||34.8%|Hard||
|0382|Linked List Random Node||54.0%|Medium||
|0383|Ransom Note||53.3%|Easy||
|0384|Shuffle an Array||53.9%|Medium||
|0385|Mini Parser|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0385.Mini-Parser)|34.5%|Medium||
|0386|Lexicographical Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0386.Lexicographical-Numbers)|54.0%|Medium||
|0387|First Unique Character in a String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0387.First-Unique-Character-in-a-String)|53.8%|Easy||
|0388|Longest Absolute File Path||42.7%|Medium||
|0389|Find the Difference|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0389.Find-the-Difference)|57.8%|Easy||
|0390|Elimination Game||44.9%|Medium||
|0391|Perfect Rectangle||31.1%|Hard||
|0392|Is Subsequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0392.Is-Subsequence)|49.5%|Easy||
|0393|UTF-8 Validation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0393.UTF-8-Validation)|38.0%|Medium||
|0394|Decode String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0394.Decode-String)|52.6%|Medium||
|0395|Longest Substring with At Least K Repeating Characters||43.5%|Medium||
|0396|Rotate Function||36.7%|Medium||
|0397|Integer Replacement|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0397.Integer-Replacement)|33.4%|Medium||
|0398|Random Pick Index||57.8%|Medium||
|0399|Evaluate Division|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0399.Evaluate-Division)|54.2%|Medium||
|0400|Nth Digit||32.4%|Medium||
|0401|Binary Watch|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0401.Binary-Watch)|48.4%|Easy||
|0402|Remove K Digits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0402.Remove-K-Digits)|28.6%|Medium||
|0403|Frog Jump||41.3%|Hard||
|0404|Sum of Left Leaves|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0404.Sum-of-Left-Leaves)|52.2%|Easy||
|0405|Convert a Number to Hexadecimal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0405.Convert-a-Number-to-Hexadecimal)|44.4%|Easy||
|0406|Queue Reconstruction by Height||68.3%|Medium||
|0407|Trapping Rain Water II||44.0%|Hard||
|0408|Valid Word Abbreviation||31.3%|Easy||
|0409|Longest Palindrome|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0409.Longest-Palindrome)|52.2%|Easy||
|0410|Split Array Largest Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0410.Split-Array-Largest-Sum)|46.2%|Hard||
|0411|Minimum Unique Word Abbreviation||37.2%|Hard||
|0412|Fizz Buzz|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0412.Fizz-Buzz)|63.5%|Easy||
|0413|Arithmetic Slices||58.6%|Medium||
|0414|Third Maximum Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0414.Third-Maximum-Number)|30.7%|Easy||
|0415|Add Strings||48.2%|Easy||
|0416|Partition Equal Subset Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0416.Partition-Equal-Subset-Sum)|44.8%|Medium||
|0417|Pacific Atlantic Water Flow||42.5%|Medium||
|0418|Sentence Screen Fitting||33.1%|Medium||
|0419|Battleships in a Board||71.0%|Medium||
|0420|Strong Password Checker||13.9%|Hard||
|0421|Maximum XOR of Two Numbers in an Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0421.Maximum-XOR-of-Two-Numbers-in-an-Array)|54.0%|Medium||
|0422|Valid Word Square||38.1%|Easy||
|0423|Reconstruct Original Digits from English||47.5%|Medium||
|0424|Longest Repeating Character Replacement|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0424.Longest-Repeating-Character-Replacement)|48.2%|Medium||
|0425|Word Squares||50.1%|Hard||
|0426|Convert Binary Search Tree to Sorted Doubly Linked List||61.1%|Medium||
|0427|Construct Quad Tree||62.5%|Medium||
|0428|Serialize and Deserialize N-ary Tree||61.4%|Hard||
|0429|N-ary Tree Level Order Traversal||66.6%|Medium||
|0430|Flatten a Multilevel Doubly Linked List||56.8%|Medium||
|0431|Encode N-ary Tree to Binary Tree||74.6%|Hard||
|0432|All O`one Data Structure||33.1%|Hard||
|0433|Minimum Genetic Mutation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0433.Minimum-Genetic-Mutation)|43.2%|Medium||
|0434|Number of Segments in a String||37.8%|Easy||
|0435|Non-overlapping Intervals|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0435.Non-overlapping-Intervals)|43.9%|Medium||
|0436|Find Right Interval|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0436.Find-Right-Interval)|48.5%|Medium||
|0437|Path Sum III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0437.Path-Sum-III)|48.1%|Medium||
|0438|Find All Anagrams in a String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0438.Find-All-Anagrams-in-a-String)|44.9%|Medium||
|0439|Ternary Expression Parser||56.7%|Medium||
|0440|K-th Smallest in Lexicographical Order||29.8%|Hard||
|0441|Arranging Coins|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0441.Arranging-Coins)|42.3%|Easy||
|0442|Find All Duplicates in an Array||68.9%|Medium||
|0443|String Compression||43.4%|Medium||
|0444|Sequence Reconstruction||23.5%|Medium||
|0445|Add Two Numbers II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0445.Add-Two-Numbers-II)|56.2%|Medium||
|0446|Arithmetic Slices II - Subsequence||33.4%|Hard||
|0447|Number of Boomerangs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0447.Number-of-Boomerangs)|52.4%|Medium||
|0448|Find All Numbers Disappeared in an Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0448.Find-All-Numbers-Disappeared-in-an-Array)|56.1%|Easy||
|0449|Serialize and Deserialize BST||54.0%|Medium||
|0450|Delete Node in a BST||45.2%|Medium||
|0451|Sort Characters By Frequency|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0451.Sort-Characters-By-Frequency)|64.3%|Medium||
|0452|Minimum Number of Arrows to Burst Balloons||49.8%|Medium||
|0453|Minimum Moves to Equal Array Elements|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0453.Minimum-Moves-to-Equal-Array-Elements)|50.8%|Easy||
|0454|4Sum II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0454.4Sum-II)|54.6%|Medium||
|0455|Assign Cookies|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0455.Assign-Cookies)|50.4%|Easy||
|0456|132 Pattern|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0456.132-Pattern)|30.6%|Medium||
|0457|Circular Array Loop|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0457.Circular-Array-Loop)|30.0%|Medium||
|0458|Poor Pigs||54.5%|Hard||
|0459|Repeated Substring Pattern||43.3%|Easy||
|0460|LFU Cache|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0460.LFU-Cache)|35.9%|Hard||
|0461|Hamming Distance|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0461.Hamming-Distance)|73.2%|Easy||
|0462|Minimum Moves to Equal Array Elements II||54.3%|Medium||
|0463|Island Perimeter|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0463.Island-Perimeter)|66.6%|Easy||
|0464|Can I Win||29.6%|Medium||
|0465|Optimal Account Balancing||48.2%|Hard||
|0466|Count The Repetitions||28.6%|Hard||
|0467|Unique Substrings in Wraparound String||36.0%|Medium||
|0468|Validate IP Address||24.9%|Medium||
|0469|Convex Polygon||37.4%|Medium||
|0470|Implement Rand10() Using Rand7()|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0470.Implement-Rand10()-Using-Rand7())|46.0%|Medium||
|0471|Encode String with Shortest Length||49.2%|Hard||
|0472|Concatenated Words||44.2%|Hard||
|0473|Matchsticks to Square||38.2%|Medium||
|0474|Ones and Zeroes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0474.Ones-and-Zeroes)|43.6%|Medium||
|0475|Heaters|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0475.Heaters)|33.6%|Medium||
|0476|Number Complement|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0476.Number-Complement)|65.1%|Easy||
|0477|Total Hamming Distance|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0477.Total-Hamming-Distance)|50.6%|Medium||
|0478|Generate Random Point in a Circle||38.9%|Medium||
|0479|Largest Palindrome Product||29.5%|Hard||
|0480|Sliding Window Median|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0480.Sliding-Window-Median)|38.6%|Hard||
|0481|Magical String||48.0%|Medium||
|0482|License Key Formatting||43.1%|Easy||
|0483|Smallest Good Base|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0483.Smallest-Good-Base)|36.2%|Hard||
|0484|Find Permutation||64.2%|Medium||
|0485|Max Consecutive Ones|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0485.Max-Consecutive-Ones)|53.0%|Easy||
|0486|Predict the Winner||48.5%|Medium||
|0487|Max Consecutive Ones II||47.9%|Medium||
|0488|Zuma Game||38.4%|Hard||
|0489|Robot Room Cleaner||72.3%|Hard||
|0490|The Maze||52.7%|Medium||
|0491|Increasing Subsequences|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0491.Increasing-Subsequences)|47.5%|Medium||
|0492|Construct the Rectangle||50.3%|Easy||
|0493|Reverse Pairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0493.Reverse-Pairs)|26.8%|Hard||
|0494|Target Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0494.Target-Sum)|45.8%|Medium||
|0495|Teemo Attacking||56.1%|Medium||
|0496|Next Greater Element I|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0496.Next-Greater-Element-I)|65.4%|Easy||
|0497|Random Point in Non-overlapping Rectangles|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0497.Random-Point-in-Non-overlapping-Rectangles)|39.1%|Medium||
|0498|Diagonal Traverse|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0498.Diagonal-Traverse)|50.2%|Medium||
|0499|The Maze III||42.4%|Hard||
|0500|Keyboard Row|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0500.Keyboard-Row)|65.6%|Easy||
|0501|Find Mode in Binary Search Tree||43.4%|Easy||
|0502|IPO||41.4%|Hard||
|0503|Next Greater Element II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0503.Next-Greater-Element-II)|58.4%|Medium||
|0504|Base 7||46.5%|Easy||
|0505|The Maze II||48.5%|Medium||
|0506|Relative Ranks||51.2%|Easy||
|0507|Perfect Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0507.Perfect-Number)|36.1%|Easy||
|0508|Most Frequent Subtree Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0508.Most-Frequent-Subtree-Sum)|59.0%|Medium||
|0509|Fibonacci Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0509.Fibonacci-Number)|67.4%|Easy||
|0510|Inorder Successor in BST II||59.9%|Medium||
|0511|Game Play Analysis I||81.3%|Easy||
|0512|Game Play Analysis II||55.9%|Easy||
|0513|Find Bottom Left Tree Value|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0513.Find-Bottom-Left-Tree-Value)|62.5%|Medium||
|0514|Freedom Trail||44.9%|Hard||
|0515|Find Largest Value in Each Tree Row|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0515.Find-Largest-Value-in-Each-Tree-Row)|62.1%|Medium||
|0516|Longest Palindromic Subsequence||55.2%|Medium||
|0517|Super Washing Machines||38.6%|Hard||
|0518|Coin Change 2||51.7%|Medium||
|0519|Random Flip Matrix||37.8%|Medium||
|0520|Detect Capital||53.9%|Easy||
|0521|Longest Uncommon Subsequence I||58.4%|Easy||
|0522|Longest Uncommon Subsequence II||34.2%|Medium||
|0523|Continuous Subarray Sum||24.7%|Medium||
|0524|Longest Word in Dictionary through Deleting|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0524.Longest-Word-in-Dictionary-through-Deleting)|49.0%|Medium||
|0525|Contiguous Array||43.4%|Medium||
|0526|Beautiful Arrangement|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0526.Beautiful-Arrangement)|61.8%|Medium||
|0527|Word Abbreviation||56.2%|Hard||
|0528|Random Pick with Weight|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0528.Random-Pick-with-Weight)|44.6%|Medium||
|0529|Minesweeper|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0529.Minesweeper)|61.0%|Medium||
|0530|Minimum Absolute Difference in BST||54.8%|Easy||
|0531|Lonely Pixel I||59.4%|Medium||
|0532|K-diff Pairs in an Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0532.K-diff-Pairs-in-an-Array)|35.1%|Medium||
|0533|Lonely Pixel II||48.1%|Medium||
|0534|Game Play Analysis III||79.1%|Medium||
|0535|Encode and Decode TinyURL||80.9%|Medium||
|0536|Construct Binary Tree from String||50.6%|Medium||
|0537|Complex Number Multiplication|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0537.Complex-Number-Multiplication)|68.3%|Medium||
|0538|Convert BST to Greater Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0538.Convert-BST-to-Greater-Tree)|59.1%|Medium||
|0539|Minimum Time Difference||52.3%|Medium||
|0540|Single Element in a Sorted Array||57.9%|Medium||
|0541|Reverse String II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0541.Reverse-String-II)|49.1%|Easy||
|0542|01 Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0542.01-Matrix)|40.8%|Medium||
|0543|Diameter of Binary Tree||49.1%|Easy||
|0544|Output Contest Matches||75.8%|Medium||
|0545|Boundary of Binary Tree||39.6%|Medium||
|0546|Remove Boxes||44.2%|Hard||
|0547|Number of Provinces|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0547.Number-of-Provinces)|60.3%|Medium||
|0548|Split Array with Equal Sum||47.9%|Medium||
|0549|Binary Tree Longest Consecutive Sequence II||47.2%|Medium||
|0550|Game Play Analysis IV||45.7%|Medium||
|0551|Student Attendance Record I||46.0%|Easy||
|0552|Student Attendance Record II||37.3%|Hard||
|0553|Optimal Division||57.4%|Medium||
|0554|Brick Wall||50.7%|Medium||
|0555|Split Concatenated Strings||42.8%|Medium||
|0556|Next Greater Element III||33.5%|Medium||
|0557|Reverse Words in a String III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0557.Reverse-Words-in-a-String-III)|71.8%|Easy||
|0558|Logical OR of Two Binary Grids Represented as Quad-Trees||45.5%|Medium||
|0559|Maximum Depth of N-ary Tree||69.5%|Easy||
|0560|Subarray Sum Equals K||43.8%|Medium||
|0561|Array Partition I|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0561.Array-Partition-I)|73.1%|Easy||
|0562|Longest Line of Consecutive One in Matrix||46.2%|Medium||
|0563|Binary Tree Tilt|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0563.Binary-Tree-Tilt)|52.9%|Easy||
|0564|Find the Closest Palindrome||20.3%|Hard||
|0565|Array Nesting||55.9%|Medium||
|0566|Reshape the Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0566.Reshape-the-Matrix)|61.0%|Easy||
|0567|Permutation in String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0567.Permutation-in-String)|44.6%|Medium||
|0568|Maximum Vacation Days||41.5%|Hard||
|0569|Median Employee Salary||61.4%|Hard||
|0570|Managers with at Least 5 Direct Reports||66.8%|Medium||
|0571|Find Median Given Frequency of Numbers||45.6%|Hard||
|0572|Subtree of Another Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0572.Subtree-of-Another-Tree)|44.5%|Easy||
|0573|Squirrel Simulation||54.5%|Medium||
|0574|Winning Candidate||51.7%|Medium||
|0575|Distribute Candies|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0575.Distribute-Candies)|62.0%|Easy||
|0576|Out of Boundary Paths||35.9%|Medium||
|0577|Employee Bonus||71.2%|Easy||
|0578|Get Highest Answer Rate Question||41.6%|Medium||
|0579|Find Cumulative Salary of an Employee||38.2%|Hard||
|0580|Count Student Number in Departments||51.5%|Medium||
|0581|Shortest Unsorted Continuous Subarray||31.7%|Medium||
|0582|Kill Process||62.6%|Medium||
|0583|Delete Operation for Two Strings||49.9%|Medium||
|0584|Find Customer Referee||73.9%|Easy||
|0585|Investments in 2016||56.7%|Medium||
|0586|Customer Placing the Largest Number of Orders||75.2%|Easy||
|0587|Erect the Fence||36.4%|Hard||
|0588|Design In-Memory File System||46.6%|Hard||
|0589|N-ary Tree Preorder Traversal||73.3%|Easy||
|0590|N-ary Tree Postorder Traversal||73.3%|Easy||
|0591|Tag Validator||34.6%|Hard||
|0592|Fraction Addition and Subtraction||50.2%|Medium||
|0593|Valid Square||43.3%|Medium||
|0594|Longest Harmonious Subsequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0594.Longest-Harmonious-Subsequence)|51.2%|Easy||
|0595|Big Countries||78.4%|Easy||
|0596|Classes More Than 5 Students||38.7%|Easy||
|0597|Friend Requests I: Overall Acceptance Rate||41.8%|Easy||
|0598|Range Addition II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0598.Range-Addition-II)|50.1%|Easy||
|0599|Minimum Index Sum of Two Lists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0599.Minimum-Index-Sum-of-Two-Lists)|51.6%|Easy||
|0600|Non-negative Integers without Consecutive Ones||34.7%|Hard||
|0601|Human Traffic of Stadium||44.9%|Hard||
|0602|Friend Requests II: Who Has the Most Friends||57.1%|Medium||
|0603|Consecutive Available Seats||66.0%|Easy||
|0604|Design Compressed String Iterator||38.1%|Easy||
|0605|Can Place Flowers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0605.Can-Place-Flowers)|31.8%|Easy||
|0606|Construct String from Binary Tree||55.1%|Easy||
|0607|Sales Person||65.0%|Easy||
|0608|Tree Node||69.4%|Medium||
|0609|Find Duplicate File in System||61.0%|Medium||
|0610|Triangle Judgement||68.3%|Easy||
|0611|Valid Triangle Number||49.2%|Medium||
|0612|Shortest Distance in a Plane||61.1%|Medium||
|0613|Shortest Distance in a Line||79.5%|Easy||
|0614|Second Degree Follower||32.4%|Medium||
|0615|Average Salary: Departments VS Company||52.0%|Hard||
|0616|Add Bold Tag in String||44.6%|Medium||
|0617|Merge Two Binary Trees||75.3%|Easy||
|0618|Students Report By Geography||59.5%|Hard||
|0619|Biggest Single Number||44.8%|Easy||
|0620|Not Boring Movies||69.7%|Easy||
|0621|Task Scheduler||51.7%|Medium||
|0622|Design Circular Queue||45.3%|Medium||
|0623|Add One Row to Tree||50.4%|Medium||
|0624|Maximum Distance in Arrays||39.5%|Medium||
|0625|Minimum Factorization||32.9%|Medium||
|0626|Exchange Seats||65.4%|Medium||
|0627|Swap Salary||77.5%|Easy||
|0628|Maximum Product of Three Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0628.Maximum-Product-of-Three-Numbers)|46.9%|Easy||
|0629|K Inverse Pairs Array||31.7%|Hard||
|0630|Course Schedule III||33.8%|Hard||
|0631|Design Excel Sum Formula||32.1%|Hard||
|0632|Smallest Range Covering Elements from K Lists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0632.Smallest-Range-Covering-Elements-from-K-Lists)|54.1%|Hard||
|0633|Sum of Square Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0633.Sum-of-Square-Numbers)|32.5%|Medium||
|0634|Find the Derangement of An Array||40.5%|Medium||
|0635|Design Log Storage System||59.7%|Medium||
|0636|Exclusive Time of Functions|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0636.Exclusive-Time-of-Functions)|54.3%|Medium||
|0637|Average of Levels in Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0637.Average-of-Levels-in-Binary-Tree)|64.7%|Easy||
|0638|Shopping Offers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0638.Shopping-Offers)|52.8%|Medium||
|0639|Decode Ways II||27.5%|Hard||
|0640|Solve the Equation||42.7%|Medium||
|0641|Design Circular Deque||54.8%|Medium||
|0642|Design Search Autocomplete System||46.1%|Hard||
|0643|Maximum Average Subarray I|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0643.Maximum-Average-Subarray-I)|42.0%|Easy||
|0644|Maximum Average Subarray II||34.0%|Hard||
|0645|Set Mismatch|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0645.Set-Mismatch)|42.5%|Easy||
|0646|Maximum Length of Pair Chain||52.9%|Medium||
|0647|Palindromic Substrings||61.8%|Medium||
|0648|Replace Words|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0648.Replace-Words)|58.4%|Medium||
|0649|Dota2 Senate||39.5%|Medium||
|0650|2 Keys Keyboard||50.0%|Medium||
|0651|4 Keys Keyboard||53.0%|Medium||
|0652|Find Duplicate Subtrees||52.3%|Medium||
|0653|Two Sum IV - Input is a BST|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0653.Two-Sum-IV---Input-is-a-BST)|56.2%|Easy||
|0654|Maximum Binary Tree||81.0%|Medium||
|0655|Print Binary Tree||56.0%|Medium||
|0656|Coin Path||29.6%|Hard||
|0657|Robot Return to Origin||73.7%|Easy||
|0658|Find K Closest Elements|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0658.Find-K-Closest-Elements)|41.9%|Medium||
|0659|Split Array into Consecutive Subsequences||44.3%|Medium||
|0660|Remove 9||54.2%|Hard||
|0661|Image Smoother|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0661.Image-Smoother)|52.3%|Easy||
|0662|Maximum Width of Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0662.Maximum-Width-of-Binary-Tree)|39.9%|Medium||
|0663|Equal Tree Partition||39.7%|Medium||
|0664|Strange Printer||41.4%|Hard||
|0665|Non-decreasing Array||19.7%|Medium||
|0666|Path Sum IV||56.0%|Medium||
|0667|Beautiful Arrangement II||55.1%|Medium||
|0668|Kth Smallest Number in Multiplication Table|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0668.Kth-Smallest-Number-in-Multiplication-Table)|47.7%|Hard||
|0669|Trim a Binary Search Tree||64.5%|Medium||
|0670|Maximum Swap||45.1%|Medium||
|0671|Second Minimum Node In a Binary Tree||42.8%|Easy||
|0672|Bulb Switcher II||51.1%|Medium||
|0673|Number of Longest Increasing Subsequence||38.5%|Medium||
|0674|Longest Continuous Increasing Subsequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0674.Longest-Continuous-Increasing-Subsequence)|46.1%|Easy||
|0675|Cut Off Trees for Golf Event||35.2%|Hard||
|0676|Implement Magic Dictionary|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0676.Implement-Magic-Dictionary)|55.2%|Medium||
|0677|Map Sum Pairs||54.0%|Medium||
|0678|Valid Parenthesis String||31.6%|Medium||
|0679|24 Game||47.2%|Hard||
|0680|Valid Palindrome II||37.1%|Easy||
|0681|Next Closest Time||45.9%|Medium||
|0682|Baseball Game|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0682.Baseball-Game)|66.7%|Easy||
|0683|K Empty Slots||36.0%|Hard||
|0684|Redundant Connection|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0684.Redundant-Connection)|58.9%|Medium||
|0685|Redundant Connection II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0685.Redundant-Connection-II)|32.9%|Hard||
|0686|Repeated String Match||32.8%|Medium||
|0687|Longest Univalue Path||37.2%|Medium||
|0688|Knight Probability in Chessboard||50.0%|Medium||
|0689|Maximum Sum of 3 Non-Overlapping Subarrays||47.2%|Hard||
|0690|Employee Importance||58.7%|Easy||
|0691|Stickers to Spell Word||44.4%|Hard||
|0692|Top K Frequent Words||53.0%|Medium||
|0693|Binary Number with Alternating Bits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0693.Binary-Number-with-Alternating-Bits)|59.8%|Easy||
|0694|Number of Distinct Islands||57.8%|Medium||
|0695|Max Area of Island|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0695.Max-Area-of-Island)|64.6%|Medium||
|0696|Count Binary Substrings||57.4%|Easy||
|0697|Degree of an Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0697.Degree-of-an-Array)|54.4%|Easy||
|0698|Partition to K Equal Sum Subsets||45.6%|Medium||
|0699|Falling Squares|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0699.Falling-Squares)|42.5%|Hard||
|0700|Search in a Binary Search Tree||73.4%|Easy||
|0701|Insert into a Binary Search Tree||75.8%|Medium||
|0702|Search in a Sorted Array of Unknown Size||68.7%|Medium||
|0703|Kth Largest Element in a Stream|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0703.Kth-Largest-Element-in-a-Stream)|50.6%|Easy||
|0704|Binary Search|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0704.Binary-Search)|54.0%|Easy||
|0705|Design HashSet|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0705.Design-HashSet)|64.7%|Easy||
|0706|Design HashMap|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0706.Design-HashMap)|62.7%|Easy||
|0707|Design Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0707.Design-Linked-List)|25.8%|Medium||
|0708|Insert into a Sorted Circular Linked List||32.5%|Medium||
|0709|To Lower Case||80.0%|Easy||
|0710|Random Pick with Blacklist|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0710.Random-Pick-with-Blacklist)|32.7%|Hard||
|0711|Number of Distinct Islands II||49.2%|Hard||
|0712|Minimum ASCII Delete Sum for Two Strings||59.4%|Medium||
|0713|Subarray Product Less Than K|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0713.Subarray-Product-Less-Than-K)|40.5%|Medium||
|0714|Best Time to Buy and Sell Stock with Transaction Fee|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0714.Best-Time-to-Buy-and-Sell-Stock-with-Transaction-Fee)|56.0%|Medium||
|0715|Range Module|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0715.Range-Module)|40.1%|Hard||
|0716|Max Stack||43.0%|Easy||
|0717|1-bit and 2-bit Characters|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0717.1-bit-and-2-bit-Characters)|47.3%|Easy||
|0718|Maximum Length of Repeated Subarray|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0718.Maximum-Length-of-Repeated-Subarray)|50.3%|Medium||
|0719|Find K-th Smallest Pair Distance|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0719.Find-K-th-Smallest-Pair-Distance)|32.5%|Hard||
|0720|Longest Word in Dictionary|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0720.Longest-Word-in-Dictionary)|49.1%|Easy||
|0721|Accounts Merge|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0721.Accounts-Merge)|51.6%|Medium||
|0722|Remove Comments||36.0%|Medium||
|0723|Candy Crush||72.6%|Medium||
|0724|Find Pivot Index|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0724.Find-Pivot-Index)|45.3%|Easy||
|0725|Split Linked List in Parts|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0725.Split-Linked-List-in-Parts)|52.9%|Medium||
|0726|Number of Atoms|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0726.Number-of-Atoms)|51.1%|Hard||
|0727|Minimum Window Subsequence||42.2%|Hard||
|0728|Self Dividing Numbers||75.4%|Easy||
|0729|My Calendar I|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0729.My-Calendar-I)|53.2%|Medium||
|0730|Count Different Palindromic Subsequences||43.4%|Hard||
|0731|My Calendar II||50.4%|Medium||
|0732|My Calendar III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0732.My-Calendar-III)|61.7%|Hard||
|0733|Flood Fill|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0733.Flood-Fill)|55.9%|Easy||
|0734|Sentence Similarity||42.3%|Easy||
|0735|Asteroid Collision|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0735.Asteroid-Collision)|43.3%|Medium||
|0736|Parse Lisp Expression||49.8%|Hard||
|0737|Sentence Similarity II||46.5%|Medium||
|0738|Monotone Increasing Digits||45.6%|Medium||
|0739|Daily Temperatures|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0739.Daily-Temperatures)|64.5%|Medium||
|0740|Delete and Earn||49.3%|Medium||
|0741|Cherry Pickup||35.1%|Hard||
|0742|Closest Leaf in a Binary Tree||44.3%|Medium||
|0743|Network Delay Time||45.4%|Medium||
|0744|Find Smallest Letter Greater Than Target|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0744.Find-Smallest-Letter-Greater-Than-Target)|45.6%|Easy||
|0745|Prefix and Suffix Search|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0745.Prefix-and-Suffix-Search)|35.2%|Hard||
|0746|Min Cost Climbing Stairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0746.Min-Cost-Climbing-Stairs)|50.9%|Easy||
|0747|Largest Number At Least Twice of Others||42.9%|Easy||
|0748|Shortest Completing Word|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0748.Shortest-Completing-Word)|57.5%|Easy||
|0749|Contain Virus||48.2%|Hard||
|0750|Number Of Corner Rectangles||67.1%|Medium||
|0751|IP to CIDR||59.8%|Medium||
|0752|Open the Lock||52.8%|Medium||
|0753|Cracking the Safe|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0753.Cracking-the-Safe)|52.3%|Hard||
|0754|Reach a Number||40.5%|Medium||
|0755|Pour Water||44.1%|Medium||
|0756|Pyramid Transition Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0756.Pyramid-Transition-Matrix)|55.5%|Medium||
|0757|Set Intersection Size At Least Two||42.1%|Hard||
|0758|Bold Words in String||47.3%|Easy||
|0759|Employee Free Time||68.1%|Hard||
|0760|Find Anagram Mappings||81.8%|Easy||
|0761|Special Binary String||58.6%|Hard||
|0762|Prime Number of Set Bits in Binary Representation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0762.Prime-Number-of-Set-Bits-in-Binary-Representation)|64.3%|Easy||
|0763|Partition Labels|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0763.Partition-Labels)|78.0%|Medium||
|0764|Largest Plus Sign||46.5%|Medium||
|0765|Couples Holding Hands|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0765.Couples-Holding-Hands)|55.4%|Hard||
|0766|Toeplitz Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0766.Toeplitz-Matrix)|65.8%|Easy||
|0767|Reorganize String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0767.Reorganize-String)|50.0%|Medium||
|0768|Max Chunks To Make Sorted II||49.7%|Hard||
|0769|Max Chunks To Make Sorted||55.6%|Medium||
|0770|Basic Calculator IV||54.1%|Hard||
|0771|Jewels and Stones|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0771.Jewels-and-Stones)|86.9%|Easy||
|0772|Basic Calculator III||43.5%|Hard||
|0773|Sliding Puzzle||60.7%|Hard||
|0774|Minimize Max Distance to Gas Station||48.2%|Hard||
|0775|Global and Local Inversions||42.6%|Medium||
|0776|Split BST||56.6%|Medium||
|0777|Swap Adjacent in LR String||35.6%|Medium||
|0778|Swim in Rising Water|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0778.Swim-in-Rising-Water)|54.8%|Hard||
|0779|K-th Symbol in Grammar||38.5%|Medium||
|0780|Reaching Points||30.3%|Hard||
|0781|Rabbits in Forest|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0781.Rabbits-in-Forest)|55.4%|Medium||
|0782|Transform to Chessboard||47.0%|Hard||
|0783|Minimum Distance Between BST Nodes||53.8%|Easy||
|0784|Letter Case Permutation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0784.Letter-Case-Permutation)|66.5%|Medium||
|0785|Is Graph Bipartite?|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0785.Is-Graph-Bipartite?)|48.3%|Medium||
|0786|K-th Smallest Prime Fraction|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0786.K-th-Smallest-Prime-Fraction)|42.3%|Hard||
|0787|Cheapest Flights Within K Stops||39.5%|Medium||
|0788|Rotated Digits||57.1%|Easy||
|0789|Escape The Ghosts||58.3%|Medium||
|0790|Domino and Tromino Tiling||40.0%|Medium||
|0791|Custom Sort String||65.9%|Medium||
|0792|Number of Matching Subsequences||48.1%|Medium||
|0793|Preimage Size of Factorial Zeroes Function|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0793.Preimage-Size-of-Factorial-Zeroes-Function)|40.7%|Hard||
|0794|Valid Tic-Tac-Toe State||33.7%|Medium||
|0795|Number of Subarrays with Bounded Maximum||47.5%|Medium||
|0796|Rotate String||49.3%|Easy||
|0797|All Paths From Source to Target||78.5%|Medium||
|0798|Smallest Rotation with Highest Score||44.8%|Hard||
|0799|Champagne Tower||44.0%|Medium||
|0800|Similar RGB Color||62.2%|Easy||
|0801|Minimum Swaps To Make Sequences Increasing||39.0%|Medium||
|0802|Find Eventual Safe States|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0802.Find-Eventual-Safe-States)|49.8%|Medium||
|0803|Bricks Falling When Hit|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0803.Bricks-Falling-When-Hit)|31.4%|Hard||
|0804|Unique Morse Code Words||78.8%|Easy||
|0805|Split Array With Same Average||26.8%|Hard||
|0806|Number of Lines To Write String||65.5%|Easy||
|0807|Max Increase to Keep City Skyline||84.3%|Medium||
|0808|Soup Servings||41.0%|Medium||
|0809|Expressive Words||46.5%|Medium||
|0810|Chalkboard XOR Game||49.6%|Hard||
|0811|Subdomain Visit Count|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0811.Subdomain-Visit-Count)|71.4%|Easy||
|0812|Largest Triangle Area|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0812.Largest-Triangle-Area)|58.9%|Easy||
|0813|Largest Sum of Averages||51.0%|Medium||
|0814|Binary Tree Pruning||72.7%|Medium||
|0815|Bus Routes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0815.Bus-Routes)|43.3%|Hard||
|0816|Ambiguous Coordinates||48.0%|Medium||
|0817|Linked List Components|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0817.Linked-List-Components)|57.7%|Medium||
|0818|Race Car||39.8%|Hard||
|0819|Most Common Word|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0819.Most-Common-Word)|45.5%|Easy||
|0820|Short Encoding of Words||51.5%|Medium||
|0821|Shortest Distance to a Character|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0821.Shortest-Distance-to-a-Character)|70.0%|Easy||
|0822|Card Flipping Game||43.6%|Medium||
|0823|Binary Trees With Factors||36.5%|Medium||
|0824|Goat Latin||66.4%|Easy||
|0825|Friends Of Appropriate Ages||43.9%|Medium||
|0826|Most Profit Assigning Work|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0826.Most-Profit-Assigning-Work)|39.0%|Medium||
|0827|Making A Large Island||47.3%|Hard||
|0828|Count Unique Characters of All Substrings of a Given String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0828.Count-Unique-Characters-of-All-Substrings-of-a-Given-String)|46.8%|Hard||
|0829|Consecutive Numbers Sum||39.3%|Hard||
|0830|Positions of Large Groups|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0830.Positions-of-Large-Groups)|50.4%|Easy||
|0831|Masking Personal Information||44.8%|Medium||
|0832|Flipping an Image|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0832.Flipping-an-Image)|78.1%|Easy||
|0833|Find And Replace in String||51.4%|Medium||
|0834|Sum of Distances in Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0834.Sum-of-Distances-in-Tree)|45.8%|Hard||
|0835|Image Overlap||61.8%|Medium||
|0836|Rectangle Overlap|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0836.Rectangle-Overlap)|44.7%|Easy||
|0837|New 21 Game||35.3%|Medium||
|0838|Push Dominoes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0838.Push-Dominoes)|49.8%|Medium||
|0839|Similar String Groups|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0839.Similar-String-Groups)|40.5%|Hard||
|0840|Magic Squares In Grid||37.8%|Medium||
|0841|Keys and Rooms|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0841.Keys-and-Rooms)|65.4%|Medium||
|0842|Split Array into Fibonacci Sequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0842.Split-Array-into-Fibonacci-Sequence)|36.8%|Medium||
|0843|Guess the Word||46.4%|Hard||
|0844|Backspace String Compare|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0844.Backspace-String-Compare)|46.9%|Easy||
|0845|Longest Mountain in Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0845.Longest-Mountain-in-Array)|38.5%|Medium||
|0846|Hand of Straights||55.3%|Medium||
|0847|Shortest Path Visiting All Nodes||53.5%|Hard||
|0848|Shifting Letters||45.2%|Medium||
|0849|Maximize Distance to Closest Person||44.5%|Medium||
|0850|Rectangle Area II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0850.Rectangle-Area-II)|48.3%|Hard||
|0851|Loud and Rich|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0851.Loud-and-Rich)|52.6%|Medium||
|0852|Peak Index in a Mountain Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0852.Peak-Index-in-a-Mountain-Array)|71.7%|Easy||
|0853|Car Fleet|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0853.Car-Fleet)|43.8%|Medium||
|0854|K-Similar Strings||38.7%|Hard||
|0855|Exam Room||43.4%|Medium||
|0856|Score of Parentheses|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0856.Score-of-Parentheses)|62.4%|Medium||
|0857|Minimum Cost to Hire K Workers||50.4%|Hard||
|0858|Mirror Reflection||59.6%|Medium||
|0859|Buddy Strings||29.5%|Easy||
|0860|Lemonade Change||51.9%|Easy||
|0861|Score After Flipping Matrix||73.6%|Medium||
|0862|Shortest Subarray with Sum at Least K|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0862.Shortest-Subarray-with-Sum-at-Least-K)|25.2%|Hard||
|0863|All Nodes Distance K in Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0863.All-Nodes-Distance-K-in-Binary-Tree)|57.8%|Medium||
|0864|Shortest Path to Get All Keys|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0864.Shortest-Path-to-Get-All-Keys)|41.8%|Hard||
|0865|Smallest Subtree with all the Deepest Nodes||64.8%|Medium||
|0866|Prime Palindrome||25.1%|Medium||
|0867|Transpose Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0867.Transpose-Matrix)|62.1%|Easy||
|0868|Binary Gap||60.9%|Easy||
|0869|Reordered Power of 2||54.1%|Medium||
|0870|Advantage Shuffle||46.7%|Medium||
|0871|Minimum Number of Refueling Stops||32.3%|Hard||
|0872|Leaf-Similar Trees|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0872.Leaf-Similar-Trees)|64.5%|Easy||
|0873|Length of Longest Fibonacci Subsequence||48.1%|Medium||
|0874|Walking Robot Simulation||36.8%|Easy||
|0875|Koko Eating Bananas|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0875.Koko-Eating-Bananas)|53.5%|Medium||
|0876|Middle of the Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0876.Middle-of-the-Linked-List)|69.1%|Easy||
|0877|Stone Game||66.4%|Medium||
|0878|Nth Magical Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0878.Nth-Magical-Number)|28.8%|Hard||
|0879|Profitable Schemes||40.0%|Hard||
|0880|Decoded String at Index|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0880.Decoded-String-at-Index)|28.3%|Medium||
|0881|Boats to Save People|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0881.Boats-to-Save-People)|48.8%|Medium||
|0882|Reachable Nodes In Subdivided Graph||42.5%|Hard||
|0883|Projection Area of 3D Shapes||68.2%|Easy||
|0884|Uncommon Words from Two Sentences|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0884.Uncommon-Words-from-Two-Sentences)|64.0%|Easy||
|0885|Spiral Matrix III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0885.Spiral-Matrix-III)|70.7%|Medium||
|0886|Possible Bipartition||45.1%|Medium||
|0887|Super Egg Drop|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0887.Super-Egg-Drop)|27.1%|Hard||
|0888|Fair Candy Swap|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0888.Fair-Candy-Swap)|58.9%|Easy||
|0889|Construct Binary Tree from Preorder and Postorder Traversal||67.5%|Medium||
|0890|Find and Replace Pattern||74.2%|Medium||
|0891|Sum of Subsequence Widths|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0891.Sum-of-Subsequence-Widths)|32.9%|Hard||
|0892|Surface Area of 3D Shapes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0892.Surface-Area-of-3D-Shapes)|59.7%|Easy||
|0893|Groups of Special-Equivalent Strings||68.4%|Easy||
|0894|All Possible Full Binary Trees||77.0%|Medium||
|0895|Maximum Frequency Stack|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0895.Maximum-Frequency-Stack)|62.3%|Hard||
|0896|Monotonic Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0896.Monotonic-Array)|58.0%|Easy||
|0897|Increasing Order Search Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0897.Increasing-Order-Search-Tree)|74.5%|Easy||
|0898|Bitwise ORs of Subarrays|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0898.Bitwise-ORs-of-Subarrays)|34.1%|Medium||
|0899|Orderly Queue||53.0%|Hard||
|0900|RLE Iterator||55.2%|Medium||
|0901|Online Stock Span|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0901.Online-Stock-Span)|61.3%|Medium||
|0902|Numbers At Most N Given Digit Set||36.1%|Hard||
|0903|Valid Permutations for DI Sequence||54.2%|Hard||
|0904|Fruit Into Baskets|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0904.Fruit-Into-Baskets)|43.0%|Medium||
|0905|Sort Array By Parity||75.0%|Easy||
|0906|Super Palindromes||32.8%|Hard||
|0907|Sum of Subarray Minimums|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0907.Sum-of-Subarray-Minimums)|33.1%|Medium||
|0908|Smallest Range I||66.4%|Easy||
|0909|Snakes and Ladders||39.2%|Medium||
|0910|Smallest Range II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0910.Smallest-Range-II)|31.3%|Medium||
|0911|Online Election|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0911.Online-Election)|51.3%|Medium||
|0912|Sort an Array||64.5%|Medium||
|0913|Cat and Mouse||34.7%|Hard||
|0914|X of a Kind in a Deck of Cards|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0914.X-of-a-Kind-in-a-Deck-of-Cards)|34.3%|Easy||
|0915|Partition Array into Disjoint Intervals||46.1%|Medium||
|0916|Word Subsets||48.2%|Medium||
|0917|Reverse Only Letters||58.7%|Easy||
|0918|Maximum Sum Circular Subarray|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0918.Maximum-Sum-Circular-Subarray)|34.2%|Medium||
|0919|Complete Binary Tree Inserter||58.8%|Medium||
|0920|Number of Music Playlists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0920.Number-of-Music-Playlists)|47.8%|Hard||
|0921|Minimum Add to Make Parentheses Valid|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0921.Minimum-Add-to-Make-Parentheses-Valid)|74.7%|Medium||
|0922|Sort Array By Parity II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0922.Sort-Array-By-Parity-II)|70.5%|Easy||
|0923|3Sum With Multiplicity|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0923.3Sum-With-Multiplicity)|36.2%|Medium||
|0924|Minimize Malware Spread|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0924.Minimize-Malware-Spread)|41.8%|Hard||
|0925|Long Pressed Name|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0925.Long-Pressed-Name)|38.1%|Easy||
|0926|Flip String to Monotone Increasing||53.1%|Medium||
|0927|Three Equal Parts|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0927.Three-Equal-Parts)|34.6%|Hard||
|0928|Minimize Malware Spread II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0928.Minimize-Malware-Spread-II)|41.2%|Hard||
|0929|Unique Email Addresses||67.1%|Easy||
|0930|Binary Subarrays With Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0930.Binary-Subarrays-With-Sum)|44.5%|Medium||
|0931|Minimum Falling Path Sum||63.3%|Medium||
|0932|Beautiful Array||61.1%|Medium||
|0933|Number of Recent Calls|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0933.Number-of-Recent-Calls)|72.0%|Easy||
|0934|Shortest Bridge||49.6%|Medium||
|0935|Knight Dialer||46.4%|Medium||
|0936|Stamping The Sequence||47.3%|Hard||
|0937|Reorder Data in Log Files||54.4%|Easy||
|0938|Range Sum of BST||83.0%|Easy||
|0939|Minimum Area Rectangle||51.9%|Medium||
|0940|Distinct Subsequences II||41.6%|Hard||
|0941|Valid Mountain Array||33.4%|Easy||
|0942|DI String Match|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0942.DI-String-Match)|73.5%|Easy||
|0943|Find the Shortest Superstring||43.4%|Hard||
|0944|Delete Columns to Make Sorted||71.1%|Easy||
|0945|Minimum Increment to Make Array Unique||46.7%|Medium||
|0946|Validate Stack Sequences|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0946.Validate-Stack-Sequences)|63.6%|Medium||
|0947|Most Stones Removed with Same Row or Column|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0947.Most-Stones-Removed-with-Same-Row-or-Column)|55.5%|Medium||
|0948|Bag of Tokens||46.1%|Medium||
|0949|Largest Time for Given Digits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0949.Largest-Time-for-Given-Digits)|36.2%|Medium||
|0950|Reveal Cards In Increasing Order||75.4%|Medium||
|0951|Flip Equivalent Binary Trees||65.7%|Medium||
|0952|Largest Component Size by Common Factor|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0952.Largest-Component-Size-by-Common-Factor)|36.1%|Hard||
|0953|Verifying an Alien Dictionary|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0953.Verifying-an-Alien-Dictionary)|52.3%|Easy||
|0954|Array of Doubled Pairs||35.2%|Medium||
|0955|Delete Columns to Make Sorted II||33.8%|Medium||
|0956|Tallest Billboard||39.9%|Hard||
|0957|Prison Cells After N Days||40.1%|Medium||
|0958|Check Completeness of a Binary Tree||52.5%|Medium||
|0959|Regions Cut By Slashes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0959.Regions-Cut-By-Slashes)|67.0%|Medium||
|0960|Delete Columns to Make Sorted III||54.8%|Hard||
|0961|N-Repeated Element in Size 2N Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0961.N-Repeated-Element-in-Size-2N-Array)|74.4%|Easy||
|0962|Maximum Width Ramp||46.5%|Medium||
|0963|Minimum Area Rectangle II||51.8%|Medium||
|0964|Least Operators to Express Number||45.0%|Hard||
|0965|Univalued Binary Tree||67.8%|Easy||
|0966|Vowel Spellchecker||48.0%|Medium||
|0967|Numbers With Same Consecutive Differences||44.7%|Medium||
|0968|Binary Tree Cameras|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0968.Binary-Tree-Cameras)|38.7%|Hard||
|0969|Pancake Sorting|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0969.Pancake-Sorting)|68.6%|Medium||
|0970|Powerful Integers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0970.Powerful-Integers)|40.0%|Medium||
|0971|Flip Binary Tree To Match Preorder Traversal||46.1%|Medium||
|0972|Equal Rational Numbers||42.1%|Hard||
|0973|K Closest Points to Origin|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0973.K-Closest-Points-to-Origin)|64.6%|Medium||
|0974|Subarray Sums Divisible by K||50.8%|Medium||
|0975|Odd Even Jump||41.6%|Hard||
|0976|Largest Perimeter Triangle|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0976.Largest-Perimeter-Triangle)|58.5%|Easy||
|0977|Squares of a Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0977.Squares-of-a-Sorted-Array)|72.1%|Easy||
|0978|Longest Turbulent Subarray|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0978.Longest-Turbulent-Subarray)|46.7%|Medium||
|0979|Distribute Coins in Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0979.Distribute-Coins-in-Binary-Tree)|69.6%|Medium||
|0980|Unique Paths III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0980.Unique-Paths-III)|77.2%|Hard||
|0981|Time Based Key-Value Store|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0981.Time-Based-Key-Value-Store)|54.1%|Medium||
|0982|Triples with Bitwise AND Equal To Zero||56.2%|Hard||
|0983|Minimum Cost For Tickets||62.7%|Medium||
|0984|String Without AAA or BBB|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0984.String-Without-AAA-or-BBB)|38.6%|Medium||
|0985|Sum of Even Numbers After Queries|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0985.Sum-of-Even-Numbers-After-Queries)|60.7%|Easy||
|0986|Interval List Intersections|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0986.Interval-List-Intersections)|68.3%|Medium||
|0987|Vertical Order Traversal of a Binary Tree||38.7%|Hard||
|0988|Smallest String Starting From Leaf||46.7%|Medium||
|0989|Add to Array-Form of Integer|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0989.Add-to-Array-Form-of-Integer)|44.8%|Easy||
|0990|Satisfiability of Equality Equations|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0990.Satisfiability-of-Equality-Equations)|46.7%|Medium||
|0991|Broken Calculator||46.5%|Medium||
|0992|Subarrays with K Different Integers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0992.Subarrays-with-K-Different-Integers)|50.6%|Hard||
|0993|Cousins in Binary Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0993.Cousins-in-Binary-Tree)|52.3%|Easy||
|0994|Rotting Oranges||49.6%|Medium||
|0995|Minimum Number of K Consecutive Bit Flips|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0995.Minimum-Number-of-K-Consecutive-Bit-Flips)|49.7%|Hard||
|0996|Number of Squareful Arrays|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0996.Number-of-Squareful-Arrays)|48.3%|Hard||
|0997|Find the Town Judge||49.8%|Easy||
|0998|Maximum Binary Tree II||63.8%|Medium||
|0999|Available Captures for Rook|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0999.Available-Captures-for-Rook)|67.7%|Easy||
|1000|Minimum Cost to Merge Stones||40.4%|Hard||
|1001|Grid Illumination||36.8%|Hard||
|1002|Find Common Characters|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1002.Find-Common-Characters)|68.5%|Easy||
|1003|Check If Word Is Valid After Substitutions|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1003.Check-If-Word-Is-Valid-After-Substitutions)|56.3%|Medium||
|1004|Max Consecutive Ones III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1004.Max-Consecutive-Ones-III)|60.7%|Medium||
|1005|Maximize Sum Of Array After K Negations|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1005.Maximize-Sum-Of-Array-After-K-Negations)|52.3%|Easy||
|1006|Clumsy Factorial||53.6%|Medium||
|1007|Minimum Domino Rotations For Equal Row||50.9%|Medium||
|1008|Construct Binary Search Tree from Preorder Traversal||78.8%|Medium||
|1009|Complement of Base 10 Integer||61.5%|Easy||
|1010|Pairs of Songs With Total Durations Divisible by 60||50.1%|Medium||
|1011|Capacity To Ship Packages Within D Days|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1011.Capacity-To-Ship-Packages-Within-D-Days)|59.7%|Medium||
|1012|Numbers With Repeated Digits||37.8%|Hard||
|1013|Partition Array Into Three Parts With Equal Sum||49.1%|Easy||
|1014|Best Sightseeing Pair||52.9%|Medium||
|1015|Smallest Integer Divisible by K||41.8%|Medium||
|1016|Binary String With Substrings Representing 1 To N||59.0%|Medium||
|1017|Convert to Base -2|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1017.Convert-to-Base--2)|59.6%|Medium||
|1018|Binary Prefix Divisible By 5|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1018.Binary-Prefix-Divisible-By-5)|47.8%|Easy||
|1019|Next Greater Node In Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1019.Next-Greater-Node-In-Linked-List)|58.2%|Medium||
|1020|Number of Enclaves|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1020.Number-of-Enclaves)|58.9%|Medium||
|1021|Remove Outermost Parentheses|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1021.Remove-Outermost-Parentheses)|78.9%|Easy||
|1022|Sum of Root To Leaf Binary Numbers||71.5%|Easy||
|1023|Camelcase Matching||57.4%|Medium||
|1024|Video Stitching||49.0%|Medium||
|1025|Divisor Game|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1025.Divisor-Game)|66.2%|Easy||
|1026|Maximum Difference Between Node and Ancestor|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1026.Maximum-Difference-Between-Node-and-Ancestor)|69.4%|Medium||
|1027|Longest Arithmetic Subsequence||49.8%|Medium||
|1028|Recover a Tree From Preorder Traversal|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1028.Recover-a-Tree-From-Preorder-Traversal)|70.9%|Hard||
|1029|Two City Scheduling||57.8%|Medium||
|1030|Matrix Cells in Distance Order|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1030.Matrix-Cells-in-Distance-Order)|66.9%|Easy||
|1031|Maximum Sum of Two Non-Overlapping Subarrays||58.9%|Medium||
|1032|Stream of Characters||48.7%|Hard||
|1033|Moving Stones Until Consecutive||43.1%|Easy||
|1034|Coloring A Border||45.6%|Medium||
|1035|Uncrossed Lines||56.0%|Medium||
|1036|Escape a Large Maze||34.8%|Hard||
|1037|Valid Boomerang|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1037.Valid-Boomerang)|37.8%|Easy||
|1038|Binary Search Tree to Greater Sum Tree|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1038.Binary-Search-Tree-to-Greater-Sum-Tree)|82.3%|Medium||
|1039|Minimum Score Triangulation of Polygon||50.1%|Medium||
|1040|Moving Stones Until Consecutive II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1040.Moving-Stones-Until-Consecutive-II)|54.1%|Medium||
|1041|Robot Bounded In Circle||55.0%|Medium||
|1042|Flower Planting With No Adjacent||48.7%|Medium||
|1043|Partition Array for Maximum Sum||66.8%|Medium||
|1044|Longest Duplicate Substring||31.5%|Hard||
|1045|Customers Who Bought All Products||68.4%|Medium||
|1046|Last Stone Weight||62.4%|Easy||
|1047|Remove All Adjacent Duplicates In String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1047.Remove-All-Adjacent-Duplicates-In-String)|70.6%|Easy||
|1048|Longest String Chain||55.4%|Medium||
|1049|Last Stone Weight II|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1049.Last-Stone-Weight-II)|45.2%|Medium||
|1050|Actors and Directors Who Cooperated At Least Three Times||72.2%|Easy||
|1051|Height Checker|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1051.Height-Checker)|72.2%|Easy||
|1052|Grumpy Bookstore Owner|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1052.Grumpy-Bookstore-Owner)|55.8%|Medium||
|1053|Previous Permutation With One Swap||51.0%|Medium||
|1054|Distant Barcodes|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1054.Distant-Barcodes)|44.2%|Medium||
|1055|Shortest Way to Form String||57.2%|Medium||
|1056|Confusing Number||47.1%|Easy||
|1057|Campus Bikes||57.7%|Medium||
|1058|Minimize Rounding Error to Meet Target||42.9%|Medium||
|1059|All Paths from Source Lead to Destination||43.2%|Medium||
|1060|Missing Element in Sorted Array||54.7%|Medium||
|1061|Lexicographically Smallest Equivalent String||66.7%|Medium||
|1062|Longest Repeating Substring||58.3%|Medium||
|1063|Number of Valid Subarrays||72.1%|Hard||
|1064|Fixed Point||65.3%|Easy||
|1065|Index Pairs of a String||61.0%|Easy||
|1066|Campus Bikes II||54.1%|Medium||
|1067|Digit Count in Range||41.3%|Hard||
|1068|Product Sales Analysis I||82.2%|Easy||
|1069|Product Sales Analysis II||83.3%|Easy||
|1070|Product Sales Analysis III||49.8%|Medium||
|1071|Greatest Common Divisor of Strings||51.3%|Easy||
|1072|Flip Columns For Maximum Number of Equal Rows||61.3%|Medium||
|1073|Adding Two Negabinary Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1073.Adding-Two-Negabinary-Numbers)|34.8%|Medium||
|1074|Number of Submatrices That Sum to Target|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1074.Number-of-Submatrices-That-Sum-to-Target)|61.7%|Hard||
|1075|Project Employees I||66.2%|Easy||
|1076|Project Employees II||53.0%|Easy||
|1077|Project Employees III||77.4%|Medium||
|1078|Occurrences After Bigram|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1078.Occurrences-After-Bigram)|64.9%|Easy||
|1079|Letter Tile Possibilities|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1079.Letter-Tile-Possibilities)|75.9%|Medium||
|1080|Insufficient Nodes in Root to Leaf Paths||49.9%|Medium||
|1081|Smallest Subsequence of Distinct Characters||53.4%|Medium||
|1082|Sales Analysis I||73.6%|Easy||
|1083|Sales Analysis II||50.8%|Easy||
|1084|Sales Analysis III||54.7%|Easy||
|1085|Sum of Digits in the Minimum Number||75.1%|Easy||
|1086|High Five||78.2%|Easy||
|1087|Brace Expansion||63.0%|Medium||
|1088|Confusing Number II||45.4%|Hard||
|1089|Duplicate Zeros|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1089.Duplicate-Zeros)|51.9%|Easy||
|1090|Largest Values From Labels||60.0%|Medium||
|1091|Shortest Path in Binary Matrix||39.1%|Medium||
|1092|Shortest Common Supersequence ||53.0%|Hard||
|1093|Statistics from a Large Sample|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1093.Statistics-from-a-Large-Sample)|49.3%|Medium||
|1094|Car Pooling||59.0%|Medium||
|1095|Find in Mountain Array||36.1%|Hard||
|1096|Brace Expansion II||62.5%|Hard||
|1097|Game Play Analysis V||56.3%|Hard||
|1098|Unpopular Books||45.4%|Medium||
|1099|Two Sum Less Than K||60.8%|Easy||
|1100|Find K-Length Substrings With No Repeated Characters||73.2%|Medium||
|1101|The Earliest Moment When Everyone Become Friends||68.0%|Medium||
|1102|Path With Maximum Minimum Value||50.4%|Medium||
|1103|Distribute Candies to People||63.5%|Easy||
|1104|Path In Zigzag Labelled Binary Tree||73.1%|Medium||
|1105|Filling Bookcase Shelves|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1105.Filling-Bookcase-Shelves)|57.5%|Medium||
|1106|Parsing A Boolean Expression||59.2%|Hard||
|1107|New Users Daily Count||45.8%|Medium||
|1108|Defanging an IP Address|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1108.Defanging-an-IP-Address)|88.4%|Easy||
|1109|Corporate Flight Bookings||54.3%|Medium||
|1110|Delete Nodes And Return Forest|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1110.Delete-Nodes-And-Return-Forest)|67.7%|Medium||
|1111|Maximum Nesting Depth of Two Valid Parentheses Strings|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1111.Maximum-Nesting-Depth-of-Two-Valid-Parentheses-Strings)|72.8%|Medium||
|1112|Highest Grade For Each Student||71.9%|Medium||
|1113|Reported Posts||65.7%|Easy||
|1114|Print in Order||67.1%|Easy||
|1115|Print FooBar Alternately||59.0%|Medium||
|1116|Print Zero Even Odd||57.7%|Medium||
|1117|Building H2O||53.1%|Medium||
|1118|Number of Days in a Month||57.5%|Easy||
|1119|Remove Vowels from a String||90.5%|Easy||
|1120|Maximum Average Subtree||63.5%|Medium||
|1121|Divide Array Into Increasing Sequences||58.0%|Hard||
|1122|Relative Sort Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1122.Relative-Sort-Array)|68.1%|Easy||
|1123|Lowest Common Ancestor of Deepest Leaves|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1123.Lowest-Common-Ancestor-of-Deepest-Leaves)|67.9%|Medium||
|1124|Longest Well-Performing Interval||33.3%|Medium||
|1125|Smallest Sufficient Team||47.1%|Hard||
|1126|Active Businesses||68.6%|Medium||
|1127|User Purchase Platform||50.2%|Hard||
|1128|Number of Equivalent Domino Pairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1128.Number-of-Equivalent-Domino-Pairs)|46.6%|Easy||
|1129|Shortest Path with Alternating Colors||40.2%|Medium||
|1130|Minimum Cost Tree From Leaf Values||67.3%|Medium||
|1131|Maximum of Absolute Value Expression||52.1%|Medium||
|1132|Reported Posts II||34.6%|Medium||
|1133|Largest Unique Number||67.2%|Easy||
|1134|Armstrong Number||78.2%|Easy||
|1135|Connecting Cities With Minimum Cost||59.3%|Medium||
|1136|Parallel Courses||61.2%|Hard||
|1137|N-th Tribonacci Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1137.N-th-Tribonacci-Number)|56.0%|Easy||
|1138|Alphabet Board Path||51.0%|Medium||
|1139|Largest 1-Bordered Square||48.6%|Medium||
|1140|Stone Game II||64.7%|Medium||
|1141|User Activity for the Past 30 Days I||54.4%|Easy||
|1142|User Activity for the Past 30 Days II||35.4%|Easy||
|1143|Longest Common Subsequence||58.6%|Medium||
|1144|Decrease Elements To Make Array Zigzag||46.1%|Medium||
|1145|Binary Tree Coloring Game|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1145.Binary-Tree-Coloring-Game)|51.4%|Medium||
|1146|Snapshot Array||36.9%|Medium||
|1147|Longest Chunked Palindrome Decomposition||59.5%|Hard||
|1148|Article Views I||77.1%|Easy||
|1149|Article Views II||48.3%|Medium||
|1150|Check If a Number Is Majority Element in a Sorted Array||57.5%|Easy||
|1151|Minimum Swaps to Group All 1's Together||58.5%|Medium||
|1152|Analyze User Website Visit Pattern||43.3%|Medium||
|1153|String Transforms Into Another String||36.0%|Hard||
|1154|Day of the Year|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1154.Day-of-the-Year)|49.3%|Easy||
|1155|Number of Dice Rolls With Target Sum||47.5%|Medium||
|1156|Swap For Longest Repeated Character Substring||47.3%|Medium||
|1157|Online Majority Element In Subarray|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1157.Online-Majority-Element-In-Subarray)|39.6%|Hard||
|1158|Market Analysis I||63.9%|Medium||
|1159|Market Analysis II||55.2%|Hard||
|1160|Find Words That Can Be Formed by Characters|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1160.Find-Words-That-Can-Be-Formed-by-Characters)|67.9%|Easy||
|1161|Maximum Level Sum of a Binary Tree||69.8%|Medium||
|1162|As Far from Land as Possible||45.2%|Medium||
|1163|Last Substring in Lexicographical Order||36.2%|Hard||
|1164|Product Price at a Given Date||68.4%|Medium||
|1165|Single-Row Keyboard||84.8%|Easy||
|1166|Design File System||58.5%|Medium||
|1167|Minimum Cost to Connect Sticks||64.4%|Medium||
|1168|Optimize Water Distribution in a Village||61.5%|Hard||
|1169|Invalid Transactions||31.2%|Medium||
|1170|Compare Strings by Frequency of the Smallest Character|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1170.Compare-Strings-by-Frequency-of-the-Smallest-Character)|59.5%|Medium||
|1171|Remove Zero Sum Consecutive Nodes from Linked List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1171.Remove-Zero-Sum-Consecutive-Nodes-from-Linked-List)|41.4%|Medium||
|1172|Dinner Plate Stacks||37.9%|Hard||
|1173|Immediate Food Delivery I||82.3%|Easy||
|1174|Immediate Food Delivery II||61.7%|Medium||
|1175|Prime Arrangements|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1175.Prime-Arrangements)|52.0%|Easy||
|1176|Diet Plan Performance||54.1%|Easy||
|1177|Can Make Palindrome from Substring||36.1%|Medium||
|1178|Number of Valid Words for Each Puzzle||38.8%|Hard||
|1179|Reformat Department Table||81.9%|Easy||
|1180|Count Substrings with Only One Distinct Letter||77.6%|Easy||
|1181|Before and After Puzzle||44.5%|Medium||
|1182|Shortest Distance to Target Color||53.4%|Medium||
|1183|Maximum Number of Ones||57.2%|Hard||
|1184|Distance Between Bus Stops|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1184.Distance-Between-Bus-Stops)|54.1%|Easy||
|1185|Day of the Week|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1185.Day-of-the-Week)|61.6%|Easy||
|1186|Maximum Subarray Sum with One Deletion||38.7%|Medium||
|1187|Make Array Strictly Increasing||41.8%|Hard||
|1188|Design Bounded Blocking Queue||72.7%|Medium||
|1189|Maximum Number of Balloons|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1189.Maximum-Number-of-Balloons)|61.7%|Easy||
|1190|Reverse Substrings Between Each Pair of Parentheses||64.1%|Medium||
|1191|K-Concatenation Maximum Sum||25.3%|Medium||
|1192|Critical Connections in a Network||50.0%|Hard||
|1193|Monthly Transactions I||69.3%|Medium||
|1194|Tournament Winners||51.9%|Hard||
|1195|Fizz Buzz Multithreaded||70.5%|Medium||
|1196|How Many Apples Can You Put into the Basket||68.0%|Easy||
|1197|Minimum Knight Moves||37.2%|Medium||
|1198|Find Smallest Common Element in All Rows||75.2%|Medium||
|1199|Minimum Time to Build Blocks||38.5%|Hard||
|1200|Minimum Absolute Difference|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1200.Minimum-Absolute-Difference)|66.9%|Easy||
|1201|Ugly Number III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1201.Ugly-Number-III)|26.4%|Medium||
|1202|Smallest String With Swaps|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1202.Smallest-String-With-Swaps)|48.7%|Medium||
|1203|Sort Items by Groups Respecting Dependencies|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1203.Sort-Items-by-Groups-Respecting-Dependencies)|48.9%|Hard||
|1204|Last Person to Fit in the Elevator||71.6%|Medium||
|1205|Monthly Transactions II||45.8%|Medium||
|1206|Design Skiplist||59.1%|Hard||
|1207|Unique Number of Occurrences|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1207.Unique-Number-of-Occurrences)|71.7%|Easy||
|1208|Get Equal Substrings Within Budget|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1208.Get-Equal-Substrings-Within-Budget)|43.9%|Medium||
|1209|Remove All Adjacent Duplicates in String II||57.5%|Medium||
|1210|Minimum Moves to Reach Target with Rotations||46.4%|Hard||
|1211|Queries Quality and Percentage||69.8%|Easy||
|1212|Team Scores in Football Tournament||56.8%|Medium||
|1213|Intersection of Three Sorted Arrays||79.3%|Easy||
|1214|Two Sum BSTs||67.7%|Medium||
|1215|Stepping Numbers||43.3%|Medium||
|1216|Valid Palindrome III||49.7%|Hard||
|1217|Minimum Cost to Move Chips to The Same Position|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1217.Minimum-Cost-to-Move-Chips-to-The-Same-Position)|71.2%|Easy||
|1218|Longest Arithmetic Subsequence of Given Difference||46.5%|Medium||
|1219|Path with Maximum Gold||65.7%|Medium||
|1220|Count Vowels Permutation||54.2%|Hard||
|1221|Split a String in Balanced Strings|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1221.Split-a-String-in-Balanced-Strings)|84.1%|Easy||
|1222|Queens That Can Attack the King||69.4%|Medium||
|1223|Dice Roll Simulation||46.7%|Hard||
|1224|Maximum Equal Frequency||34.5%|Hard||
|1225|Report Contiguous Dates||62.6%|Hard||
|1226|The Dining Philosophers||59.8%|Medium||
|1227|Airplane Seat Assignment Probability||62.0%|Medium||
|1228|Missing Number In Arithmetic Progression||51.5%|Easy||
|1229|Meeting Scheduler||54.4%|Medium||
|1230|Toss Strange Coins||49.9%|Medium||
|1231|Divide Chocolate||53.5%|Hard||
|1232|Check If It Is a Straight Line|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1232.Check-If-It-Is-a-Straight-Line)|43.7%|Easy||
|1233|Remove Sub-Folders from the Filesystem||61.9%|Medium||
|1234|Replace the Substring for Balanced String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1234.Replace-the-Substring-for-Balanced-String)|34.3%|Medium||
|1235|Maximum Profit in Job Scheduling|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1235.Maximum-Profit-in-Job-Scheduling)|47.3%|Hard||
|1236|Web Crawler||64.7%|Medium||
|1237|Find Positive Integer Solution for a Given Equation||69.6%|Medium||
|1238|Circular Permutation in Binary Representation||66.1%|Medium||
|1239|Maximum Length of a Concatenated String with Unique Characters||49.5%|Medium||
|1240|Tiling a Rectangle with the Fewest Squares||52.4%|Hard||
|1241|Number of Comments per Post||67.7%|Easy||
|1242|Web Crawler Multithreaded||47.7%|Medium||
|1243|Array Transformation||50.2%|Easy||
|1244|Design A Leaderboard||65.7%|Medium||
|1245|Tree Diameter||61.3%|Medium||
|1246|Palindrome Removal||45.9%|Hard||
|1247|Minimum Swaps to Make Strings Equal||62.3%|Medium||
|1248|Count Number of Nice Subarrays||56.1%|Medium||
|1249|Minimum Remove to Make Valid Parentheses||63.7%|Medium||
|1250|Check If It Is a Good Array||56.2%|Hard||
|1251|Average Selling Price||82.5%|Easy||
|1252|Cells with Odd Values in a Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1252.Cells-with-Odd-Values-in-a-Matrix)|78.8%|Easy||
|1253|Reconstruct a 2-Row Binary Matrix||41.6%|Medium||
|1254|Number of Closed Islands|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1254.Number-of-Closed-Islands)|61.5%|Medium||
|1255|Maximum Score Words Formed by Letters||69.8%|Hard||
|1256|Encode Number||67.6%|Medium||
|1257|Smallest Common Region||60.8%|Medium||
|1258|Synonymous Sentences||63.7%|Medium||
|1259|Handshakes That Don't Cross||54.2%|Hard||
|1260|Shift 2D Grid|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1260.Shift-2D-Grid)|61.9%|Easy||
|1261|Find Elements in a Contaminated Binary Tree||74.6%|Medium||
|1262|Greatest Sum Divisible by Three||49.5%|Medium||
|1263|Minimum Moves to Move a Box to Their Target Location||43.2%|Hard||
|1264|Page Recommendations||69.2%|Medium||
|1265|Print Immutable Linked List in Reverse||94.4%|Medium||
|1266|Minimum Time Visiting All Points|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1266.Minimum-Time-Visiting-All-Points)|79.5%|Easy||
|1267|Count Servers that Communicate||57.8%|Medium||
|1268|Search Suggestions System||64.6%|Medium||
|1269|Number of Ways to Stay in the Same Place After Some Steps||43.2%|Hard||
|1270|All People Report to the Given Manager||88.2%|Medium||
|1271|Hexspeak||55.4%|Easy||
|1272|Remove Interval||57.9%|Medium||
|1273|Delete Tree Nodes||62.3%|Medium||
|1274|Number of Ships in a Rectangle||66.2%|Hard||
|1275|Find Winner on a Tic Tac Toe Game|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1275.Find-Winner-on-a-Tic-Tac-Toe-Game)|52.9%|Easy||
|1276|Number of Burgers with No Waste of Ingredients||50.0%|Medium||
|1277|Count Square Submatrices with All Ones||72.9%|Medium||
|1278|Palindrome Partitioning III||60.9%|Hard||
|1279|Traffic Light Controlled Intersection||75.7%|Easy||
|1280|Students and Examinations||74.7%|Easy||
|1281|Subtract the Product and Sum of Digits of an Integer|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer)|85.6%|Easy||
|1282|Group the People Given the Group Size They Belong To||84.4%|Medium||
|1283|Find the Smallest Divisor Given a Threshold|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1283.Find-the-Smallest-Divisor-Given-a-Threshold)|49.4%|Medium||
|1284|Minimum Number of Flips to Convert Binary Matrix to Zero Matrix||70.4%|Hard||
|1285|Find the Start and End Number of Continuous Ranges||87.0%|Medium||
|1286|Iterator for Combination||70.9%|Medium||
|1287|Element Appearing More Than 25% In Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1287.Element-Appearing-More-Than-25%-In-Sorted-Array)|60.2%|Easy||
|1288|Remove Covered Intervals||57.2%|Medium||
|1289|Minimum Falling Path Sum II||62.4%|Hard||
|1290|Convert Binary Number in a Linked List to Integer|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1290.Convert-Binary-Number-in-a-Linked-List-to-Integer)|81.8%|Easy||
|1291|Sequential Digits||57.4%|Medium||
|1292|Maximum Side Length of a Square with Sum Less than or Equal to Threshold||50.5%|Medium||
|1293|Shortest Path in a Grid with Obstacles Elimination||42.9%|Hard||
|1294|Weather Type in Each Country||66.7%|Easy||
|1295|Find Numbers with Even Number of Digits|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1295.Find-Numbers-with-Even-Number-of-Digits)|79.2%|Easy||
|1296|Divide Array in Sets of K Consecutive Numbers||55.4%|Medium||
|1297|Maximum Number of Occurrences of a Substring||49.8%|Medium||
|1298|Maximum Candies You Can Get from Boxes||59.7%|Hard||
|1299|Replace Elements with Greatest Element on Right Side|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1299.Replace-Elements-with-Greatest-Element-on-Right-Side)|74.5%|Easy||
|1300|Sum of Mutated Array Closest to Target|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1300.Sum-of-Mutated-Array-Closest-to-Target)|43.2%|Medium||
|1301|Number of Paths with Max Score||38.1%|Hard||
|1302|Deepest Leaves Sum|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1302.Deepest-Leaves-Sum)|84.2%|Medium||
|1303|Find the Team Size||89.5%|Easy||
|1304|Find N Unique Integers Sum up to Zero|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1304.Find-N-Unique-Integers-Sum-up-to-Zero)|76.9%|Easy||
|1305|All Elements in Two Binary Search Trees|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1305.All-Elements-in-Two-Binary-Search-Trees)|77.8%|Medium||
|1306|Jump Game III|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1306.Jump-Game-III)|62.6%|Medium||
|1307|Verbal Arithmetic Puzzle||36.9%|Hard||
|1308|Running Total for Different Genders||87.4%|Medium||
|1309|Decrypt String from Alphabet to Integer Mapping||77.3%|Easy||
|1310|XOR Queries of a Subarray||69.2%|Medium||
|1311|Get Watched Videos by Your Friends||44.1%|Medium||
|1312|Minimum Insertion Steps to Make a String Palindrome||59.5%|Hard||
|1313|Decompress Run-Length Encoded List|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1313.Decompress-Run-Length-Encoded-List)|85.4%|Easy||
|1314|Matrix Block Sum||73.6%|Medium||
|1315|Sum of Nodes with Even-Valued Grandparent||84.2%|Medium||
|1316|Distinct Echo Substrings||49.6%|Hard||
|1317|Convert Integer to the Sum of Two No-Zero Integers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1317.Convert-Integer-to-the-Sum-of-Two-No-Zero-Integers)|56.8%|Easy||
|1318|Minimum Flips to Make a OR b Equal to c||63.9%|Medium||
|1319|Number of Operations to Make Network Connected|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1319.Number-of-Operations-to-Make-Network-Connected)|55.3%|Medium||
|1320|Minimum Distance to Type a Word Using Two Fingers||62.1%|Hard||
|1321|Restaurant Growth||71.0%|Medium||
|1322|Ads Performance||58.0%|Easy||
|1323|Maximum 69 Number||78.0%|Easy||
|1324|Print Words Vertically||58.8%|Medium||
|1325|Delete Leaves With a Given Value||73.7%|Medium||
|1326|Minimum Number of Taps to Open to Water a Garden||46.4%|Hard||
|1327|List the Products Ordered in a Period||77.5%|Easy||
|1328|Break a Palindrome||46.1%|Medium||
|1329|Sort the Matrix Diagonally|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1329.Sort-the-Matrix-Diagonally)|81.8%|Medium||
|1330|Reverse Subarray To Maximize Array Value||36.6%|Hard||
|1331|Rank Transform of an Array||57.6%|Easy||
|1332|Remove Palindromic Subsequences||62.7%|Easy||
|1333|Filter Restaurants by Vegan-Friendly, Price and Distance||57.1%|Medium||
|1334|Find the City With the Smallest Number of Neighbors at a Threshold Distance||46.8%|Medium||
|1335|Minimum Difficulty of a Job Schedule||57.0%|Hard||
|1336|Number of Transactions per Visit||47.9%|Hard||
|1337|The K Weakest Rows in a Matrix||70.5%|Easy||
|1338|Reduce Array Size to The Half||67.6%|Medium||
|1339|Maximum Product of Splitted Binary Tree||38.3%|Medium||
|1340|Jump Game V||59.2%|Hard||
|1341|Movie Rating||58.6%|Medium||
|1342|Number of Steps to Reduce a Number to Zero||85.6%|Easy||
|1343|Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold||64.7%|Medium||
|1344|Angle Between Hands of a Clock||61.3%|Medium||
|1345|Jump Game IV||41.9%|Hard||
|1346|Check If N and Its Double Exist||36.3%|Easy||
|1347|Minimum Number of Steps to Make Two Strings Anagram||75.3%|Medium||
|1348|Tweet Counts Per Frequency||35.3%|Medium||
|1349|Maximum Students Taking Exam||44.2%|Hard||
|1350|Students With Invalid Departments||90.4%|Easy||
|1351|Count Negative Numbers in a Sorted Matrix||76.0%|Easy||
|1352|Product of the Last K Numbers||44.2%|Medium||
|1353|Maximum Number of Events That Can Be Attended||30.5%|Medium||
|1354|Construct Target Array With Multiple Sums||31.3%|Hard||
|1355|Activity Participants||74.4%|Medium||
|1356|Sort Integers by The Number of 1 Bits||69.7%|Easy||
|1357|Apply Discount Every n Orders||66.7%|Medium||
|1358|Number of Substrings Containing All Three Characters||60.5%|Medium||
|1359|Count All Valid Pickup and Delivery Options||56.5%|Hard||
|1360|Number of Days Between Two Dates||46.9%|Easy||
|1361|Validate Binary Tree Nodes||43.9%|Medium||
|1362|Closest Divisors||57.6%|Medium||
|1363|Largest Multiple of Three||34.1%|Hard||
|1364|Number of Trusted Contacts of a Customer||78.4%|Medium||
|1365|How Many Numbers Are Smaller Than the Current Number||86.0%|Easy||
|1366|Rank Teams by Votes||55.1%|Medium||
|1367|Linked List in Binary Tree||41.2%|Medium||
|1368|Minimum Cost to Make at Least One Valid Path in a Grid||57.0%|Hard||
|1369|Get the Second Most Recent Activity||68.7%|Hard||
|1370|Increasing Decreasing String||76.6%|Easy||
|1371|Find the Longest Substring Containing Vowels in Even Counts||61.0%|Medium||
|1372|Longest ZigZag Path in a Binary Tree||54.8%|Medium||
|1373|Maximum Sum BST in Binary Tree||37.3%|Hard||
|1374|Generate a String With Characters That Have Odd Counts||76.2%|Easy||
|1375|Bulb Switcher III||64.1%|Medium||
|1376|Time Needed to Inform All Employees||56.4%|Medium||
|1377|Frog Position After T Seconds||35.2%|Hard||
|1378|Replace Employee ID With The Unique Identifier||89.9%|Easy||
|1379|Find a Corresponding Node of a Binary Tree in a Clone of That Tree||84.8%|Medium||
|1380|Lucky Numbers in a Matrix|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1380.Lucky-Numbers-in-a-Matrix)|70.8%|Easy||
|1381|Design a Stack With Increment Operation||76.1%|Medium||
|1382|Balance a Binary Search Tree||76.1%|Medium||
|1383|Maximum Performance of a Team||35.5%|Hard||
|1384|Total Sales Amount by Year||64.7%|Hard||
|1385|Find the Distance Value Between Two Arrays|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1385.Find-the-Distance-Value-Between-Two-Arrays)|66.5%|Easy||
|1386|Cinema Seat Allocation||35.7%|Medium||
|1387|Sort Integers by The Power Value||70.4%|Medium||
|1388|Pizza With 3n Slices||45.6%|Hard||
|1389|Create Target Array in the Given Order|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1389.Create-Target-Array-in-the-Given-Order)|84.8%|Easy||
|1390|Four Divisors||39.1%|Medium||
|1391|Check if There is a Valid Path in a Grid||44.9%|Medium||
|1392|Longest Happy Prefix||41.8%|Hard||
|1393|Capital Gain/Loss||90.5%|Medium||
|1394|Find Lucky Integer in an Array||63.2%|Easy||
|1395|Count Number of Teams||78.8%|Medium||
|1396|Design Underground System||69.5%|Medium||
|1397|Find All Good Strings||38.3%|Hard||
|1398|Customers Who Bought Products A and B but Not C||81.9%|Medium||
|1399|Count Largest Group||65.3%|Easy||
|1400|Construct K Palindrome Strings||62.8%|Medium||
|1401|Circle and Rectangle Overlapping||42.4%|Medium||
|1402|Reducing Dishes||72.3%|Hard||
|1403|Minimum Subsequence in Non-Increasing Order||71.1%|Easy||
|1404|Number of Steps to Reduce a Number in Binary Representation to One||49.9%|Medium||
|1405|Longest Happy String||52.3%|Medium||
|1406|Stone Game III||57.3%|Hard||
|1407|Top Travellers||83.7%|Easy||
|1408|String Matching in an Array||62.7%|Easy||
|1409|Queries on a Permutation With Key||81.5%|Medium||
|1410|HTML Entity Parser||54.4%|Medium||
|1411|Number of Ways to Paint N × 3 Grid||60.4%|Hard||
|1412|Find the Quiet Students in All Exams||65.1%|Hard||
|1413|Minimum Value to Get Positive Step by Step Sum||65.5%|Easy||
|1414|Find the Minimum Number of Fibonacci Numbers Whose Sum Is K||63.1%|Medium||
|1415|The k-th Lexicographical String of All Happy Strings of Length n||69.8%|Medium||
|1416|Restore The Array||36.3%|Hard||
|1417|Reformat The String||55.2%|Easy||
|1418|Display Table of Food Orders in a Restaurant||68.7%|Medium||
|1419|Minimum Number of Frogs Croaking||47.4%|Medium||
|1420|Build Array Where You Can Find The Maximum Exactly K Comparisons||64.4%|Hard||
|1421|NPV Queries||82.0%|Medium||
|1422|Maximum Score After Splitting a String||56.0%|Easy||
|1423|Maximum Points You Can Obtain from Cards||46.6%|Medium||
|1424|Diagonal Traverse II||45.8%|Medium||
|1425|Constrained Subsequence Sum||45.0%|Hard||
|1426|Counting Elements||59.1%|Easy||
|1427|Perform String Shifts||53.4%|Easy||
|1428|Leftmost Column with at Least a One||49.0%|Medium||
|1429|First Unique Number||49.5%|Medium||
|1430|Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree||45.1%|Medium||
|1431|Kids With the Greatest Number of Candies||88.4%|Easy||
|1432|Max Difference You Can Get From Changing an Integer||42.9%|Medium||
|1433|Check If a String Can Break Another String||67.2%|Medium||
|1434|Number of Ways to Wear Different Hats to Each Other||39.3%|Hard||
|1435|Create a Session Bar Chart||77.8%|Easy||
|1436|Destination City||77.0%|Easy||
|1437|Check If All 1's Are at Least Length K Places Away|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1437.Check-If-All-1's-Are-at-Least-Length-K-Places-Away)|62.8%|Easy||
|1438|Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit||44.2%|Medium||
|1439|Find the Kth Smallest Sum of a Matrix With Sorted Rows||60.0%|Hard||
|1440|Evaluate Boolean Expression||74.4%|Medium||
|1441|Build an Array With Stack Operations||69.7%|Easy||
|1442|Count Triplets That Can Form Two Arrays of Equal XOR||71.3%|Medium||
|1443|Minimum Time to Collect All Apples in a Tree||54.6%|Medium||
|1444|Number of Ways of Cutting a Pizza||53.7%|Hard||
|1445|Apples & Oranges||90.6%|Medium||
|1446|Consecutive Characters||61.0%|Easy||
|1447|Simplified Fractions||62.3%|Medium||
|1448|Count Good Nodes in Binary Tree||70.5%|Medium||
|1449|Form Largest Integer With Digits That Add up to Target||43.9%|Hard||
|1450|Number of Students Doing Homework at a Given Time||77.3%|Easy||
|1451|Rearrange Words in a Sentence||59.4%|Medium||
|1452|People Whose List of Favorite Companies Is Not a Subset of Another List||55.0%|Medium||
|1453|Maximum Number of Darts Inside of a Circular Dartboard||35.3%|Hard||
|1454|Active Users||38.8%|Medium||
|1455|Check If a Word Occurs As a Prefix of Any Word in a Sentence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1455.Check-If-a-Word-Occurs-As-a-Prefix-of-Any-Word-in-a-Sentence)|64.5%|Easy||
|1456|Maximum Number of Vowels in a Substring of Given Length||54.5%|Medium||
|1457|Pseudo-Palindromic Paths in a Binary Tree||70.9%|Medium||
|1458|Max Dot Product of Two Subsequences||43.1%|Hard||
|1459|Rectangles Area||64.6%|Medium||
|1460|Make Two Arrays Equal by Reversing Sub-arrays||72.2%|Easy||
|1461|Check If a String Contains All Binary Codes of Size K||47.3%|Medium||
|1462|Course Schedule IV||44.5%|Medium||
|1463|Cherry Pickup II||69.3%|Hard||
|1464|Maximum Product of Two Elements in an Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1464.Maximum-Product-of-Two-Elements-in-an-Array)|77.2%|Easy||
|1465|Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts||32.5%|Medium||
|1466|Reorder Routes to Make All Paths Lead to the City Zero||61.5%|Medium||
|1467|Probability of a Two Boxes Having The Same Number of Distinct Balls||61.0%|Hard||
|1468|Calculate Salaries||81.7%|Medium||
|1469|Find All The Lonely Nodes||80.6%|Easy||
|1470|Shuffle the Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1470.Shuffle-the-Array)|88.4%|Easy||
|1471|The k Strongest Values in an Array||58.4%|Medium||
|1472|Design Browser History||71.0%|Medium||
|1473|Paint House III||48.7%|Hard||
|1474|Delete N Nodes After M Nodes of a Linked List||74.5%|Easy||
|1475|Final Prices With a Special Discount in a Shop||75.2%|Easy||
|1476|Subrectangle Queries||88.4%|Medium||
|1477|Find Two Non-overlapping Sub-arrays Each With Target Sum||34.4%|Medium||
|1478|Allocate Mailboxes||54.0%|Hard||
|1479|Sales by Day of the Week||83.9%|Hard||
|1480|Running Sum of 1d Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1480.Running-Sum-of-1d-Array)|89.2%|Easy||
|1481|Least Number of Unique Integers after K Removals||55.7%|Medium||
|1482|Minimum Number of Days to Make m Bouquets||50.0%|Medium||
|1483|Kth Ancestor of a Tree Node||30.6%|Hard||
|1484|Group Sold Products By The Date||85.7%|Easy||
|1485|Clone Binary Tree With Random Pointer||79.4%|Medium||
|1486|XOR Operation in an Array||84.2%|Easy||
|1487|Making File Names Unique||30.6%|Medium||
|1488|Avoid Flood in The City||24.6%|Medium||
|1489|Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree||52.7%|Hard||
|1490|Clone N-ary Tree||83.4%|Medium||
|1491|Average Salary Excluding the Minimum and Maximum Salary||68.7%|Easy||
|1492|The kth Factor of n||63.4%|Medium||
|1493|Longest Subarray of 1's After Deleting One Element||58.0%|Medium||
|1494|Parallel Courses II||31.4%|Hard||
|1495|Friendly Movies Streamed Last Month||51.3%|Easy||
|1496|Path Crossing||55.3%|Easy||
|1497|Check If Array Pairs Are Divisible by k||40.5%|Medium||
|1498|Number of Subsequences That Satisfy the Given Sum Condition||38.3%|Medium||
|1499|Max Value of Equation||45.1%|Hard||
|1500|Design a File Sharing System||46.9%|Medium||
|1501|Countries You Can Safely Invest In||60.5%|Medium||
|1502|Can Make Arithmetic Progression From Sequence||71.2%|Easy||
|1503|Last Moment Before All Ants Fall Out of a Plank||53.1%|Medium||
|1504|Count Submatrices With All Ones||60.9%|Medium||
|1505|Minimum Possible Integer After at Most K Adjacent Swaps On Digits||36.3%|Hard||
|1506|Find Root of N-Ary Tree||80.4%|Medium||
|1507|Reformat Date||60.1%|Easy||
|1508|Range Sum of Sorted Subarray Sums||61.3%|Medium||
|1509|Minimum Difference Between Largest and Smallest Value in Three Moves||52.0%|Medium||
|1510|Stone Game IV||58.9%|Hard||
|1511|Customer Order Frequency||73.6%|Easy||
|1512|Number of Good Pairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1512.Number-of-Good-Pairs)|87.7%|Easy||
|1513|Number of Substrings With Only 1s||41.6%|Medium||
|1514|Path with Maximum Probability||39.9%|Medium||
|1515|Best Position for a Service Centre||37.1%|Hard||
|1516|Move Sub-Tree of N-Ary Tree||63.2%|Hard||
|1517|Find Users With Valid E-Mails||71.8%|Easy||
|1518|Water Bottles||60.6%|Easy||
|1519|Number of Nodes in the Sub-Tree With the Same Label||37.0%|Medium||
|1520|Maximum Number of Non-Overlapping Substrings||36.2%|Hard||
|1521|Find a Value of a Mysterious Function Closest to Target||44.3%|Hard||
|1522|Diameter of N-Ary Tree||69.4%|Medium||
|1523|Count Odd Numbers in an Interval Range||54.7%|Easy||
|1524|Number of Sub-arrays With Odd Sum||39.8%|Medium||
|1525|Number of Good Ways to Split a String||66.4%|Medium||
|1526|Minimum Number of Increments on Subarrays to Form a Target Array||60.0%|Hard||
|1527|Patients With a Condition||69.5%|Easy||
|1528|Shuffle String||85.8%|Easy||
|1529|Bulb Switcher IV||70.9%|Medium||
|1530|Number of Good Leaf Nodes Pairs||56.1%|Medium||
|1531|String Compression II||33.8%|Hard||
|1532|The Most Recent Three Orders||72.9%|Medium||
|1533|Find the Index of the Large Integer||54.7%|Medium||
|1534|Count Good Triplets||80.4%|Easy||
|1535|Find the Winner of an Array Game||47.4%|Medium||
|1536|Minimum Swaps to Arrange a Binary Grid||43.3%|Medium||
|1537|Get the Maximum Score||36.4%|Hard||
|1538|Guess the Majority in a Hidden Array||61.3%|Medium||
|1539|Kth Missing Positive Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1539.Kth-Missing-Positive-Number)|55.1%|Easy||
|1540|Can Convert String in K Moves||30.7%|Medium||
|1541|Minimum Insertions to Balance a Parentheses String||42.7%|Medium||
|1542|Find Longest Awesome Substring||36.6%|Hard||
|1543|Fix Product Name Format||67.6%|Easy||
|1544|Make The String Great||55.3%|Easy||
|1545|Find Kth Bit in Nth Binary String||57.4%|Medium||
|1546|Maximum Number of Non-Overlapping Subarrays With Sum Equals Target||43.7%|Medium||
|1547|Minimum Cost to Cut a Stick||52.4%|Hard||
|1548|The Most Similar Path in a Graph||54.2%|Hard||
|1549|The Most Recent Orders for Each Product||66.7%|Medium||
|1550|Three Consecutive Odds||64.9%|Easy||
|1551|Minimum Operations to Make Array Equal||77.9%|Medium||
|1552|Magnetic Force Between Two Balls||48.8%|Medium||
|1553|Minimum Number of Days to Eat N Oranges||29.5%|Hard||
|1554|Strings Differ by One Character||63.4%|Medium||
|1555|Bank Account Summary||52.7%|Medium||
|1556|Thousand Separator||57.8%|Easy||
|1557|Minimum Number of Vertices to Reach All Nodes||75.5%|Medium||
|1558|Minimum Numbers of Function Calls to Make Target Array||63.5%|Medium||
|1559|Detect Cycles in 2D Grid||44.3%|Hard||
|1560|Most Visited Sector in  a Circular Track||56.9%|Easy||
|1561|Maximum Number of Coins You Can Get||77.7%|Medium||
|1562|Find Latest Group of Size M||39.6%|Medium||
|1563|Stone Game V||40.1%|Hard||
|1564|Put Boxes Into the Warehouse I||65.7%|Medium||
|1565|Unique Orders and Customers Per Month||82.7%|Easy||
|1566|Detect Pattern of Length M Repeated K or More Times||42.5%|Easy||
|1567|Maximum Length of Subarray With Positive Product||36.7%|Medium||
|1568|Minimum Number of Days to Disconnect Island||50.3%|Hard||
|1569|Number of Ways to Reorder Array to Get Same BST||50.0%|Hard||
|1570|Dot Product of Two Sparse Vectors||91.2%|Medium||
|1571|Warehouse Manager||89.9%|Easy||
|1572|Matrix Diagonal Sum||78.3%|Easy||
|1573|Number of Ways to Split a String|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1573.Number-of-Ways-to-Split-a-String)|31.0%|Medium||
|1574|Shortest Subarray to be Removed to Make Array Sorted||33.3%|Medium||
|1575|Count All Possible Routes||57.3%|Hard||
|1576|Replace All ?'s to Avoid Consecutive Repeating Characters||48.3%|Easy||
|1577|Number of Ways Where Square of Number Is Equal to Product of Two Numbers||37.4%|Medium||
|1578|Minimum Deletion Cost to Avoid Repeating Letters||60.5%|Medium||
|1579|Remove Max Number of Edges to Keep Graph Fully Traversable||46.2%|Hard||
|1580|Put Boxes Into the Warehouse II||61.8%|Medium||
|1581|Customer Who Visited but Did Not Make Any Transactions||89.8%|Easy||
|1582|Special Positions in a Binary Matrix||64.3%|Easy||
|1583|Count Unhappy Friends||54.3%|Medium||
|1584|Min Cost to Connect All Points||51.6%|Medium||
|1585|Check If String Is Transformable With Substring Sort Operations||48.2%|Hard||
|1586|Binary Search Tree Iterator II||66.5%|Medium||
|1587|Bank Account Summary II||89.9%|Easy||
|1588|Sum of All Odd Length Subarrays||82.0%|Easy||
|1589|Maximum Sum Obtained of Any Permutation||34.8%|Medium||
|1590|Make Sum Divisible by P||27.1%|Medium||
|1591|Strange Printer II||55.1%|Hard||
|1592|Rearrange Spaces Between Words||43.8%|Easy||
|1593|Split a String Into the Max Number of Unique Substrings||48.9%|Medium||
|1594|Maximum Non Negative Product in a Matrix||32.0%|Medium||
|1595|Minimum Cost to Connect Two Groups of Points||42.9%|Hard||
|1596|The Most Frequently Ordered Products for Each Customer||83.9%|Medium||
|1597|Build Binary Expression Tree From Infix Expression||63.1%|Hard||
|1598|Crawler Log Folder||64.1%|Easy||
|1599|Maximum Profit of Operating a Centennial Wheel||43.6%|Medium||
|1600|Throne Inheritance||60.2%|Medium||
|1601|Maximum Number of Achievable Transfer Requests||47.8%|Hard||
|1602|Find Nearest Right Node in Binary Tree||73.9%|Medium||
|1603|Design Parking System||86.5%|Easy||
|1604|Alert Using Same Key-Card Three or More Times in a One Hour Period||42.3%|Medium||
|1605|Find Valid Matrix Given Row and Column Sums||77.2%|Medium||
|1606|Find Servers That Handled Most Number of Requests||36.9%|Hard||
|1607|Sellers With No Sales||55.5%|Easy||
|1608|Special Array With X Elements Greater Than or Equal X||61.7%|Easy||
|1609|Even Odd Tree||53.0%|Medium||
|1610|Maximum Number of Visible Points||29.0%|Hard||
|1611|Minimum One Bit Operations to Make Integers Zero||57.5%|Hard||
|1612|Check If Two Expression Trees are Equivalent||70.0%|Medium||
|1613|Find the Missing IDs||73.4%|Medium||
|1614|Maximum Nesting Depth of the Parentheses||83.2%|Easy||
|1615|Maximal Network Rank||51.9%|Medium||
|1616|Split Two Strings to Make Palindrome||36.2%|Medium||
|1617|Count Subtrees With Max Distance Between Cities||63.7%|Hard||
|1618|Maximum Font to Fit a Sentence in a Screen||57.3%|Medium||
|1619|Mean of Array After Removing Some Elements||65.1%|Easy||
|1620|Coordinate With Maximum Network Quality||37.1%|Medium||
|1621|Number of Sets of K Non-Overlapping Line Segments||41.0%|Medium||
|1622|Fancy Sequence||15.3%|Hard||
|1623|All Valid Triplets That Can Represent a Country||88.9%|Easy||
|1624|Largest Substring Between Two Equal Characters||59.0%|Easy||
|1625|Lexicographically Smallest String After Applying Operations||63.7%|Medium||
|1626|Best Team With No Conflicts||37.7%|Medium||
|1627|Graph Connectivity With Threshold||38.5%|Hard||
|1628|Design an Expression Tree With Evaluate Function||80.2%|Medium||
|1629|Slowest Key||59.2%|Easy||
|1630|Arithmetic Subarrays||77.7%|Medium||
|1631|Path With Minimum Effort||50.1%|Medium||
|1632|Rank Transform of a Matrix||30.7%|Hard||
|1633|Percentage of Users Attended a Contest||72.1%|Easy||
|1634|Add Two Polynomials Represented as Linked Lists||55.7%|Medium||
|1635|Hopper Company Queries I||56.4%|Hard||
|1636|Sort Array by Increasing Frequency||66.8%|Easy||
|1637|Widest Vertical Area Between Two Points Containing No Points||83.9%|Medium||
|1638|Count Substrings That Differ by One Character||68.8%|Medium||
|1639|Number of Ways to Form a Target String Given a Dictionary||39.5%|Hard||
|1640|Check Array Formation Through Concatenation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1640.Check-Array-Formation-Through-Concatenation)|60.2%|Easy||
|1641|Count Sorted Vowel Strings|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1641.Count-Sorted-Vowel-Strings)|77.0%|Medium||
|1642|Furthest Building You Can Reach||50.6%|Medium||
|1643|Kth Smallest Instructions||43.4%|Hard||
|1644|Lowest Common Ancestor of a Binary Tree II||57.2%|Medium||
|1645|Hopper Company Queries II||39.7%|Hard||
|1646|Get Maximum in Generated Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1646.Get-Maximum-in-Generated-Array)|53.5%|Easy||
|1647|Minimum Deletions to Make Character Frequencies Unique|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1647.Minimum-Deletions-to-Make-Character-Frequencies-Unique)|54.4%|Medium||
|1648|Sell Diminishing-Valued Colored Balls|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1648.Sell-Diminishing-Valued-Colored-Balls)|30.8%|Medium||
|1649|Create Sorted Array through Instructions|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1649.Create-Sorted-Array-through-Instructions)|36.3%|Hard||
|1650|Lowest Common Ancestor of a Binary Tree III||77.0%|Medium||
|1651|Hopper Company Queries III||67.4%|Hard||
|1652|Defuse the Bomb|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1652.Defuse-the-Bomb)|63.0%|Easy||
|1653|Minimum Deletions to Make String Balanced|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1653.Minimum-Deletions-to-Make-String-Balanced)|50.6%|Medium||
|1654|Minimum Jumps to Reach Home|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1654.Minimum-Jumps-to-Reach-Home)|26.2%|Medium||
|1655|Distribute Repeating Integers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1655.Distribute-Repeating-Integers)|40.4%|Hard||
|1656|Design an Ordered Stream|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1656.Design-an-Ordered-Stream)|82.2%|Easy||
|1657|Determine if Two Strings Are Close|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1657.Determine-if-Two-Strings-Are-Close)|55.4%|Medium||
|1658|Minimum Operations to Reduce X to Zero|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1658.Minimum-Operations-to-Reduce-X-to-Zero)|33.4%|Medium||
|1659|Maximize Grid Happiness|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1659.Maximize-Grid-Happiness)|35.4%|Hard||
|1660|Correct a Binary Tree||78.7%|Medium||
|1661|Average Time of Process per Machine||79.3%|Easy||
|1662|Check If Two String Arrays are Equivalent|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1662.Check-If-Two-String-Arrays-are-Equivalent)|83.3%|Easy||
|1663|Smallest String With A Given Numeric Value|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1663.Smallest-String-With-A-Given-Numeric-Value)|64.3%|Medium||
|1664|Ways to Make a Fair Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1664.Ways-to-Make-a-Fair-Array)|60.6%|Medium||
|1665|Minimum Initial Energy to Finish Tasks|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1665.Minimum-Initial-Energy-to-Finish-Tasks)|66.6%|Hard||
|1666|Change the Root of a Binary Tree||69.5%|Medium||
|1667|Fix Names in a Table||63.5%|Easy||
|1668|Maximum Repeating Substring|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1668.Maximum-Repeating-Substring)|38.5%|Easy||
|1669|Merge In Between Linked Lists|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1669.Merge-In-Between-Linked-Lists)|77.2%|Medium||
|1670|Design Front Middle Back Queue|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1670.Design-Front-Middle-Back-Queue)|54.5%|Medium||
|1671|Minimum Number of Removals to Make Mountain Array||45.6%|Hard||
|1672|Richest Customer Wealth|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1672.Richest-Customer-Wealth)|88.4%|Easy||
|1673|Find the Most Competitive Subsequence|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1673.Find-the-Most-Competitive-Subsequence)|45.2%|Medium||
|1674|Minimum Moves to Make Array Complementary|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1674.Minimum-Moves-to-Make-Array-Complementary)|34.7%|Medium||
|1675|Minimize Deviation in Array||48.6%|Hard||
|1676|Lowest Common Ancestor of a Binary Tree IV||78.5%|Medium||
|1677|Product's Worth Over Invoices||73.9%|Easy||
|1678|Goal Parser Interpretation|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1678.Goal-Parser-Interpretation)|85.7%|Easy||
|1679|Max Number of K-Sum Pairs|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1679.Max-Number-of-K-Sum-Pairs)|54.2%|Medium||
|1680|Concatenation of Consecutive Binary Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1680.Concatenation-of-Consecutive-Binary-Numbers)|52.4%|Medium||
|1681|Minimum Incompatibility|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1681.Minimum-Incompatibility)|35.4%|Hard||
|1682|Longest Palindromic Subsequence II||50.8%|Medium||
|1683|Invalid Tweets||91.2%|Easy||
|1684|Count the Number of Consistent Strings|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1684.Count-the-Number-of-Consistent-Strings)|83.4%|Easy||
|1685|Sum of Absolute Differences in a Sorted Array|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1685.Sum-of-Absolute-Differences-in-a-Sorted-Array)|61.9%|Medium||
|1686|Stone Game VI||49.0%|Medium||
|1687|Delivering Boxes from Storage to Ports||34.6%|Hard||
|1688|Count of Matches in Tournament|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1688.Count-of-Matches-in-Tournament)|82.6%|Easy||
|1689|Partitioning Into Minimum Number Of Deci-Binary Numbers|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1689.Partitioning-Into-Minimum-Number-Of-Deci-Binary-Numbers)|87.4%|Medium||
|1690|Stone Game VII|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1690.Stone-Game-VII)|47.5%|Medium||
|1691|Maximum Height by Stacking Cuboids ||49.9%|Hard||
|1692|Count Ways to Distribute Candies||62.1%|Hard||
|1693|Daily Leads and Partners||90.9%|Easy||
|1694|Reformat Phone Number|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1694.Reformat-Phone-Number)|66.2%|Easy||
|1695|Maximum Erasure Value|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1695.Maximum-Erasure-Value)|49.6%|Medium||
|1696|Jump Game VI|[Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1696.Jump-Game-VI)|54.0%|Medium||
|1697|Checking Existence of Edge Length Limited Paths||55.7%|Hard||
|1698|Number of Distinct Substrings in a String||57.4%|Medium||
|1699|Number of Calls Between Two Persons||86.7%|Medium||
|1700|Number of Students Unable to Eat Lunch||69.4%|Easy||
|1701|Average Waiting Time||61.4%|Medium||
|1702|Maximum Binary String After Change||60.4%|Medium||
|1703|Minimum Adjacent Swaps for K Consecutive Ones||40.1%|Hard||
|1704|Determine if String Halves Are Alike||77.7%|Easy||
|1705|Maximum Number of Eaten Apples||42.1%|Medium||
|1706|Where Will the Ball Fall||59.1%|Medium||
|1707|Maximum XOR With an Element From Array||47.2%|Hard||
|1708|Largest Subarray Length K||63.9%|Easy||
|1709|Biggest Window Between Visits||82.4%|Medium||
|1710|Maximum Units on a Truck||70.6%|Easy||
|1711|Count Good Meals||26.0%|Medium||
|1712|Ways to Split Array Into Three Subarrays||29.8%|Medium||
|1713|Minimum Operations to Make a Subsequence||45.8%|Hard||
|1714|Sum Of Special Evenly-Spaced Elements In Array||47.9%|Hard||
|1715|Count Apples and Oranges||79.3%|Medium||
|1716|Calculate Money in Leetcode Bank||67.5%|Easy||
|1717|Maximum Score From Removing Substrings||39.6%|Medium||
|1718|Construct the Lexicographically Largest Valid Sequence||44.4%|Medium||
|1719|Number Of Ways To Reconstruct A Tree||39.4%|Hard||
|1720|Decode XORed Array||85.8%|Easy||
|1721|Swapping Nodes in a Linked List||64.8%|Medium||
|1722|Minimize Hamming Distance After Swap Operations||55.3%|Medium||
|1723|Find Minimum Time to Finish All Jobs||44.4%|Hard||
|1724|Checking Existence of Edge Length Limited Paths II||58.6%|Hard||
|1725|Number Of Rectangles That Can Form The Largest Square||77.8%|Easy||
|1726|Tuple with Same Product||55.3%|Medium||
|1727|Largest Submatrix With Rearrangements||58.4%|Medium||
|1728|Cat and Mouse II||41.0%|Hard||
|1729|Find Followers Count||72.2%|Easy||
|1730|Shortest Path to Get Food||59.2%|Medium||
|1731|The Number of Employees Which Report to Each Employee||49.6%|Easy||
|1732|Find the Highest Altitude||81.8%|Easy||
|1733|Minimum Number of People to Teach||36.6%|Medium||
|1734|Decode XORed Permutation||51.1%|Medium||
|1735|Count Ways to Make Array With Product||49.1%|Hard||
|1736|Latest Time by Replacing Hidden Digits||41.3%|Easy||
|1737|Change Minimum Characters to Satisfy One of Three Conditions||29.0%|Medium||
|1738|Find Kth Largest XOR Coordinate Value||62.2%|Medium||
|1739|Building Boxes||49.5%|Hard||
|1740|Find Distance in a Binary Tree||70.7%|Medium||
|1741|Find Total Time Spent by Each Employee||90.9%|Easy||
|1742|Maximum Number of Balls in a Box||75.0%|Easy||
|1743|Restore the Array From Adjacent Pairs||62.0%|Medium||
|1744|Can You Eat Your Favorite Candy on Your Favorite Day?||30.3%|Medium||
|1745|Palindrome Partitioning IV||48.8%|Hard||
|1746|Maximum Subarray Sum After One Operation||62.3%|Medium||
|1747|Leetflex Banned Accounts||70.9%|Medium||
|1748|Sum of Unique Elements||77.8%|Easy||
|1749|Maximum Absolute Sum of Any Subarray||49.5%|Medium||
|1750|Minimum Length of String After Deleting Similar Ends||41.2%|Medium||
|1751|Maximum Number of Events That Can Be Attended II||47.0%|Hard||
|1752|Check if Array Is Sorted and Rotated||72.7%|Easy||
|1753|Maximum Score From Removing Stones||60.3%|Medium||
|1754|Largest Merge Of Two Strings||38.0%|Medium||
|1755|Closest Subsequence Sum||34.5%|Hard||
|1756|Design Most Recently Used Queue||77.3%|Medium||
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|

------------------------------------------------------------------

下面这些是免费的算法题，但是暂时还不能使用 Go 解答的：

暂无

------------------------------------------------------------------


## 三.分类

## Array


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Array/)



## String

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/String/)


## Two Pointers

![](./topic/Two_pointers.png)

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


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Two_Pointers/)


## Linked List

![](./topic/Linked_List.png)


- 巧妙的构造虚拟头结点。可以使遍历处理逻辑更加统一。
- 灵活使用递归。构造递归条件，使用递归可以巧妙的解题。不过需要注意有些题目不能使用递归，因为递归深度太深会导致超时和栈溢出。
- 链表区间逆序。第 92 题。
- 链表寻找中间节点。第 876 题。链表寻找倒数第 n 个节点。第 19 题。只需要一次遍历就可以得到答案。
- 合并 K 个有序链表。第 21 题，第 23 题。
- 链表归类。第 86 题，第 328 题。
- 链表排序，时间复杂度要求 O(n * log n)，空间复杂度 O(1)。只有一种做法，归并排序，至顶向下归并。第 148 题。
- 判断链表是否存在环，如果有环，输出环的交叉点的下标；判断 2 个链表是否有交叉点，如果有交叉点，输出交叉点。第 141 题，第 142 题，第 160 题。


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/)




## Stack

![](./topic/Stack.png)

- 括号匹配问题及类似问题。第 20 题，第 921 题，第 1021 题。
- 栈的基本 pop 和 push 操作。第 71 题，第 150 题，第 155 题，第 224 题，第 225 题，第 232 题，第 946 题，第 1047 题。
- 利用栈进行编码问题。第 394 题，第 682 题，第 856 题，第 880 题。
- **单调栈**。**利用栈维护一个单调递增或者递减的下标数组**。第 84 题，第 456 题，第 496 题，第 503 题，第 739 题，第 901 题，第 907 题，第 1019 题。

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Stack/)



## Tree


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Tree/)





## Dynamic Programming

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Dynamic_Programming/)



## Backtracking

![](./topic/Backtracking.png)

- 排列问题 Permutations。第 46 题，第 47 题。第 60 题，第 526 题，第 996 题。
- 组合问题 Combination。第 39 题，第 40 题，第 77 题，第 216 题。
- 排列和组合杂交问题。第 1079 题。
- N 皇后终极解法(二进制解法)。第 51 题，第 52 题。
- 数独问题。第 37 题。
- 四个方向搜索。第 79 题，第 212 题，第 980 题。
- 子集合问题。第 78 题，第 90 题。
- Trie。第 208 题，第 211 题。
- BFS 优化。第 126 题，第 127 题。
- DFS 模板。(只是一个例子，不对应任何题)

```go
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum2(candidates, target, 0, c, &res)
	return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // 这里是去重的关键逻辑
			continue
		}
		if target >= nums[i] {
			c = append(c, nums[i])
			findcombinationSum2(nums, target-nums[i], i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
```
- BFS 模板。(只是一个例子，不对应任何题)

```go
func updateMatrix_BFS(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	queue := make([][]int, 0)
	for i, _ := range matrix {
		res[i] = make([]int, len(matrix[0]))
		for j, _ := range res[i] {
			if matrix[i][j] == 0 {
				res[i][j] = -1
				queue = append(queue, []int{i, j})
			}
		}
	}
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size -= 1
			node := queue[0]
			queue = queue[1:]
			i, j := node[0], node[1]
			for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				x := i + direction[0]
				y := j + direction[1]
				if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || res[x][y] < 0 || res[x][y] > 0 {
					continue
				}
				res[x][y] = level
				queue = append(queue, []int{x, y})
			}
		}
		level++
	}
	for i, row := range res {
		for j, cell := range row {
			if cell == -1 {
				res[i][j] = 0
			}
		}
	}
	return res
}
```


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Backtracking/)


## Depth First Search


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Depth_First_Search/)




## Breadth First Search



Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Breadth_First_Search/)





## Binary Search

![]()

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

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Binary_Search/)



## Math


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Math/)




## Hash Table


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Hash_Table/)



## Sort

![](./topic/Sort.png)

- 深刻的理解多路快排。第 75 题。
- 链表的排序，插入排序(第 147 题)和归并排序(第 148 题)
- 桶排序和基数排序。第 164 题。
- "摆动排序"。第 324 题。
- 两两不相邻的排序。第 767 题，第 1054 题。
- "饼子排序"。第 969 题。

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Sort/)


## Bit Manipulation

![](./topic/Bit_Manipulation.png)

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

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Bit_Manipulation/)


## Union Find

![](./topic/Union_Find.png)

- 灵活使用并查集的思想，熟练掌握并查集的[模板](https://github.com/halfrost/LeetCode-Go/blob/master/template/UnionFind.go)，模板中有两种并查集的实现方式，一种是路径压缩 + 秩优化的版本，另外一种是计算每个集合中元素的个数 + 最大集合元素个数的版本，这两种版本都有各自使用的地方。能使用第一类并查集模板的题目有：第 128 题，第 130 题，第 547 题，第 684 题，第 721 题，第 765 题，第 778 题，第 839 题，第 924 题，第 928 题，第 947 题，第 952 题，第 959 题，第 990 题。能使用第二类并查集模板的题目有：第 803 题，第 952 题。第 803 题秩优化和统计集合个数这些地方会卡时间，如果不优化，会 TLE。
- 并查集是一种思想，有些题需要灵活使用这种思想，而不是死套模板，如第 399 题，这一题是 stringUnionFind，利用并查集思想实现的。这里每个节点是基于字符串和 map 的，而不是单纯的用 int 节点编号实现的。
- 有些题死套模板反而做不出来，比如第 685 题，这一题不能路径压缩和秩优化，因为题目中涉及到有向图，需要知道节点的前驱节点，如果路径压缩了，这一题就没法做了。这一题不需要路径压缩和秩优化。
- 灵活的抽象题目给的信息，将给定的信息合理的编号，使用并查集解题，并用 map 降低时间复杂度，如第 721 题，第 959 题。
- 关于地图，砖块，网格的题目，可以新建一个特殊节点，将四周边缘的砖块或者网格都 union() 到这个特殊节点上。第 130 题，第 803 题。
- 能用并查集的题目，一般也可以用 DFS 和 BFS 解答，只不过时间复杂度会高一点。


Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/)



## Sliding Window

![](./topic/Sliding_Window.png)

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

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Sliding_Window/)


## Segment Tree

![](./topic/Segment_Tree.png)

- 线段树的经典数组实现写法。将合并两个节点 pushUp 逻辑抽象出来了，可以实现任意操作(常见的操作有：加法，取 max，min 等等)。第 218 题，第 303 题，第 307 题，第 699 题。
- 计数线段树的经典写法。第 315 题，第 327 题，第 493 题。
- 线段树的树的实现写法。第 715 题，第 732 题。
- 区间懒惰更新。第 218 题，第 699 题。
- 离散化。离散化需要注意一个特殊情况：假如三个区间为 [1,10] [1,4] [6,10]，离散化后 x[1]=1,x[2]=4,x[3]=6,x[4]=10。第一个区间为 [1,4]，第二个区间为 [1,2]，第三个区间为 [3,4]，这样一来，区间一 = 区间二 + 区间三，这和离散前的模型不符，离散前，很明显，区间一 > 区间二 + 区间三。正确的做法是：在相差大于 1 的数间加一个数，例如在上面 1 4 6 10 中间加 5，即可 x[1]=1,x[2]=4,x[3]=5,x[4]=6,x[5]=10。这样处理之后，区间一是 1-5 ，区间二是 1-2 ，区间三是 4-5 。
- 灵活构建线段树。线段树节点可以存储多条信息，合并两个节点的 pushUp 操作也可以是多样的。第 850 题，第 1157 题。


线段树[题型](https://blog.csdn.net/xuechelingxiao/article/details/38313105)从简单到困难:

1. 单点更新:  
	[HDU 1166 敌兵布阵](http://acm.hdu.edu.cn/showproblem.php?pid=1166) update:单点增减 query:区间求和  
	[HDU 1754 I Hate It](http://acm.hdu.edu.cn/showproblem.php?pid=1754) update:单点替换 query:区间最值  
	[HDU 1394 Minimum Inversion Number](http://acm.hdu.edu.cn/showproblem.php?pid=1394) update:单点增减 query:区间求和  
	[HDU 2795 Billboard](http://acm.hdu.edu.cn/showproblem.php?pid=2795) query:区间求最大值的位子(直接把update的操作在query里做了)
2. 区间更新:  
	[HDU 1698 Just a Hook](http://acm.hdu.edu.cn/showproblem.php?pid=1698) update:成段替换 (由于只query一次总区间,所以可以直接输出 1 结点的信息)  
	[POJ 3468 A Simple Problem with Integers](http://poj.org/problem?id=3468) update:成段增减 query:区间求和  
	[POJ 2528 Mayor’s posters](http://poj.org/problem?id=2528) 离散化 + update:成段替换 query:简单hash  
	[POJ 3225 Help with Intervals](http://poj.org/problem?id=3225) update:成段替换,区间异或 query:简单hash
3. 区间合并(这类题目会询问区间中满足条件的连续最长区间,所以PushUp的时候需要对左右儿子的区间进行合并):  
	[POJ 3667 Hotel](http://poj.org/problem?id=3667) update:区间替换 query:询问满足条件的最左端点
4. 扫描线(这类题目需要将一些操作排序,然后从左到右用一根扫描线扫过去最典型的就是矩形面积并,周长并等题):  
	[HDU 1542 Atlantis](http://acm.hdu.edu.cn/showproblem.php?pid=1542) update:区间增减 query:直接取根节点的值  
	[HDU 1828 Picture](http://acm.hdu.edu.cn/showproblem.php?pid=1828) update:区间增减 query:直接取根节点的值

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Segment_Tree/)


## Binary Indexed Tree

![](./topic/Binary_Indexed_Tree.png)

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Binary_Indexed_Tree/)


----------------------------------------------------------------------------------------

<p align='center'>
<a href="https://github.com/halfrost/LeetCode-Go/releases/tag/Special"><img src="https://img.halfrost.com/Leetcode/ACM-ICPC_Algorithm_Template.png"></a>
</p>

Thank you for reading here. This is bonus. You can download my [《ACM-ICPC Algorithm Template》](https://github.com/halfrost/LeetCode-Go/releases/tag/Special/)



## ♥️ Thanks

Thanks for your Star！

[![Stargazers over time](https://starchart.cc/halfrost/LeetCode-Go.svg)](https://starchart.cc/halfrost/LeetCode-Go)

