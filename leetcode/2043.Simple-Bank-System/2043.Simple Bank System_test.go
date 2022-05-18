package leetcode

import (
	"fmt"
	"testing"
)

type question2043 struct {
	para2043
	ans2043
}

// para 是参数
type para2043 struct {
	ops  []string
	para [][]int64
}

// ans 是答案
type ans2043 struct {
	ans []bool
}

func Test_Problem2043(t *testing.T) {

	qs := []question2043{

		{
			para2043{
				[]string{"Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"},
				[][]int64{{10, 100, 20, 50, 30}, {3, 10}, {5, 1, 20}, {5, 20}, {3, 4, 15}, {10, 50}}},
			ans2043{[]bool{true, true, true, false, false}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2043------------------------\n")

	for _, q := range qs {
		var b Bank
		var res []bool
		_, p := q.ans2043, q.para2043
		for i, op := range p.ops {
			if op == "Bank" {
				b = Constructor(q.para[i])
			} else if op == "withdraw" {
				isSuccess := b.Withdraw(int(p.para[i][0]), p.para[i][1])
				res = append(res, isSuccess)
			} else if op == "transfer" {
				isSuccess := b.Transfer(int(p.para[i][0]), int(p.para[i][0]), p.para[i][2])
				res = append(res, isSuccess)
			} else if op == "deposit" {
				isSuccess := b.Deposit(int(p.para[i][0]), p.para[i][1])
				res = append(res, isSuccess)
			} else {
				fmt.Println("unknown operation")
			}
		}
		fmt.Printf("【input】:%v      \n", p)
		fmt.Printf("【output】:%v      \n", res)
	}
	fmt.Printf("\n\n\n")
}
