package leetcode

import "fmt"

type FooBar struct {
	chans   []chan bool
	limit   int
	printed []int
}

func NewFooBar(n int) *FooBar {
	pt := FooBar{}
	pt.limit = n
	pt.printed = make([]int, 2)
	pt.chans = make([]chan bool, 2)
	for idx := range pt.chans {
		pt.chans[idx] = make(chan bool, 1)
	}
	pt.chans[0] <- true
	return &pt
}

func (this *FooBar) foo() {
	for this.printed[0] < this.limit {
		select {
		case <-this.chans[0]:
			if this.printed[0] >= this.limit {
				return
			}
			fmt.Printf("foo")
			this.printed[0]++
			this.chans[1] <- true
		}
	}
}

func (this *FooBar) bar() {
	for this.printed[1] < this.limit {
		select {
		case <-this.chans[1]:
			if this.printed[1] >= this.limit {
				return
			}
			fmt.Printf("bar")
			this.printed[1]++
			this.chans[0] <- true
		}
	}
}
