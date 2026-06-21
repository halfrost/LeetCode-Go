package leetcode

import (
	"fmt"
	"testing"
)

type question794 struct {
	para794
	ans794
}

// para 是参数
type para794 struct {
	board []string
}

// ans 是答案
type ans794 struct {
	ans bool
}

func Test_Problem794(t *testing.T) {

	qs := []question794{

		{
			para794{[]string{"O  ", "   ", "   "}},
			ans794{false},
		},

		{
			para794{[]string{"XOX", " X ", "   "}},
			ans794{false},
		},

		{
			para794{[]string{"XXX", "   ", "OOO"}},
			ans794{false},
		},

		{
			para794{[]string{"XOX", "O O", "XOX"}},
			ans794{true},
		},

		{
			// X wins via right column, cntX == cntO -> process(board, 'X') column branch
			para794{[]string{"  X", " OX", "OOX"}},
			ans794{false},
		},

		{
			// O wins via anti-diagonal, cntX == cntO -> process(board, 'O')
			para794{[]string{"XXO", "XOX", "OXO"}},
			ans794{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 794------------------------\n")

	for _, q := range qs {
		_, p := q.ans794, q.para794
		fmt.Printf("【input】:%v    【output】:%v\n", p.board, validTicTacToe(p.board))
	}
	fmt.Printf("\n\n\n")
}
