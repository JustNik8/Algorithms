package main

import "math/rand"

type RandomizedSet struct {
	indexes map[int]int
	elems   []int
}

func Constructor() RandomizedSet {
	indexes := make(map[int]int)
	elems := make([]int, 0)

	return RandomizedSet{
		indexes: indexes,
		elems:   elems,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, contains := this.indexes[val]

	if contains {
		return false
	}

	index := len(this.elems)

	this.elems = append(this.elems, val)
	this.indexes[val] = index

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, contains := this.indexes[val]

	if !contains {
		return false
	}

	last := len(this.elems) - 1
	lastElem := this.elems[last]

	this.elems[index] = lastElem
	this.elems = this.elems[:last]

	this.indexes[lastElem] = index
	delete(this.indexes, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.elems))

	return this.elems[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
