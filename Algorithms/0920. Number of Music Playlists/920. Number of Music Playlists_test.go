package leetcode

import (
	"fmt"
	"testing"
)

type question920 struct {
	para920
	ans920
}

// para 是参数
// one 代表第一个参数
type para920 struct {
	N int
	L int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans920 struct {
	one int
}

func Test_Problem920(t *testing.T) {

	qs := []question920{

		question920{
			para920{3, 3, 1},
			ans920{6},
		},

		question920{
			para920{2, 3, 0},
			ans920{6},
		},

		question920{
			para920{2, 3, 1},
			ans920{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 920------------------------\n")

	for _, q := range qs {
		_, p := q.ans920, q.para920
		fmt.Printf("【input】:%v       【output】:%v\n", p, numMusicPlaylists(p.N, p.L, p.K))
	}
	fmt.Printf("\n\n\n")
}
