package deque

import (
	. "github.com/liyue201/gostl/uitls/iterator"
)

//ArrayIterator is a SortableIterator
var _ SortableIterator = (*DequeIterator)(nil)

type DequeIterator struct {
	dq       *Deque
	position int
}

func (this *DequeIterator) IsValid() bool {
	if this.position >= 0 && this.position < this.dq.Size() {
		return true
	}
	return false
}

func (this *DequeIterator) Value() interface{} {
	return this.dq.At(this.position)
}

func (this *DequeIterator) SetValue(val interface{}) error {
	return this.dq.Set(this.position, val)
}

func (this *DequeIterator) Next() ConstIterator {
	if this.position < this.dq.Size() {
		this.position++
	}
	return this
}

func (this *DequeIterator) Prev() ConstBidIterator {
	if this.position >= 0 {
		this.position--
	}
	return this
}

func (this *DequeIterator) Clone() interface{} {
	return &DequeIterator{dq: this.dq, position: this.position}
}

func (this *DequeIterator) IteratorAt(position int) SortableIterator {
	return &DequeIterator{dq: this.dq, position: position}
}

func (this *DequeIterator) Position() int {
	return this.position
}
