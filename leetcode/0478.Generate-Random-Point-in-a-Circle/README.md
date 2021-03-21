# [478. Generate Random Point in a Circle](https://leetcode.com/problems/generate-random-point-in-a-circle/)


## 题目

Given the radius and x-y positions of the center of a circle, write a function `randPoint` which generates a uniform random point in the circle.

Note:

1. input and output values are in [floating-point](https://www.webopedia.com/TERM/F/floating_point_number.html).
2. radius and x-y position of the center of the circle is passed into the class constructor.
3. a point on the circumference of the circle is considered to be in the circle.
4. `randPoint` returns a size 2 array containing x-position and y-position of the random point, in that order.

**Example 1:**

```
Input: 
["Solution","randPoint","randPoint","randPoint"]
[[1,0,0],[],[],[]]
Output: [null,[-0.72939,-0.65505],[-0.78502,-0.28626],[-0.83119,-0.19803]]

```

**Example 2:**

```
Input: 
["Solution","randPoint","randPoint","randPoint"]
[[10,5,-7.5],[],[],[]]
Output: [null,[11.52438,-8.33273],[2.46992,-16.21705],[11.13430,-12.42337]]
```

**Explanation of Input Syntax:**

The input is two lists: the subroutines called and their arguments. `Solution`'s constructor has three arguments, the radius, x-position of the center, and y-position of the center of the circle. `randPoint` has no arguments. Arguments are always wrapped with a list, even if there aren't any.

## 题目大意

给定圆的半径和圆心的 x、y 坐标，写一个在圆中产生均匀随机点的函数 randPoint 。

说明:

- 输入值和输出值都将是浮点数。
- 圆的半径和圆心的 x、y 坐标将作为参数传递给类的构造函数。
- 圆周上的点也认为是在圆中。
- randPoint 返回一个包含随机点的x坐标和y坐标的大小为2的数组。

## 解题思路

- 随机产生一个圆内的点，这个点一定满足定义 `(x-a)^2+(y-b)^2 ≤ R^2`，其中 `(a,b)` 是圆的圆心坐标，`R` 是半径。
- 先假设圆心坐标在 (0,0)，这样方便计算，最终输出坐标的时候整体加上圆心的偏移量即可。`rand.Float64()` 产生一个 `[0.0,1.0)` 区间的浮点数。`-R ≤ 2 * R * rand() - R < R`，利用随机产生坐标点的横纵坐标 `(x,y)` 与半径 R 的关系，如果 `x^2 + y^2 ≤ R^2`，那么说明产生的点在圆内。最终输出的时候要记得加上圆心坐标的偏移值。

## 代码

```go
package leetcode

import (
	"math"
	"math/rand"
	"time"
)

type Solution struct {
	r float64
	x float64
	y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	/*
	   a := angle()
	   r := this.r * math.Sqrt(rand.Float64())
	   x := r * math.Cos(a) + this.x
	   y := r * math.Sin(a) + this.y
	   return []float64{x, y}*/
	for {
		rx := 2*rand.Float64() - 1.0
		ry := 2*rand.Float64() - 1.0
		x := this.r * rx
		y := this.r * ry
		if x*x+y*y <= this.r*this.r {
			return []float64{x + this.x, y + this.y}
		}
	}
}

func angle() float64 {
	return rand.Float64() * 2 * math.Pi
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
```