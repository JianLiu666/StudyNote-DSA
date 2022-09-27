package p00138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	res := &Node{}

	origin, copyright := head, res
	mapping := map[*Node]*Node{}

	// 第一次遍歷 head 只處理兩件事情
	// 1. 複製出另一份與 head 一模一樣長度與 value 的 linked list
	// 2. 紀錄 linked list 長度中相同位置的新舊節點對照表
	for origin != nil {
		copyright.Val = origin.Val
		mapping[origin] = copyright

		// 末端位置邊界判斷
		if origin.Next != nil {
			copyright.Next = &Node{}
		}
		origin = origin.Next
		copyright = copyright.Next
	}

	// 第二次遍歷 head 按照對照表更新 linked list 的 random field
	origin, copyright = head, res
	for origin != nil {
		// 只處理確實有 random node 的節點
		if origin.Random != nil {
			copyright.Random = mapping[origin.Random]
		}
		origin = origin.Next
		copyright = copyright.Next
	}

	return res
}
