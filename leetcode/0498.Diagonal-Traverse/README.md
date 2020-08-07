# [498. Diagonal Traverse](https://leetcode.com/problems/diagonal-traverse/)


## 题目:

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
