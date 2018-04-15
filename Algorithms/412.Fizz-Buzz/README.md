# [412. Fizz Buzz](https://leetcode.com/problems/fizz-buzz/)

## 题目

Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

```text
n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
```

## 题目大意

3的倍数输出 "Fizz"，5的倍数输出 "Buzz"，15的倍数输出 "FizzBuzz"，其他时候都输出原本的数字。