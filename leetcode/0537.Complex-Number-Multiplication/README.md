# [537. Complex Number Multiplication](https://leetcode.com/problems/complex-number-multiplication/)


## 题目

Given two strings representing two [complex numbers](https://en.wikipedia.org/wiki/Complex_number).

You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

**Example 1**:

```
Input: "1+1i", "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
```

**Example 2**:

```
Input: "1+-1i", "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
```

**Note**:

1. The input strings will not have extra blank.
2. The input strings will be given in the form of **a+bi**, where the integer **a** and **b** will both belong to the range of [-100, 100]. And **the output should be also in this form**.

## 题目大意

给定两个表示复数的字符串。返回表示它们乘积的字符串。注意，根据定义 i^2 = -1 。

注意:

- 输入字符串不包含额外的空格。
- 输入字符串将以 a+bi 的形式给出，其中整数 a 和 b 的范围均在 [-100, 100] 之间。输出也应当符合这种形式。



## 解题思路

- 给定 2 个字符串，要求这两个复数的乘积，输出也是字符串格式。
- 数学题。按照复数的运算法则，i^2 = -1，最后输出字符串结果即可。

## 代码

```go

package leetcode

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	realA, imagA := parse(a)
	realB, imagB := parse(b)
	real := realA*realB - imagA*imagB
	imag := realA*imagB + realB*imagA
	return strconv.Itoa(real) + "+" + strconv.Itoa(imag) + "i"
}

func parse(s string) (int, int) {
	ss := strings.Split(s, "+")
	r, _ := strconv.Atoi(ss[0])
	i, _ := strconv.Atoi(ss[1][:len(ss[1])-1])
	return r, i
}

```