# [911. Online Election](https://leetcode.com/problems/online-election/)


## 题目:

In an election, the `i`-th vote was cast for `persons[i]` at time `times[i]`.

Now, we would like to implement the following query function: `TopVotedCandidate.q(int t)` will return the number of the person that was leading the election at time `t`.

Votes cast at time `t` will count towards our query. In the case of a tie, the most recent vote (among tied candidates) wins.

**Example 1:**

    Input: ["TopVotedCandidate","q","q","q","q","q","q"], [[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]
    Output: [null,0,1,1,0,0,1]
    Explanation: 
    At time 3, the votes are [0], and 0 is leading.
    At time 12, the votes are [0,1,1], and 1 is leading.
    At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
    This continues for 3 more queries at time 15, 24, and 8.

**Note:**

1. `1 <= persons.length = times.length <= 5000`
2. `0 <= persons[i] <= persons.length`
3. `times` is a strictly increasing array with all elements in `[0, 10^9]`.
4. `TopVotedCandidate.q` is called at most `10000` times per test case.
5. `TopVotedCandidate.q(int t)` is always called with `t >= times[0]`.


## 题目大意

在选举中，第 i 张票是在时间为 times[i] 时投给 persons[i] 的。

现在，我们想要实现下面的查询函数： TopVotedCandidate.q(int t) 将返回在 t 时刻主导选举的候选人的编号。

在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

提示：

1. 1 <= persons.length = times.length <= 5000
2. 0 <= persons[i] <= persons.length
3. times 是严格递增的数组，所有元素都在 [0, 10^9] 范围中。
4. 每个测试用例最多调用 10000 次 TopVotedCandidate.q。
5. TopVotedCandidate.q(int t) 被调用时总是满足 t >= times[0]。




## 解题思路

- 给出一个 2 个数组，分别代表第 `i` 人在第 `t` 时刻获得的票数。需要实现一个查询功能的函数，查询在任意 `t` 时刻，输出谁的选票领先。
- `persons[]` 数组里面装的是获得选票人的编号，`times[]` 数组里面对应的是每个选票的时刻。`times[]` 数组默认是有序的，从小到大排列。先计算出每个时刻哪个人选票领先，放在一个数组中，实现查询函数的时候，只需要先对 `times[]` 数组二分搜索，找到比查询时间 `t` 小的最大时刻 `i`，再在选票领先的数组里面输出对应时刻领先的人的编号即可。
