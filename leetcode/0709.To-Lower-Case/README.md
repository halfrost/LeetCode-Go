# [709. To Lower Case](https://leetcode.com/problems/to-lower-case/)


## 题目

Given a string `s`, return *the string after replacing every uppercase letter with the same lowercase letter*.

**Example 1:**

```
Input: s = "Hello"
Output: "hello"
```

**Example 2:**

```
Input: s = "here"
Output: "here"
```

**Example 3:**

```
Input: s = "LOVELY"
Output: "lovely"
```

**Constraints:**

- `1 <= s.length <= 100`
- `s` consists of printable ASCII characters.

## 题目大意

给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。

## 解题思路

- 简单题，将字符串中的大写字母转换成小写字母。

## 代码

```go
func toLowerCase(s string) string {
    runes := [] rune(s)
    diff := 'a' - 'A'
    for i := 0; i < len(s); i++ {
        if runes[i] >= 'A' && runes[i] <= 'Z' {
            runes[i] += diff
        }
    }
    return string(runes)
}
```