package leetcode

import (
	"fmt"
	"testing"
)

type question820 struct {
	para820
	ans820
}

// para 是参数
// one 代表第一个参数
type para820 struct {
	words []string
}

// ans 是答案
// one 代表第一个答案
type ans820 struct {
	one int
}

func Test_Problem820(t *testing.T) {

	qs := []question820{

		{
			para820{[]string{"time", "me", "bell"}},
			ans820{10},
		},

		{
			para820{[]string{"t"}},
			ans820{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 820------------------------\n")

	for _, q := range qs {
		a, p := q.ans820, q.para820
		out := minimumLengthEncoding(p.words)
		fmt.Printf("【input】:%v       【output】:%v\n", p, out)
		if out != a.one {
			t.Fatalf("minimumLengthEncoding(%v) = %d, want %d", p.words, out, a.one)
		}
		out1 := minimumLengthEncoding1(p.words)
		if out1 != a.one {
			t.Fatalf("minimumLengthEncoding1(%v) = %d, want %d", p.words, out1, a.one)
		}
	}

	// 覆盖 node 方法的 nil 守卫与未命中分支
	var nilNode *node
	if _, ok := nilNode.has('a'); ok {
		t.Fatalf("nil node has('a') should be false")
	}
	if nilNode.isLeaf() {
		t.Fatalf("nil node isLeaf should be false")
	}
	tree := new(node)
	tree.add([]byte("abc"))
	if got := tree.endNodeOf([]byte("xyz")); got != nil {
		t.Fatalf("endNodeOf(\"xyz\") = %v, want nil", got)
	}

	fmt.Printf("\n\n\n")
}
