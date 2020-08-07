# [497. Random Point in Non-overlapping Rectangles](https://leetcode.com/problems/random-point-in-non-overlapping-rectangles)


## 题目:

Given a list of **non-overlapping** axis-aligned rectangles `rects`, write a function `pick` which randomly and uniformily picks an **integer point** in the space covered by the rectangles.

Note:

1. An **integer point** is a point that has integer coordinates.
2. A point on the perimeter of a rectangle is **included** in the space covered by the rectangles.
3. `i`th rectangle = `rects[i]` = `[x1,y1,x2,y2]`, where `[x1, y1]` are the integer coordinates of the bottom-left corner, and `[x2, y2]` are the integer coordinates of the top-right corner.
4. length and width of each rectangle does not exceed `2000`.
5. `1 <= rects.length <= 100`
6. `pick` return a point as an array of integer coordinates `[p_x, p_y]`
7. `pick` is called at most `10000` times.

**Example 1:**

    Input: 
    ["Solution","pick","pick","pick"]
    [[[[1,1,5,5]]],[],[],[]]
    Output: 
    [null,[4,1],[4,1],[3,3]]

**Example 2:**

    Input: 
    ["Solution","pick","pick","pick","pick","pick"]
    [[[[-2,-2,-1,-1],[1,0,3,0]]],[],[],[],[],[]]
    Output: 
    [null,[-1,-2],[2,0],[-2,-1],[3,0],[-2,-2]]

**Explanation of Input Syntax:**

The input is two lists: the subroutines called and their arguments. `Solution`'s constructor has one argument, the array of rectangles `rects`. `pick` has no arguments. Arguments are always wrapped with a list, even if there aren't any.


## 题目大意

给定一个非重叠轴对齐矩形的列表 rects，写一个函数 pick 随机均匀地选取矩形覆盖的空间中的整数点。

提示：

1. 整数点是具有整数坐标的点。
2. 矩形周边上的点包含在矩形覆盖的空间中。
3. 第 i 个矩形 rects [i] = [x1，y1，x2，y2]，其中 [x1，y1] 是左下角的整数坐标，[x2，y2] 是右上角的整数坐标。
4. 每个矩形的长度和宽度不超过 2000。
5. 1 <= rects.length <= 100
6. pick 以整数坐标数组 [p_x, p_y] 的形式返回一个点。
7. pick 最多被调用10000次。


输入语法的说明：

输入是两个列表：调用的子例程及其参数。Solution 的构造函数有一个参数，即矩形数组 rects。pick 没有参数。参数总是用列表包装的，即使没有也是如此。


## 解题思路


- 给出一个非重叠轴对齐矩形列表，每个矩形用左下角和右上角的两个坐标表示。要求 `pick()` 随机均匀地选取矩形覆盖的空间中的整数点。
- 这一题是第 528 题的变种题，这一题权重是面积，按权重（面积）选择一个矩形，然后再从矩形中随机选择一个点即可。思路和代码和第 528 题一样。
