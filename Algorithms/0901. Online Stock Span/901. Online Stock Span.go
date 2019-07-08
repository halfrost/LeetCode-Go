package leetcode

import "fmt"

// node pair
type Node struct {
	Val int
	res int
}

// slice
type StockSpanner struct {
	Item []Node
}

func Constructor901() StockSpanner {
	stockSpanner := StockSpanner{make([]Node, 0)}
	return stockSpanner
}

// need refactor later
func (this *StockSpanner) Next(price int) int {
	res := 1
	if len(this.Item) == 0 {
		this.Item = append(this.Item, Node{price, res})
		return res
	}
	for len(this.Item) > 0 && this.Item[len(this.Item)-1].Val <= price {
		res = res + this.Item[len(this.Item)-1].res
		this.Item = this.Item[:len(this.Item)-1]
	}
	this.Item = append(this.Item, Node{price, res})
	fmt.Printf("this.Item = %v\n", this.Item)
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
