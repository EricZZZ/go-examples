package main

import "math"

type MinStack struct {
	min_x int
	items []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		items: make([]int, 0),
		min_x: math.MaxInt,
	}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	this.min_x = min(this.min_x, x)
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
	this.min_x = math.MaxInt
	for _, v := range this.items {
		this.min_x = min(this.min_x, v)
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.min_x
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
