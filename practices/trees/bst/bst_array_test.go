package bst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	ast := assert.New(t)

	nums := []int{5, 2, 6, 1, 4, 7, 3}
	bst := CreateBSTArray(nums)

	ast.Equal("5,2,6,1,4,x,7,x,x,3,x,x,x,x", bst.String())
	ast.Equal(2, bst.Search(2))
	ast.Equal(7, bst.Search(7))
	ast.Equal(0, bst.Search(10))

	bst.Add(12)
	bst.Add(15)
	bst.Add(10)
	fmt.Println(bst.String())
	ast.Equal("5,2,6,1,4,x,7,x,x,3,x,x,x,x,12,x,x,x,x,x,x,x,x,x,x,x,x,x,x,10,15,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x", bst.String())
}
