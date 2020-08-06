package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	ast := assert.New(t)

	s := NewStack()
	ast.True(s.IsEmpty(), "检查新建的 s 是否为空")

	start, end := 0, 100

	for i := start; i < end; i++ {
		s.Push(i)
		ast.Equal(i-start+1, s.Len(), "Push 后检查 q 的长度。")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "从 s 中 pop 出数来。")
	}

	ast.True(s.IsEmpty(), "检查 Pop 完毕后的 s 是否为空")
}
