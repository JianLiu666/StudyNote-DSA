package p00297

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueNode struct {
	Val  *TreeNode
	Next *QueueNode
	Prev *QueueNode
}

func CreateQueueNode(node *TreeNode) *QueueNode {
	return &QueueNode{
		Val:  node,
		Next: nil,
		Prev: nil,
	}
}

type Codec struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func Constructor() Codec {
	return Codec{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (c *Codec) put(node *TreeNode) {
	queueNode := CreateQueueNode(node)
	if c.size == 0 {
		c.head = queueNode
		c.tail = queueNode
	} else {
		queueNode.Next = c.head
		c.head.Prev = queueNode
		c.head = queueNode
	}
	c.size++
}

func (c *Codec) pop() *TreeNode {
	if c.size == 0 {
		return nil
	}

	node := c.tail
	if c.size == 1 {
		c.head = nil
		c.tail = nil
	} else {
		c.tail = c.tail.Prev
		c.tail.Next = nil
	}
	node.Prev = nil
	c.size--

	return node.Val
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	this.put(root)

	var res strings.Builder
	for this.size > 0 {
		node := this.pop()
		if node == nil {
			res.WriteString("null,")
			continue
		}
		res.WriteString(strconv.Itoa(node.Val))
		res.WriteString(",")
		this.put(node.Left)
		this.put(node.Right)
	}

	str := res.String()
	return str[:len(str)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	strs := strings.Split(data, ",")
	num, _ := strconv.Atoi(strs[0])
	root := &TreeNode{Val: num}

	this.put(root)

	strIdx := 1
	for strIdx < len(strs) {
		node := this.pop()
		if strs[strIdx] != "null" {
			num, _ := strconv.Atoi(strs[strIdx])
			node.Left = &TreeNode{Val: num}
			this.put(node.Left)
		}
		strIdx++
		if strs[strIdx] != "null" {
			num, _ := strconv.Atoi(strs[strIdx])
			node.Right = &TreeNode{Val: num}
			this.put(node.Right)
		}
		strIdx++

	}

	return root
}
