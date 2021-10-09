# [498. Diagonal Traverse](https://leetcode.com/problems/diagonal-traverse/)


## 题目

Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

**Example:**

    Input:
    [
     [ 1, 2, 3 ],
     [ 4, 5, 6 ],
     [ 7, 8, 9 ]
    ]
    
    Output:  [1,2,4,7,5,3,6,8,9]
    
    Explanation:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/diagonal_traverse.png)

**Note:**

The total number of elements of the given matrix will not exceed 10,000.


## 题目大意

给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/diagonal_traverse.png)

说明: 给定矩阵中的元素总数不会超过 100000 。

## 解题思路

- 给出一个二维数组，要求按照如图的方式遍历整个数组。
- 这一题用模拟的方式就可以解出来。需要注意的是边界条件：比如二维数组为空，二维数组退化为一行或者一列，退化为一个元素。具体例子见测试用例。
- 解题关键是在判断下一个位置，将矩阵想像成一个X,Y坐标轴。那么可分为以下几种情况，
    1、斜角向右上遍历时，
    当右上角在坐标轴内， 正常计算 即， x+1(X轴向右移动)， y-1(Y轴向上移动)
    当右上角在坐标轴外，那么当前位置只能在 第一行X坐标轴 ，或者 最后一列Y坐标轴 ， 即判断该两种情况下�应该X坐标往右，或者 Y坐标往上
    2、同理 斜角向下遍历时
    当左下角在坐标轴内，正常计算 即， x-1(X轴向右移动)， y+1(Y轴向下移动)
    当左下角在坐标轴外，那么当前位置只能在 第一列Y坐标轴，或者 最后一行X坐标轴， 即判断该两种情况下�应该X坐标往左，或者 Y坐标往下
