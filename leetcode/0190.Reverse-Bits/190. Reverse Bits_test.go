package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

type question190 struct {
	para190
	ans190
}

// para 是参数
// one 代表第一个参数
type para190 struct {
	one uint32
}

// ans 是答案
// one 代表第一个答案
type ans190 struct {
	one uint32
}

func Test_Problem190(t *testing.T) {
	qs := []question190{

		{
			para190{43261596},
			ans190{964176192},
		},

		{
			para190{4294967293},
			ans190{3221225471},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 190------------------------\n")

	for _, q := range qs {
		_, p := q.ans190, q.para190
		input := strconv.FormatUint(uint64(p.one), 2) // 32位无符号整数转换为二进制字符串
		input = fmt.Sprintf("%0*v", 32, input)        // 格式化输出32位,保留前置0
		output := reverseBits(p.one)
		outputBin := strconv.FormatUint(uint64(output), 2) // 32位无符号整数转换为二进制字符串
		outputBin = fmt.Sprintf("%0*v", 32, outputBin)     // 格式化输出32位,保留前置0
		fmt.Printf("【input】:%v       【output】:%v (%v)\n", input, output, outputBin)
	}
	fmt.Printf("\n\n\n")
}
