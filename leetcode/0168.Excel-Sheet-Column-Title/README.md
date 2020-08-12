# [168. Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/)

## 题目

Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

```
		1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...
```

**Example 1**:

```
Input: 1
Output: "A"
```

**Example 2**:

```
Input: 28
Output: "AB"
```

**Example 3**:

```
Input: 701
Output: "ZY"
```

## 题目大意

给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...


## 解题思路

- 给定一个正整数，返回它在 Excel 表中的对应的列名称
- 简单题。这一题就类似短除法的计算过程。以 26 进制的字母编码。按照短除法先除，然后余数逆序输出即可。

## 代码

```go

package leetcode

func convertToTitle(n int) string {
	result := []byte{}
	for n > 0 {
		result = append(result, 'A'+byte((n-1)%26))
		n = (n - 1) / 26
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

```