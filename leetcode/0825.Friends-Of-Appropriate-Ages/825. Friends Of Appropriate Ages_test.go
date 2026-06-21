package leetcocde

import (
	"fmt"
	"testing"
)

type question825 struct {
	para825
	ans825
}

// para 是参数
// one 代表第一个参数
type para825 struct {
	ages []int
}

// ans 是答案
// one 代表第一个答案
type ans825 struct {
	one int
}

func Test_Problem825(t *testing.T) {

	qs := []question825{

		{
			para825{[]int{16, 16}},
			ans825{2},
		},

		{
			para825{[]int{16, 17, 18}},
			ans825{2},
		},

		{
			para825{[]int{20, 30, 100, 110, 120}},
			ans825{3},
		},

		{
			para825{[]int{10, 16, 16}},
			ans825{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 825------------------------\n")

	for _, q := range qs {
		a, p := q.ans825, q.para825
		got := numFriendRequests(append([]int{}, p.ages...))
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("numFriendRequests(%v) = %d, want %d", p.ages, got, a.one)
		}
		if got1 := numFriendRequests1(append([]int{}, p.ages...)); got1 != a.one {
			t.Fatalf("numFriendRequests1(%v) = %d, want %d", p.ages, got1, a.one)
		}
		if got2 := numFriendRequests2(append([]int{}, p.ages...)); got2 != a.one {
			t.Fatalf("numFriendRequests2(%v) = %d, want %d", p.ages, got2, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
