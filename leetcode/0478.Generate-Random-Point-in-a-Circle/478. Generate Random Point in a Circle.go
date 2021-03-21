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
