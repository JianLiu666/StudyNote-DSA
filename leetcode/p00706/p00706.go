package p00706

type HashMap_BruteForce struct {
	arr []int
}

func Constructor() HashMap_BruteForce {
	arr := make([]int, 1000001)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}

	return HashMap_BruteForce{
		arr: arr,
	}
}

func (this *HashMap_BruteForce) Put(key int, value int) {
	this.arr[key] = value
}

func (this *HashMap_BruteForce) Get(key int) int {
	return this.arr[key]
}

func (this *HashMap_BruteForce) Remove(key int) {
	this.arr[key] = -1
}
