package leetcode

import (
	"fmt"
	"testing"
)

type question864 struct {
	para864
	ans864
}

// para 是参数
// one 代表第一个参数
type para864 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans864 struct {
	one int
}

func Test_Problem864(t *testing.T) {

	qs := []question864{

		question864{
			para864{[]string{".##..##..B", "##...#...#", "..##..#...", ".#..#b...#", "#.##.a.###", ".#....#...", ".##..#.#..", ".....###@.", "..........", ".........A"}},
			ans864{11},
		},

		question864{
			para864{[]string{"Dd#b@", ".fE.e", "##.B.", "#.cA.", "aF.#C"}},
			ans864{14},
		},

		question864{
			para864{[]string{"@...a", ".###A", "b.BCc"}},
			ans864{10},
		},

		question864{
			para864{[]string{"@Aa"}},
			ans864{-1},
		},

		question864{
			para864{[]string{"b", "A", "a", "@", "B"}},
			ans864{3},
		},

		question864{
			para864{[]string{"@.a.#", "#####", "b.A.B"}},
			ans864{-1},
		},

		question864{
			para864{[]string{"@.a.#", "###.#", "b.A.B"}},
			ans864{8},
		},

		question864{
			para864{[]string{"@..aA", "..B#.", "....b"}},
			ans864{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 864------------------------\n")

	for _, q := range qs {
		_, p := q.ans864, q.para864
		fmt.Printf("【input】:%v       【output】:%v\n", p, shortestPathAllKeys(p.one))
	}
	fmt.Printf("\n\n\n")
}
