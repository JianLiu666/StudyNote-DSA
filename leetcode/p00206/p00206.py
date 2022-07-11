from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return None

        current = head.next
        previous = head
        previous.next = None
        
        while current:
            next = current.next
            current.next = previous
            current, previous = next, current
        
        return previous

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res

if __name__ == '__main__':
    s = Solution()

    # Example 1
    list = [ListNode(1), ListNode(2), ListNode(3), ListNode(4), ListNode(5)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.reverseList(list[0])) == [5,4,3,2,1]

    # Example 2
    list = [ListNode(1), ListNode(2)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.reverseList(list[0])) == [2,1]

    # Example 3
    assert getVals(s.reverseList(None)) == []