package leetcode

//Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	nextEl int
	hasEl  bool
	iter   *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.hasEl {
		return true
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.hasEl {
		this.hasEl = false
		return this.nextEl
	} else {
		return this.iter.next()
	}
}

func (this *PeekingIterator) peek() int {
	if this.hasEl {
		return this.nextEl
	}
	this.hasEl = true
	this.nextEl = this.iter.next()
	return this.nextEl
}
