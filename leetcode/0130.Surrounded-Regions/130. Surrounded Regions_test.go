package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func clone130(board [][]byte) [][]byte {
	res := make([][]byte, len(board))
	for i := range board {
		res[i] = make([]byte, len(board[i]))
		copy(res[i], board[i])
	}
	return res
}

type question130 struct {
	para130
	ans130
}

// para 是参数
// one 代表第一个参数
type para130 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans130 struct {
	one [][]byte
}

func Test_Problem130(t *testing.T) {

	qs := []question130{

		{
			para130{[][]byte{}},
			ans130{[][]byte{}},
		},

		{
			para130{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}},
			ans130{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		},

		{
			para130{[][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}},
			ans130{[][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 130------------------------\n")

	for _, q := range qs {
		a, p := q.ans130, q.para130
		fmt.Printf("【input】:%v      ", p)
		b1 := clone130(p.one)
		solve1(b1)
		b2 := clone130(p.one)
		solve(b2)
		if !reflect.DeepEqual(b1, b2) {
			t.Fatalf("solve and solve1 differ: solve1=%v solve=%v", b1, b2)
		}
		if len(a.one) != 0 && !reflect.DeepEqual(b2, a.one) {
			t.Fatalf("got %v, want %v", b2, a.one)
		}
		fmt.Printf("【output】:%v      \n", b2)
	}
	fmt.Printf("\n\n\n")
}
