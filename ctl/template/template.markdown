
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
<a href="https://codecov.io/gh/halfrost/LeetCode-Go"><img src="https://codecov.io/gh/halfrost/LeetCode-Go/branch/master/graph/badge.svg" /></a>
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

{{.PersonalData}}

## 二. 目录

{{.TotalNum}}

{{.AvailableTable}}

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

