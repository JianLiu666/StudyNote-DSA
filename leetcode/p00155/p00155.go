package p00155

type MinStack struct {
	arr []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		arr: make([]int, 0),
		min: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)

	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if this.arr[len(this.arr)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
