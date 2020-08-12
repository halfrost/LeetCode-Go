# [832. Flipping an Image](https://leetcode.com/problems/flipping-an-image/)


## 题目

Given a binary matrix `A`, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed. For example, flipping `[1, 1, 0]` horizontally results in `[0, 1, 1]`.

To invert an image means that each `0` is replaced by `1`, and each `1` is replaced by `0`. For example, inverting `[0, 1, 1]` results in `[1, 0, 0]`.

**Example 1**:

```
Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
```

**Example 2**:

```
Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
```

**Notes**:

- `1 <= A.length = A[0].length <= 20`
- `0 <= A[i][j] <= 1`

## 题目大意

给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。


## 解题思路

- 给定一个二进制矩阵，要求先水平翻转，然后再反转( 1→0 , 0→1 )。
- 简单题，按照题意先水平翻转，再反转即可。

## 代码

```go

package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		for a, b := 0, len(A[i])-1; a < b; a, b = a+1, b-1 {
			A[i][a], A[i][b] = A[i][b], A[i][a]
		}
		for a := 0; a < len(A[i]); a++ {
			A[i][a] = (A[i][a] + 1) % 2
		}
	}
	return A
}

```