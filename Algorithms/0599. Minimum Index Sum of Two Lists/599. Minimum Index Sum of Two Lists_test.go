package leetcode

import (
	"fmt"
	"testing"
)

type question599 struct {
	para599
	ans599
}

// para 是参数
// one 代表第一个参数
type para599 struct {
	one []string
	two []string
}

// ans 是答案
// one 代表第一个答案
type ans599 struct {
	one []string
}

func Test_Problem599(t *testing.T) {

	qs := []question599{

		question599{
			para599{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}},
			ans599{[]string{"Shogun"}},
		},

		question599{
			para599{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}},
			ans599{[]string{"Shogun"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 599------------------------\n")

	for _, q := range qs {
		_, p := q.ans599, q.para599
		fmt.Printf("【input】:%v       【output】:%v\n", p, findRestaurant(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
