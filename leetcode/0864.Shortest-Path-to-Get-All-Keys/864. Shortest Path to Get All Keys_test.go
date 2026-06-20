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

		{
			para864{[]string{".##..##..B", "##...#...#", "..##..#...", ".#..#b...#", "#.##.a.###", ".#....#...", ".##..#.#..", ".....###@.", "..........", ".........A"}},
			ans864{11},
		},

		{
			para864{[]string{"Dd#b@", ".fE.e", "##.B.", "#.cA.", "aF.#C"}},
			ans864{14},
		},

		{
			para864{[]string{"@...a", ".###A", "b.BCc"}},
			ans864{10},
		},

		{
			para864{[]string{"@Aa"}},
			ans864{-1},
		},

		{
			para864{[]string{"b", "A", "a", "@", "B"}},
			ans864{3},
		},

		{
			para864{[]string{"@.a.#", "#####", "b.A.B"}},
			ans864{-1},
		},

		{
			para864{[]string{"@.a.#", "###.#", "b.A.B"}},
			ans864{8},
		},

		{
			para864{[]string{"@..aA", "..B#.", "....b"}},
			ans864{6},
		},

		{
			para864{[]string{}},
			ans864{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 864------------------------\n")

	for _, q := range qs {
		a, p := q.ans864, q.para864
		got := shortestPathAllKeys(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("input %v: expected %v, got %v", p.one, a.one, got)
		}
		// 解法二是指数级 DFS（剪枝不足，作者已注明会超时）。大网格（如第一个 10x10
		// 用例）会跑到分钟级导致测试卡住，所以只在小网格上调用它来覆盖代码，
		// 大网格只验证解法一。
		if len(p.one) <= 5 {
			if got1 := shortestPathAllKeys1(p.one); got1 != a.one {
				t.Fatalf("input %v: expected %v, shortestPathAllKeys1 got %v", p.one, a.one, got1)
			}
		}
	}
	fmt.Printf("\n\n\n")
}
