# [65. Valid Number](https://leetcode.com/problems/valid-number/)


## 题目

A **valid number** can be split up into these components (in order):

   1. A **decimal number** or an integer. 
   2. (Optional) An 'e' or 'E', followed by an **integer.**
   
A **decimal number** can be split up into these components (in order):

   1. (Optional) A sign character (either '+' or '-').
   2. One of the following formats:
        1. One or more digits, followed by a dot '.'.
        2. One or more digits, followed by a dot '.', followed by one or more digits.
        3. A dot '.', followed by one or more digits.
        
An **integer** can be split up into these components (in order):

   1. (Optional) A sign character (either '+' or '-').
   2. One or more digits.
   
For example, all the following are valid numbers: `["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]`, while the following are not valid numbers: `["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"].`

Given a string s, return true if s is a **valid number.**

**Example:**

    Input: s = "0"
    Output: true
    
    Input: s = "e"
    Output: false

## 题目大意

给定一个字符串S，请根据以上的规则判断该字符串是否是一个有效的数字字符串。


## 解题思路

- 用三个变量分别标记是否出现过数字、是否出现过'.'和 是否出现过 'e/E'
- 从左到右依次遍历字符串中的每一个元素
    - 如果是数字，则标记数字出现过
    - 如果是 '.', 则需要 '.'没有出现过，并且 'e/E' 没有出现过，才会进行标记
    - 如果是 'e/E', 则需要 'e/E'没有出现过，并且前面出现过数字，才会进行标记
    - 如果是 '+/-', 则需要是第一个字符，或者前一个字符是 'e/E'，才会进行标记，并重置数字出现的标识
    - 最后返回时，需要字符串中至少出现过数字，避免下列case: s == '.' or 'e/E' or '+/e' and etc...

## 代码

```go

package leetcode

func isNumber(s string) bool {
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && !eFlag && numFlag {
			eFlag = true
			numFlag = false // reJudge integer after 'e' or 'E'
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			continue
		} else {
			return false
		}
	}
	// avoid case: s == '.' or 'e/E' or '+/-' and etc...
	// string s must have num
	return numFlag
}

```