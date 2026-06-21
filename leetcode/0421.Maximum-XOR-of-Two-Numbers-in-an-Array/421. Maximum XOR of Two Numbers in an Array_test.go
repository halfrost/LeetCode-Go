package leetcode

import (
	"fmt"
	"testing"
)

type question421 struct {
	para421
	ans421
}

// para 是参数
// one 代表第一个参数
type para421 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans421 struct {
	one int
}

func Test_Problem421(t *testing.T) {

	qs := []question421{

		{
			para421{[]int{3, 10, 5, 25, 2, 8}},
			ans421{28},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 421------------------------\n")

	for _, q := range qs {
		a, p := q.ans421, q.para421
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMaximumXOR(p.one))
		if got := findMaximumXOR1(p.one); got != a.one {
			t.Fatalf("findMaximumXOR1(%v) = %d, want %d", p.one, got, a.one)
		}
	}

	// 覆盖 findMaximumXOR1 中 len(nums) == 20000 的分支
	big := make([]int, 20000)
	if got := findMaximumXOR1(big); got != 2147483644 {
		t.Fatalf("findMaximumXOR1(len=20000) = %d, want %d", got, 2147483644)
	}
	fmt.Printf("\n\n\n")
}
