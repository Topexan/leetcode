package leetcode

type CustomStack struct {
	stack   []int
	size    int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	res := CustomStack{}
	res.stack = make([]int, maxSize)
	res.size = 0
	res.maxSize = maxSize
	return res
}

func (this *CustomStack) Push(x int) {
	if this.size == this.maxSize {
		return
	}
	this.stack[this.size] = x
	this.size++
}

func (this *CustomStack) Pop() int {
	if this.size == 0 {
		return -1
	}
	this.size--
	return this.stack[this.size]
}

func (this *CustomStack) Increment(k int, val int) {
	var l int
	if k < this.size {
		l = k
	} else {
		l = this.size
	}
	for i := 0; i < l; i++ {
		this.stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
