package p00895

type FreqStack struct {
	freq    map[int]int
	stack   map[int][]int
	maxfreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:    map[int]int{},
		stack:   map[int][]int{},
		maxfreq: 0,
	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++
	this.stack[this.freq[val]] = append(this.stack[this.freq[val]], val)
	this.maxfreq = max(this.maxfreq, this.freq[val])
}

func (this *FreqStack) Pop() int {
	l := len(this.stack[this.maxfreq])
	val := this.stack[this.maxfreq][l-1]
	this.stack[this.maxfreq] = this.stack[this.maxfreq][:l-1]
	this.freq[val]--

	if len(this.stack[this.maxfreq]) == 0 {
		this.maxfreq--
	}

	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
