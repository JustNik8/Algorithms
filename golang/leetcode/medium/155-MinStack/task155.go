package main

func main() {

}

type elem struct {
	data     int
	minIndex int
}

type MinStack struct {
	stack    []elem
	minIndex int
}

func Constructor() MinStack {
	stack := make([]elem, 0)
	return MinStack{
		stack:    stack,
		minIndex: -1,
	}
}

func (this *MinStack) Push(val int) {
	e := elem{data: val}
	if this.minIndex == -1 {
		this.minIndex = 0
		e.minIndex = 0
		this.stack = append(this.stack, e)
		return
	}

	minElem := this.stack[this.minIndex]

	if val < minElem.data {
		e.minIndex = len(this.stack)
		this.minIndex = len(this.stack)
	} else {
		e.minIndex = this.minIndex
	}

	this.stack = append(this.stack, e)
}

func (this *MinStack) Pop() {
	size := len(this.stack)

	if this.minIndex == size-1 {
		if size-2 >= 0 {
			this.minIndex = this.stack[size-2].minIndex
		} else {
			this.minIndex = -1
		}
	}

	this.stack = this.stack[:size-1]
}

func (this *MinStack) Top() int {
	last := len(this.stack) - 1
	return this.stack[last].data
}

func (this *MinStack) GetMin() int {
	if this.minIndex == -1 {
		return -1
	}
	return this.stack[this.minIndex].data
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
