package bht

import (
	"fmt"
	"testing"
)

func TestBHTArray(t *testing.T) {
	// ast := assert.New(t)

	nums := []int{66, 78, 27, 35, 6, 2, 44, 58, 29, 76}
	bht := CreateBHTArray(nums)
	fmt.Println(bht.bht)

}
