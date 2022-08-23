package bht

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBHTArray(t *testing.T) {
	ast := assert.New(t)

	nums := []int{66, 78, 27, 35, 6, 2, 44, 58, 29, 76}
	bht := CreateBHTArray(nums)

	ast.Equal("78,76,44,58,66,2,27,35,29,6", bht.String())

	bht.Remove()
	ast.Equal("76,66,44,58,6,2,27,35,29", bht.String())

	bht.Add(100)
	bht.Add(120)
	ast.Equal("120,100,44,58,76,2,27,35,29,6,66", bht.String())
}
