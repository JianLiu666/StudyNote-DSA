from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None
        
        size = 0
        current = head
        while current:
            size += 1
            current = current.next
        
        k %= size

        tail = head
        for _ in range(k):
            tail = tail.next
        
        tmp = head
        while tail.next:
            tail = tail.next
            tmp = tmp.next
        
        tail.next = head
        new_head = tmp.next
        tmp.next = None

        return new_head

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res


if __name__ == '__main__':
    s = Solution()

    list = [ListNode(1), ListNode(2), ListNode(3), ListNode(4), ListNode(5)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.rotateRight(list[0], 2)) == [4,5,1,2,3]

    list = [ListNode(0), ListNode(1), ListNode(2)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.rotateRight(list[0], 4)) == [2,0,1]
