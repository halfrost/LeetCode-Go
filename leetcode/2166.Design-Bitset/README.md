# [2166. Design Bitset](https://leetcode.com/problems/design-bitset/)


## 题目

A **Bitset** is a data structure that compactly stores bits.

Implement the `Bitset` class:

- `Bitset(int size)` Initializes the Bitset with `size` bits, all of which are `0`.
- `void fix(int idx)` Updates the value of the bit at the index `idx` to `1`. If the value was already `1`, no change occurs.
- `void unfix(int idx)` Updates the value of the bit at the index `idx` to `0`. If the value was already `0`, no change occurs.
- `void flip()` Flips the values of each bit in the Bitset. In other words, all bits with value `0` will now have value `1` and vice versa.
- `boolean all()` Checks if the value of **each** bit in the Bitset is `1`. Returns `true` if it satisfies the condition, `false` otherwise.
- `boolean one()` Checks if there is **at least one** bit in the Bitset with value `1`. Returns `true` if it satisfies the condition, `false` otherwise.
- `int count()` Returns the **total number** of bits in the Bitset which have value `1`.
- `String toString()` Returns the current composition of the Bitset. Note that in the resultant string, the character at the `ith` index should coincide with the value at the `ith` bit of the Bitset.

**Example 1:**

```
Input
["Bitset", "fix", "fix", "flip", "all", "unfix", "flip", "one", "unfix", "count", "toString"]
[[5], [3], [1], [], [], [0], [], [], [0], [], []]
Output
[null, null, null, null, false, null, null, true, null, 2, "01010"]

Explanation
Bitset bs = new Bitset(5); // bitset = "00000".
bs.fix(3);     // the value at idx = 3 is updated to 1, so bitset = "00010".
bs.fix(1);     // the value at idx = 1 is updated to 1, so bitset = "01010".
bs.flip();     // the value of each bit is flipped, so bitset = "10101".
bs.all();      // return False, as not all values of the bitset are 1.
bs.unfix(0);   // the value at idx = 0 is updated to 0, so bitset = "00101".
bs.flip();     // the value of each bit is flipped, so bitset = "11010".
bs.one();      // return True, as there is at least 1 index with value 1.
bs.unfix(0);   // the value at idx = 0 is updated to 0, so bitset = "01010".
bs.count();    // return 2, as there are 2 bits with value 1.
bs.toString(); // return "01010", which is the composition of bitset.

```

**Constraints:**

- `1 <= size <= 10^5`
- `0 <= idx <= size - 1`
- At most `10^5` calls will be made **in total** to `fix`, `unfix`, `flip`, `all`, `one`, `count`, and `toString`.
- At least one call will be made to `all`, `one`, `count`, or `toString`.
- At most `5` calls will be made to `toString`.

## 题目大意

位集 Bitset 是一种能以紧凑形式存储位的数据结构。

请你实现 Bitset 类。

- Bitset(int size) 用 size 个位初始化 Bitset ，所有位都是 0 。
- void fix(int idx) 将下标为 idx 的位上的值更新为 1 。如果值已经是 1 ，则不会发生任何改变。
- void unfix(int idx) 将下标为 idx 的位上的值更新为 0 。如果值已经是 0 ，则不会发生任何改变。
- void flip() 翻转 Bitset 中每一位上的值。换句话说，所有值为 0 的位将会变成 1 ，反之亦然。
- boolean all() 检查 Bitset 中 每一位 的值是否都是 1 。如果满足此条件，返回 true ；否则，返回 false 。
- boolean one() 检查 Bitset 中 是否 至少一位 的值是 1 。如果满足此条件，返回 true ；否则，返回 false 。
- int count() 返回 Bitset 中值为 1 的位的 总数 。
- String toString() 返回 Bitset 的当前组成情况。注意，在结果字符串中，第 i 个下标处的字符应该与 Bitset 中的第 i 位一致。

提示：

- 1 <= size <= 10^5
- 0 <= idx <= size - 1
- 至多调用 fix、unfix、flip、all、one、count 和 toString 方法 总共 10^5 次
- 至少调用 all、one、count 或 toString 方法一次
- 至多调用 toString 方法 5 次

## 解题思路

- 题目中给出了 size 大小，10^5 位二进制。所以不能用 int64 数据类型。
- 用数组模拟二进制位的一系列操作。flip 操作并不需要每次去翻转，偶数次翻转等于没有翻转，奇数次翻转记下标记，同时更新 1 的个数。这次懒操作在调用 fix 和 unfix 时，更新到原来数组中。
- fix 和 unfix 根据懒数组中的标记对应更新二进制位。同时更新 1 的个数。
- all，one，count 都是判断 1 的个数。toString 输出即可。

## 代码

```go
package leetcode

type Bitset struct {
	set      []byte
	flipped  []byte
	oneCount int
	size     int
}

func Constructor(size int) Bitset {
	set := make([]byte, size)
	flipped := make([]byte, size)
	for i := 0; i < size; i++ {
		set[i] = byte('0')
		flipped[i] = byte('1')
	}
	return Bitset{
		set:      set,
		flipped:  flipped,
		oneCount: 0,
		size:     size,
	}
}

func (this *Bitset) Fix(idx int) {
	if this.set[idx] == byte('0') {
		this.set[idx] = byte('1')
		this.flipped[idx] = byte('0')
		this.oneCount++
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.set[idx] == byte('1') {
		this.set[idx] = byte('0')
		this.flipped[idx] = byte('1')
		this.oneCount--
	}
}

func (this *Bitset) Flip() {
	this.set, this.flipped = this.flipped, this.set
	this.oneCount = this.size - this.oneCount
}

func (this *Bitset) All() bool {
	return this.oneCount == this.size
}

func (this *Bitset) One() bool {
	return this.oneCount != 0
}

func (this *Bitset) Count() int {
	return this.oneCount
}

func (this *Bitset) ToString() string {
	return string(this.set)
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
```