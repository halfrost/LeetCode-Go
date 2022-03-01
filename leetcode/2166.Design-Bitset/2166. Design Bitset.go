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
