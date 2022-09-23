package p00147

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n^2), where n is the length of linked list
// Space Complexity: O(1)
func insertionSortList(head *ListNode) *ListNode {
	// 定義一個 dummy head 省去之後在插入至 head 時的例外判斷
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	// 指針紀錄已經排序完畢後的末端位置
	sorted := head
	// 指針紀錄目前準備要排序的位置
	unsorted := head.Next

	for unsorted != nil {
		tmp := unsorted.Next

		// 定義兩個指針從頭開始尋訪, 直到找到適合的插入點為止
		previous := dummy
		cursor := dummy.Next
		for cursor.Val < unsorted.Val && cursor != unsorted {
			previous, cursor = cursor, cursor.Next
		}

		// 表示 unsorted 為目前最大值不需要往前插入, 將 sorted 更新到當前位置
		if cursor == unsorted {
			sorted, unsorted = cursor, tmp
			continue
		}

		// 節點移動流程
		sorted.Next = tmp
		previous.Next = unsorted
		unsorted.Next = cursor
		unsorted = tmp
	}

	return dummy.Next
}
