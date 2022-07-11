from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        first = head
        second = head

        for _ in range(n):
            first = first.next
        
        while first and first.next:
            first = first.next
            second = second.next

        if first == None and second == head:
            head = head.next
        else:
            second.next = second.next.next
        
        return head

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
    assert getVals(s.removeNthFromEnd(list[0], 2)) == [1,2,3,5]

    list = [ListNode(1)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.removeNthFromEnd(list[0], 1)) == []

    list = [ListNode(1), ListNode(2)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.removeNthFromEnd(list[0], 1)) == [1]

    list = [ListNode(1), ListNode(2)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.removeNthFromEnd(list[0], 2)) == [2]