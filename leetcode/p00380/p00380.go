package p00380

import "math/rand"

type RandomizedSet struct {
	lookup map[int]int
	nums   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		lookup: map[int]int{},
		nums:   []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.lookup[val]; exists {
		return false
	}

	this.nums = append(this.nums, val)
	this.lookup[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.lookup[val]; !exists {
		return false
	}

	val_idx := this.lookup[val]
	last_val := this.nums[len(this.nums)-1]
	this.lookup[last_val] = val_idx
	this.nums[val_idx] = last_val

	this.nums = this.nums[:len(this.nums)-1]
	delete(this.lookup, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
