# [93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/)


## 题目

Given a string containing only digits, restore it by returning all possible valid IP address combinations.

**Example:**

    Input: "25525511135"
    Output: ["255.255.11.135", "255.255.111.35"]

## 题目大意

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

## 解题思路

- DFS 深搜
- 需要注意的点是 IP 的规则，以 0 开头的数字和超过 255 的数字都为非法的。

