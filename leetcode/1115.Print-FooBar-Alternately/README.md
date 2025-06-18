# [1115. Print FooBar Alternately](https://leetcode.com/problems/print-foobar-alternately/)

## 题目

Suppose you are given the following code:

```java
class FooBar {
  public void foo() {
    for (int i = 0; i < n; i++) {
      print("foo");
    }
  }

  public void bar() {
    for (int i = 0; i < n; i++) {
      print("bar");
    }
  }
}
```

The same instance of `FooBar` will be passed to two different threads:
- thread A will call `foo()`, while
- thread B will call `bar()`.
Modify the given program to output `"foobar"` n times.

**Example 1:**

    Input: n = 1
    Output: "foobar"
    Explanation: There are two threads being fired asynchronously. One of them calls foo(), while the other calls bar(). "foobar" is being output 1 time.

**Example 2:**

    Input: nums = 2
    Output: "foobarfoobar"
    Explanation: "foobar" is being output 2 times.

**Constraints:**

- `1 <= n <= 1000`


## 题目大意

实现一个可以指定次数打印“foobar”字符串的类，foo和bar由两个线程并行打印。

## 解题思路

- 限制打印顺序依赖的思路同1114题，通过Channel实现。额外增加计数字段解决打印计数问题。注意重入问题：每次打印前要检查是否超过打印次数，如超过则直接退出。
