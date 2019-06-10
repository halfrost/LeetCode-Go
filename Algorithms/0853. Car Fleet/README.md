# [853. Car Fleet](https://leetcode.com/problems/car-fleet/)

## 题目

N cars are going to the same destination along a one lane road.  The destination is target miles away.

Each car i has a constant speed speed[i] (in miles per hour), and initial position position[i] miles towards the target along the road.

A car can never pass another car ahead of it, but it can catch up to it, and drive bumper to bumper at the same speed.

The distance between these two cars is ignored - they are assumed to have the same position.

A car fleet is some non-empty set of cars driving at the same position and same speed.  Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.


How many car fleets will arrive at the destination?

 

Example 1:

```c
Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
Output: 3
Explanation:
The cars starting at 10 and 8 become a fleet, meeting each other at 12.
The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
The cars starting at 5 and 3 become a fleet, meeting each other at 6.
Note that no other cars meet these fleets before the destination, so the answer is 3.
```

Note:

1. 0 <= N <= 10 ^ 4
2. 0 < target <= 10 ^ 6
3. 0 < speed[i] <= 10 ^ 6
4. 0 <= position[i] < target
5. All initial positions are different.


## 题目大意

汽车舰队。有 N 个车辆，每个车辆都想行驶到终点，但是每个车辆距离终点的距离不同，行驶速度也不同。当一辆汽车追上前一辆汽车以后，会形成“汽车舰队”，即形成舰队的车辆并以第一辆车的速度前进。形成舰队以后，多辆车的长度可以忽略，可以认为这些车辆都叠在一起了。一辆车也是一个舰队，当多辆车一起到达终点，也算是一起形成了舰队。问，最后有多少个“汽车舰队”到达终点了？


## 解题思路

- 根据每辆车距离终点和速度，计算每辆车到达终点的时间，并按照距离从大到小排序(position 越大代表距离终点越近)
- 从头往后扫描排序以后的数组，时间一旦大于前一个 car 的时间，就会生成一个新的 car fleet，最终计数加一即可。


