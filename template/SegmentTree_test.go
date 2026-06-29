package template

import "testing"

func sum(i, j int) int { return i + j }
func maxv(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 单点更新 + 区间查询（求和语义）
func Test_SegmentTree_PointUpdate(t *testing.T) {
	st := SegmentTree{}
	st.Init([]int{1, 3, 5, 7, 9, 11}, sum)

	checks := []struct {
		l, r, want int
	}{
		{0, 5, 36},
		{1, 3, 15},
		{2, 2, 5},
		{0, 0, 1},
	}
	for _, c := range checks {
		if got := st.Query(c.l, c.r); got != c.want {
			t.Fatalf("Query(%d,%d) = %d, want %d", c.l, c.r, got, c.want)
		}
	}

	st.Update(2, 6) // nums[2]: 5 -> 6
	if got := st.Query(1, 3); got != 16 {
		t.Fatalf("after Update, Query(1,3) = %d, want 16", got)
	}
	if got := st.Query(0, 5); got != 37 {
		t.Fatalf("after Update, Query(0,5) = %d, want 37", got)
	}
}

// 区间更新 + 区间查询（max 幂等语义，等价于 Falling Squares / Skyline 的用法）
// 这是本次把 lazy 下推从 O(区间) 改成 O(1) 单次 merge 所影响的路径。
func Test_SegmentTree_Lazy_Max(t *testing.T) {
	st := SegmentTree{}
	st.Init([]int{0, 0, 0, 0, 0}, maxv)

	st.UpdateLazy(0, 2, 5) // [5,5,5,0,0]
	if got := st.QueryLazy(0, 4); got != 5 {
		t.Fatalf("QueryLazy(0,4) = %d, want 5", got)
	}
	if got := st.QueryLazy(3, 4); got != 0 {
		t.Fatalf("QueryLazy(3,4) = %d, want 0", got)
	}
	if got := st.QueryLazy(1, 2); got != 5 {
		t.Fatalf("QueryLazy(1,2) = %d, want 5", got)
	}

	st.UpdateLazy(2, 4, 3) // max into [2,4]: [5,5,5,3,3]
	cases := []struct {
		l, r, want int
	}{
		{0, 4, 5},
		{3, 4, 3},
		{2, 2, 5},
		{4, 4, 3},
		{0, 0, 5},
	}
	for _, c := range cases {
		if got := st.QueryLazy(c.l, c.r); got != c.want {
			t.Fatalf("QueryLazy(%d,%d) = %d, want %d", c.l, c.r, got, c.want)
		}
	}
}

// 计数线段树：按值域区间统计已插入元素个数（327/493/315/1649 的用法）
func Test_SegmentCountTree(t *testing.T) {
	st := SegmentCountTree{}
	st.Init([]int{1, 2, 3, 4, 5}, sum) // 有序去重的值域

	for _, v := range []int{3, 3, 5, 1} { // 插入 3,3,5,1
		st.UpdateCount(v)
	}

	cases := []struct {
		lo, hi, want int // 统计值在 [lo,hi] 的个数
	}{
		{1, 5, 4},
		{3, 3, 2},
		{3, 5, 3},
		{1, 2, 1},
		{4, 5, 1},
		{2, 2, 0},
	}
	for _, c := range cases {
		if got := st.Query(c.lo, c.hi); got != c.want {
			t.Fatalf("Query(%d,%d) = %d, want %d", c.lo, c.hi, got, c.want)
		}
	}
}
