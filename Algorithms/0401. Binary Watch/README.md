# [401. Binary Watch](https://leetcode.com/problems/binary-watch/)


## 题目

A binary watch has 4 LEDs on the top which represent the **hours** (**0-11**), and the 6 LEDs on the bottom represent the **minutes** (**0-59**).

Each LED represents a zero or one, with the least significant bit on the right.

![](https://upload.wikimedia.org/wikipedia/commons/8/8b/Binary_clock_samui_moon.jpg)

For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

**Example:**

    Input: n = 1
    Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

**Note:**

- The order of output does not matter.
- The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
- The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".


## 题目大意

二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。

给定一个非负整数 n 代表当前 LED 亮着的数量，返回二进制表所有可能的时间。


## 解题思路


- 给出数字 n，要求输出二进制表中所有可能的时间
- 题目中比较坑的是，分钟大于 60 的都不应该打印出来，小时大于 12 的也不应该打印出来，因为是非法的。给出的 num 大于 8 的也是非法值，最终结果应该输出空字符串数组。
- 这道题的数据量不大，可以直接用打表法，具体打表函数见 `findReadBinaryWatchMinute()` 和 `findReadBinaryWatchHour()` 这两个函数。
