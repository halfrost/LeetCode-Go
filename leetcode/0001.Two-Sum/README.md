# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```



## 题目大意

在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

## 解题思路

这道题最优的做法时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。

C++: 一遍哈希表

```c++
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int,int> map_t;
        vector<int> result(2,-1);
        for(int i = 0; i < nums.size(); i++){
            int another = target - nums[i];
            if(map_t.count(another) > 0){
                result[0] = i;
                result[1] = map_t[another];
                break;
            }
            map_t[nums[i]] = i;
        }
        return result;

    }
};
```

C++: 暴力算法

```c++
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> result(2,-1);
        for(int i = 0; i < nums.size()-1; i++){
            for(int j = i + 1; j < nums.size(); j++)
                if(nums[i] + nums[j] == target){
                    result[0] = i;
                    result[1] = j;
                }
            }
        return result;
    }
};
```

