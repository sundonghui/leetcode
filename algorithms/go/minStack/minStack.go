package minstack

import "math"

type MinStack struct {
    minStack []int
    stack []int
}


func Constructor() MinStack {
    return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }
    this.stack = this.stack[:len(this.stack)-1]
    this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.minStack) == 1 {
        return 0
    }
    return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */