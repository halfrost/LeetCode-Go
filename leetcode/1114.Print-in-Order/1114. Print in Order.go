package leetcode

import "fmt"

type Foo struct {
	chans []chan bool
}

func NewFoo() *Foo {
	pt := Foo{}
	pt.chans = make([]chan bool, 0)
	for i := 0; i < 2; i++ {
		pt.chans = append(pt.chans, make(chan bool, 1))
	}
	return &pt
}

func (this *Foo) first() {
	fmt.Print("first")
	this.chans[0] <- true
}

func (this *Foo) second() {
	select {
	case <-this.chans[0]:
		fmt.Print("second")
		this.chans[1] <- true
	}
}

func (this *Foo) third() {
	select {
	case <-this.chans[1]:
		fmt.Print("third")
	}
}
