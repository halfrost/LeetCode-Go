package leetcode

import (
	"fmt"
	"testing"
)

type question526 struct {
	para526
	ans526
}

// para 是参数
// one 代表第一个参数
type para526 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans526 struct {
	one int
}

func Test_Problem526(t *testing.T) {

	qs := []question526{

		{
			para526{1},
			ans526{1},
		},

		{
			para526{2},
			ans526{2},
		},

		{
			para526{3},
			ans526{3},
		},

		{
			para526{4},
			ans526{8},
		},

		{
			para526{5},
			ans526{10},
		},

		{
			para526{6},
			ans526{36},
		},

		{
			para526{7},
			ans526{41},
		},

		{
			para526{8},
			ans526{132},
		},

		{
			para526{9},
			ans526{250},
		},

		{
			para526{10},
			ans526{700},
		},

		{
			para526{11},
			ans526{750},
		},

		{
			para526{12},
			ans526{4010},
		},

		{
			para526{13},
			ans526{4237},
		},

		{
			para526{14},
			ans526{10680},
		},

		{
			para526{15},
			ans526{24679},
		},

		{
			para526{16},
			ans526{87328},
		},

		{
			para526{17},
			ans526{90478},
		},

		{
			para526{18},
			ans526{435812},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 526------------------------\n")

	for _, q := range qs {
		_, p := q.ans526, q.para526
		fmt.Printf("【input】:%v       【output】:%v\n", p, countArrangement(p.one))
	}
	fmt.Printf("\n\n\n")
}
