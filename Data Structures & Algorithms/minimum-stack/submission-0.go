type MinStack struct {
	data []int
	min  int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val - this.min)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if val < 0 {
		this.min -= val
	}
}

func (this *MinStack) Top() int {
	val := this.data[len(this.data)-1]
	if val < 0 {
		return this.min
	}
	return val + this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
