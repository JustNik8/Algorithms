package main

// Идея в том , что мы пушим в стек1. А pop и peek берут элементы из стека1, и пушат в стек2.
// Таким образом, в стек2 элементы будут находится в порядке FIFO.

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		this.moveElems()
	}

	if len(this.out) == 0 {
		return 0
	}

	last := len(this.out) - 1
	elem := this.out[last]
	this.out = this.out[:last]
	return elem
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		this.moveElems()
	}

	if len(this.out) == 0 {
		return 0
	}

	last := len(this.out) - 1
	return this.out[last]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func (this *MyQueue) moveElems() {
	for len(this.in) > 0 {
		last := len(this.in) - 1
		elem := this.in[last]

		this.out = append(this.out, elem)
		this.in = this.in[:last]
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
