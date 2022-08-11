package p00173

// type BSTIterator struct {
// 	stack  Stack
// 	memo   map[*TreeNode]bool
// 	cursor *TreeNode
// }

// func Constructor(root *TreeNode) BSTIterator {
// 	iterator := BSTIterator{
// 		stack:  CreateStack(),
// 		memo:   map[*TreeNode]bool{},
// 		cursor: nil,
// 	}

// 	iterator.stack.Put(root)
// 	for {
// 		node := iterator.stack.Pop()
// 		if !iterator.memo[node] {
// 			iterator.memo[node] = true
// 			if node.Right != nil {
// 				iterator.stack.Put(node.Right)
// 			}
// 			if node.Left != nil {
// 				iterator.stack.Put(node)
// 				iterator.stack.Put(node.Left)
// 			}
// 		}
// 		if node.Left == nil || iterator.memo[node.Left] {
// 			iterator.cursor = node
// 			break
// 		}
// 	}

// 	return iterator
// }

// func (this *BSTIterator) Next() int {
// 	prev := this.cursor

// 	this.cursor = nil
// 	for !this.stack.Empty() {
// 		node := this.stack.Pop()
// 		if !this.memo[node] {
// 			this.memo[node] = true
// 			if node.Right != nil {
// 				this.stack.Put(node.Right)
// 			}
// 			if node.Left != nil {
// 				this.stack.Put(node)
// 				this.stack.Put(node.Left)
// 			}
// 		}
// 		if node.Left == nil || this.memo[node.Left] {
// 			this.cursor = node
// 			break
// 		}
// 	}

// 	return prev.Val
// }

// func (this *BSTIterator) HasNext() bool {
// 	return this.cursor != nil
// }
