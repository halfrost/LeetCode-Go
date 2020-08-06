package structures

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 同一个 TreeNode 的不同表达方式
	//            1
	//      	/  \
	//         2    3
	//        / \  /  \
	//       4  5  6   7
	LeetCodeOrder = []int{1, 2, 3, 4, 5, 6, 7}
	preOrder      = []int{1, 2, 4, 5, 3, 6, 7}
	inOrder       = []int{4, 2, 5, 1, 6, 3, 7}
	postOrder     = []int{4, 5, 2, 6, 7, 3, 1}
)

func Test_Ints2TreeNode(t *testing.T) {
	ast := assert.New(t)

	expected := PreIn2Tree(preOrder, inOrder)
	actual := Ints2TreeNode(LeetCodeOrder)
	ast.Equal(expected, actual)

	actual = Ints2TreeNode([]int{})
	ast.Nil(actual)
}

func Test_preIn2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := Tree2Postorder(PreIn2Tree(preOrder, inOrder))
	expected := postOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { PreIn2Tree([]int{1}, []int{}) })

	ast.Nil(PreIn2Tree([]int{}, []int{}))
}

func Test_inPost2Tree(t *testing.T) {
	ast := assert.New(t)

	actual := Tree2Preorder(InPost2Tree(inOrder, postOrder))
	expected := preOrder
	ast.Equal(expected, actual)

	ast.Panics(func() { InPost2Tree([]int{1}, []int{}) })

	ast.Nil(InPost2Tree([]int{}, []int{}))
}

func Test_tree2Ints(t *testing.T) {
	ast := assert.New(t)
	root := PreIn2Tree(preOrder, inOrder)

	ast.Equal(preOrder, Tree2Preorder(root))
	ast.Nil(Tree2Preorder(nil))

	ast.Equal(inOrder, Tree2Inorder(root))
	ast.Nil(Tree2Inorder(nil))

	ast.Equal(postOrder, Tree2Postorder(root))
	ast.Nil(Tree2Postorder(nil))
}

func Test_indexOf(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(1, indexOf(1, []int{0, 1, 2, 3}))

	ast.Panics(func() { indexOf(0, []int{1, 2, 3}) })
}

func Test_TreeNode_Equal(t *testing.T) {
	type args struct {
		a *TreeNode
	}
	tests := []struct {
		name   string
		fields args
		args   args
		want   bool
	}{

		{
			"相等",
			args{Ints2TreeNode([]int{1, 2, 3, 4, 5})},
			args{Ints2TreeNode([]int{1, 2, 3, 4, 5})},
			true,
		},

		{
			"不相等",
			args{Ints2TreeNode([]int{1, 2, 3, 4, 5})},
			args{Ints2TreeNode([]int{1, 2, 3, NULL, 5})},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := tt.fields.a
			if got := tn.Equal(tt.args.a); got != tt.want {
				t.Errorf("TreeNode.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetTargetNode(t *testing.T) {
	ints := []int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}
	root := Ints2TreeNode(ints)

	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{

		{
			"找到 root.Right.Right",
			args{
				root:   root,
				target: 8,
			},
			root.Right.Right,
		},

		{
			"找到 root.Left.Left",
			args{
				root:   root,
				target: 6,
			},
			root.Left.Left,
		},

		//
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTargetNode(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTargetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tree2ints(t *testing.T) {
	ast := assert.New(t)

	root := PreIn2Tree(preOrder, inOrder)
	actual := LeetCodeOrder
	expected := Tree2ints(root)
	ast.Equal(expected, actual)
}
