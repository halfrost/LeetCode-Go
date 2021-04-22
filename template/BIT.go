package template

// BinaryIndexedTree define
type BinaryIndexedTree struct {
	tree     []int
	capacity int
}

// Init define
func (bit *BinaryIndexedTree) Init(capacity int) {
	bit.tree, bit.capacity = make([]int, capacity+1), capacity
}

// Add define
func (bit *BinaryIndexedTree) Add(index int, val int) {
	for ; index <= bit.capacity; index += index & -index {
		bit.tree[index] += val
	}
}

// Query define
func (bit *BinaryIndexedTree) Query(index int) int {
	sum := 0
	for ; index > 0; index -= index & -index {
		sum += bit.tree[index]
	}
	return sum
}

// InitWithNums define
func (bit *BinaryIndexedTree) InitWithNums(nums []int) {
	bit.tree, bit.capacity = make([]int, len(nums)+1), len(nums)
	for i := 1; i <= len(nums); i++ {
		bit.tree[i] += nums[i-1]
		for j := i - 2; j >= i-lowbit(i); j-- {
			bit.tree[i] += nums[j]
		}
	}
}

func lowbit(x int) int {
	return x & -x
}

// BinaryIndexedTree2D define
type BinaryIndexedTree2D struct {
	tree [][]int
	row  int
	col  int
}

// Add define
func (bit2 *BinaryIndexedTree2D) Add(i, j int, val int) {
	for i <= bit2.row {
		k := j
		for k <= bit2.col {
			bit2.tree[i][k] += val
			k += lowbit(k)
		}
		i += lowbit(i)
	}
}

// Query define
func (bit2 *BinaryIndexedTree2D) Query(i, j int) int {
	sum := 0
	for i >= 1 {
		k := j
		for k >= 1 {
			sum += bit2.tree[i][k]
			k -= lowbit(k)
		}
		i -= lowbit(i)
	}
	return sum
}
