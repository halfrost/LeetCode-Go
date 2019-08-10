# [324. Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii/)

## 题目

Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

Example 1:

```c
Input: nums = [1, 5, 1, 1, 6, 4]
Output: One possible answer is [1, 4, 1, 5, 1, 6].
```

Example 2:

```c
Input: nums = [1, 3, 2, 2, 3, 1]
Output: One possible answer is [2, 3, 1, 3, 1, 2].
```

Note:   

You may assume all input has valid answer.

Follow Up:   

Can you do it in O(n) time and/or in-place with O(1) extra space?


## 题目大意

给定一个数组，要求给它进行“摆动排序”，“摆动排序”即：nums[0] < nums[1] > nums[2] < nums[3]...

## 解题思路

这一题最直接的方法是先排序，然后用 2 个指针，一个指向下标为 0 的位置，另一个指向下标为 n/2 的位置。最终的数组的奇数位从下标为 0 开始往后取，偶数位从下标为 n/2 中间位置开始往后取。这种方法时间复杂度为 O(n log n)。

题目要求用时间复杂度 O(n) 和 空间复杂度 O(1) 的方法解决。思路如下，先找到数组中间大小的数字，然后把数组分为 2 部分：

```c
Index :       0   1   2   3   4   5
Small half:   M       S       S    
Large half:       L       L       L(M)
```

奇数位排中间数和小于中间数的数字，偶数位排大于中间数的数字和中间数。如果中间数字有多个，那么偶数位最后几位也是中间数，奇数位开头的前几位也是中间数。

举例，给定一个数组如下，中间数是 5 。有 2 个 5 。

```c
13   6   5   5   4   2

         M
```

```c
Step 1: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        13   6    5    5    4    2 
             Left
              i
                                      Right
```
                                      
nums[Mapped_idx[i]] = nums[1] = 6 > 5, 所以可以把 6 放在第 1 个奇数位的位置。left 和 i 同时右移。

```c
Step 2: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        13   6    5    5    4    2 
                  Left
                   i
                                      Right
```
                                    
nums[3] = 5 = 5, 5 可以放在下标为 3 的位置，由于 5 已经和中间数相等了，所以只后移 i 。


```c
Step 3: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        13   6    5    5    4    2 
                  Left
                        i
                                     Right
```

nums[5] = 2 < 5, 因为比中位数小，所以应该放在偶数位的最后 1 位。这里的例子而言，应该放在下标为 4 的位置上。交换 nums[Mapped_idx[i]] 和 nums[Mapped_idx[Right]]，交换完成以后 right 向左移。


```c
Step 4: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        13   6    5    5    2    4 
                  Left
                        i
                               Right
```

nums[5] = 4 < 5, 因为比中位数小，所以应该放在偶数位的当前倒数第一位。这里的例子而言，应该放在下标为 2 的位置上。交换 nums[Mapped\_idx[i]] 和 nums[Mapped\_idx[Right]]，交换完成以后 right 向左移。


```c
Step 5: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        13   6    4    5    2    5 
                  Left
                        i
                            Right
```

nums[5] = 5 = 5, 由于 5 已经和中间数相等了，所以只后移 i 。

```c
Step 6: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        13   6    4    5    2    5 
                  Left
                             i
                            Right
```


nums[0] = 13 > 5, 由于 13 比中位数大，所以可以把 13 放在第 2 个奇数位的位置，并移动 left 和 i 。


```c
Step Final: 
Original idx: 0    1    2    3    4    5  
Mapped idx:   1    3    5    0    2    4 
Array:        5    6    4    13   2    5 
                      Left
                                  i
                            Right
```
                            
i > Right, 退出循环，最终摆动排序的结果是 5 6 4 13 2 5 。

具体时间见代码，时间复杂度 O(n) 和 空间复杂度 O(1)。

















