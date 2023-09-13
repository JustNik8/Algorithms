package main

import (
	"container/heap"
	"fmt"
)

func main() {

}

type NextTask struct {
	name     byte
	nextTime int
}

type Task struct {
	name byte
	cnt  int
}

type TaskHeap []Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].cnt > h[j].cnt }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x any)        { *h = append(*h, x.(Task)) }
func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	cntMap := make(map[byte]int)
	for _, task := range tasks {
		cntMap[task] += 1
	}

	tmpHeap := make(TaskHeap, 0)
	for name, cnt := range cntMap {
		tmpHeap = append(tmpHeap, Task{name: name, cnt: cnt})
	}

	maxHeap := &tmpHeap
	heap.Init(maxHeap)
	q := make([]NextTask, 0)

	time := 0
	for len(*maxHeap) > 0 || len(q) > 0 {
		if len(q) > 0 {
			task := q[0]
			if task.nextTime < time {
				heap.Push(maxHeap, Task{task.name, cntMap[task.name]})
				q = q[1:]
			}
		}

		if len(*maxHeap) > 0 {
			task := heap.Pop(maxHeap).(Task)
			cntMap[task.name] -= 1

			if cntMap[task.name] > 0 {
				q = append(q, NextTask{name: task.name, nextTime: time + n})
			}
		}

		time++
	}

	fmt.Println()
	return time
}
