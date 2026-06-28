---
title: 2.16 ✅ Union Find
type: docs
weight: 16
---

# Union Find

![](https://img.halfrost.com/Leetcode/Union_Find.png)

- Flexibly use the idea of union find, and be proficient with the union find [template]({{< relref "/ChapterThree/UnionFind.md" >}}). There are two implementations of union find in the template: one is the path compression + union by rank version, and the other is the version that calculates the number of elements in each set + the number of elements in the largest set. Both versions have their own use cases. Problems that can use the first type of union find template include: Problem 128, Problem 130, Problem 547, Problem 684, Problem 721, Problem 765, Problem 778, Problem 839, Problem 924, Problem 928, Problem 947, Problem 952, Problem 959, Problem 990. Problems that can use the second type of union find template include: Problem 803, Problem 952. In Problem 803, union by rank and counting the number of elements in a set may cause time limits; without optimization, it will TLE.
- Union find is an idea. Some problems require flexible use of this idea rather than rigidly applying a template, such as Problem 399. This problem is stringUnionFind, implemented using the union find idea. Here each node is based on strings and maps, rather than simply being implemented with int node numbers.
- For some problems, rigidly applying the template may make them unsolvable, such as Problem 685. This problem cannot use path compression and union by rank, because the problem involves a directed graph and needs to know each node's predecessor. If path compression is used, this problem cannot be solved. This problem does not require path compression or union by rank.
- Flexibly abstract the information given by the problem, reasonably number the given information, use union find to solve the problem, and use maps to reduce time complexity, such as Problem 721 and Problem 959.
- For problems about maps, bricks, and grids, you can create a special node and union() the bricks or grids around the edges with this special node. Problem 130, Problem 803.
- Problems that can be solved with union find can generally also be solved with DFS and BFS, but the time complexity will be a bit higher.



| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0128|Longest Consecutive Sequence|[Go]({{< relref "/ChapterFour/0100~0199/0128.Longest-Consecutive-Sequence.md" >}})|Medium| O(n)| O(n)|❤️|48.5%|
|0130|Surrounded Regions|[Go]({{< relref "/ChapterFour/0100~0199/0130.Surrounded-Regions.md" >}})|Medium| O(m\*n)| O(m\*n)||36.7%|
|0200|Number of Islands|[Go]({{< relref "/ChapterFour/0200~0299/0200.Number-of-Islands.md" >}})|Medium| O(m\*n)| O(m\*n)||57.0%|
|0399|Evaluate Division|[Go]({{< relref "/ChapterFour/0300~0399/0399.Evaluate-Division.md" >}})|Medium| O(n)| O(n)||59.6%|
|0547|Number of Provinces|[Go]({{< relref "/ChapterFour/0500~0599/0547.Number-of-Provinces.md" >}})|Medium| O(n^2)| O(n)||63.7%|
|0684|Redundant Connection|[Go]({{< relref "/ChapterFour/0600~0699/0684.Redundant-Connection.md" >}})|Medium| O(n)| O(n)||62.2%|
|0685|Redundant Connection II|[Go]({{< relref "/ChapterFour/0600~0699/0685.Redundant-Connection-II.md" >}})|Hard| O(n)| O(n)||34.1%|
|0695|Max Area of Island|[Go]({{< relref "/ChapterFour/0600~0699/0695.Max-Area-of-Island.md" >}})|Medium||||71.8%|
|0721|Accounts Merge|[Go]({{< relref "/ChapterFour/0700~0799/0721.Accounts-Merge.md" >}})|Medium| O(n)| O(n)|❤️|56.3%|
|0765|Couples Holding Hands|[Go]({{< relref "/ChapterFour/0700~0799/0765.Couples-Holding-Hands.md" >}})|Hard| O(n)| O(n)|❤️|56.6%|
|0778|Swim in Rising Water|[Go]({{< relref "/ChapterFour/0700~0799/0778.Swim-in-Rising-Water.md" >}})|Hard| O(n^2)| O(n)|❤️|59.8%|
|0785|Is Graph Bipartite?|[Go]({{< relref "/ChapterFour/0700~0799/0785.Is-Graph-Bipartite.md" >}})|Medium||||53.1%|
|0803|Bricks Falling When Hit|[Go]({{< relref "/ChapterFour/0800~0899/0803.Bricks-Falling-When-Hit.md" >}})|Hard| O(n^2)| O(n)|❤️|34.4%|
|0839|Similar String Groups|[Go]({{< relref "/ChapterFour/0800~0899/0839.Similar-String-Groups.md" >}})|Hard| O(n^2)| O(n)||48.0%|
|0924|Minimize Malware Spread|[Go]({{< relref "/ChapterFour/0900~0999/0924.Minimize-Malware-Spread.md" >}})|Hard| O(m\*n)| O(n)||42.1%|
|0928|Minimize Malware Spread II|[Go]({{< relref "/ChapterFour/0900~0999/0928.Minimize-Malware-Spread-II.md" >}})|Hard| O(m\*n)| O(n)|❤️|42.7%|
|0947|Most Stones Removed with Same Row or Column|[Go]({{< relref "/ChapterFour/0900~0999/0947.Most-Stones-Removed-with-Same-Row-or-Column.md" >}})|Medium| O(n)| O(n)||58.9%|
|0952|Largest Component Size by Common Factor|[Go]({{< relref "/ChapterFour/0900~0999/0952.Largest-Component-Size-by-Common-Factor.md" >}})|Hard| O(n)| O(n)|❤️|40.0%|
|0959|Regions Cut By Slashes|[Go]({{< relref "/ChapterFour/0900~0999/0959.Regions-Cut-By-Slashes.md" >}})|Medium| O(n^2)| O(n^2)|❤️|69.1%|
|0990|Satisfiability of Equality Equations|[Go]({{< relref "/ChapterFour/0900~0999/0990.Satisfiability-of-Equality-Equations.md" >}})|Medium| O(n)| O(n)||50.5%|
|1020|Number of Enclaves|[Go]({{< relref "/ChapterFour/1000~1099/1020.Number-of-Enclaves.md" >}})|Medium||||65.5%|
|1202|Smallest String With Swaps|[Go]({{< relref "/ChapterFour/1200~1299/1202.Smallest-String-With-Swaps.md" >}})|Medium||||57.7%|
|1254|Number of Closed Islands|[Go]({{< relref "/ChapterFour/1200~1299/1254.Number-of-Closed-Islands.md" >}})|Medium||||64.1%|
|1319|Number of Operations to Make Network Connected|[Go]({{< relref "/ChapterFour/1300~1399/1319.Number-of-Operations-to-Make-Network-Connected.md" >}})|Medium||||62.1%|
|1579|Remove Max Number of Edges to Keep Graph Fully Traversable|[Go]({{< relref "/ChapterFour/1500~1599/1579.Remove-Max-Number-of-Edges-to-Keep-Graph-Fully-Traversable.md" >}})|Hard||||53.2%|
|1631|Path With Minimum Effort|[Go]({{< relref "/ChapterFour/1600~1699/1631.Path-With-Minimum-Effort.md" >}})|Medium||||55.7%|
