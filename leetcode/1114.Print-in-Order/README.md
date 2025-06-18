# [1114. Print in Order](https://leetcode.com/problems/print-in-order/)

## 题目

Suppose we have a class:

```java
public class Foo {
  public void first() { print("first"); }
  public void second() { print("second"); }
  public void third() { print("third"); }
}
```

The same instance of Foo will be passed to three different threads. Thread A will call `first()`, thread B will call `second()`, and thread C will call `third()`. Design a mechanism and modify the program to ensure that `second()` is executed after `first()`, and `third()` is executed after `second()`.

**Note:**

We do not know how the threads will be scheduled in the operating system, even though the numbers in the input seem to imply the ordering. The input format you see is mainly to ensure our tests' comprehensiveness.

**Example 1:**

    Input: nums = [1,2,3]
    Output: "firstsecondthird"
    Explanation: There are three threads being fired asynchronously. The input [1,2,3] means thread A calls first(), thread B calls second(), and thread C calls third(). "firstsecondthird" is the correct output.

**Example 2:**

    Input: nums = [1,3,2]
    Output: "firstsecondthird"
    Explanation: The input [1,3,2] means thread A calls first(), thread B calls third(), and thread C calls second(). "firstsecondthird" is the correct output.

**Constraints:**

- `nums` is a permutation of `[1, 2, 3]`.


## 题目大意

实现一个可以按需打印“first”、“second”和“third”的类，无论并发调用的顺序如何，均可保持按从小到大打印的位次。

## 解题思路

- Golang中可以使用Channel实现多个协程执行的顺序依赖。通过有缓冲区的Channel实现写入不阻塞，读取阻塞的功能；通过select使预期靠后执行的协程阻塞等待其它协程执行完毕。
