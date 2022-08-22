package bst

type node struct {
	Val   int
	Left  *node
	Right *node
}

func CreateNode(val int) *node {
	return &node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
