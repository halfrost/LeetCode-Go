# [475. Heaters](https://leetcode.com/problems/heaters/)

## 题目:

Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.

Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.

So, your input will be the positions of houses and heaters seperately, and your expected output will be the minimum radius standard of heaters.

**Note:**

1. Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
2. Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
3. As long as a house is in the heaters' warm radius range, it can be warmed.
4. All the heaters follow your radius standard and the warm radius will the same.

**Example 1:**

    Input: [1,2,3],[2]
    Output: 1
    Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.

**Example 2:**

    Input: [1,2,3,4],[1,4]
    Output: 1
    Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.



## 题目大意


冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径。

说明:

- 给出的房屋和供暖器的数目是非负数且不会超过 25000。
- 给出的房屋和供暖器的位置均是非负数且不会超过10^9。
- 只要房屋位于供暖器的半径内(包括在边缘上)，它就可以得到供暖。
- 所有供暖器都遵循你的半径标准，加热的半径也一样。



## 解题思路


- 给出一个房子坐标的数组，和一些供暖器坐标的数组，分别表示房子的坐标和供暖器的坐标。要求找到供暖器最小的半径能使得所有的房子都能享受到暖气。
- 这一题可以用暴力的解法，暴力解法依次遍历每个房子的坐标，再遍历每个供暖器，找到距离房子最近的供暖器坐标。在所有这些距离的长度里面找到最大值，这个距离的最大值就是供暖器半径的最小值。时间复杂度 O(n^2)。
- 这一题最优解是二分搜索。在查找距离房子最近的供暖器的时候，先将供暖器排序，然后用二分搜索的方法查找。其他的做法和暴力解法一致。时间复杂度 O(n log n)。
