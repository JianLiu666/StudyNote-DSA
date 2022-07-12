from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        while head and head.val == val:
            head = head.next
        
        if head == None:
            return None
        
        current = head
        while current.next:
            if current.next.val == val:
                current.next = current.next.next
            else:
                current = current.next
        
        return head

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res

if __name__ == '__main__':
    s = Solution()
    
    # Example 1
    list = [ListNode(1), ListNode(2), ListNode(6), ListNode(3), ListNode(4), ListNode(5), ListNode(6)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.removeElements(list[0], 6)) == [1,2,3,4,5]

    # Example 2
    assert getVals(s.removeElements(None, 1)) == []

    # Example 3
    list = [ListNode(7), ListNode(7), ListNode(7), ListNode(7)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.removeElements(list[0], 7)) == []