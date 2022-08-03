from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None
        if not head.next:
            return head

        current = head
        newHead = head.next
        current.next = current.next.next
        newHead.next = current

        while current.next and current.next.next:
            tmp = current.next
            current.next = current.next.next
            current = current.next
            tmp.next = current.next
            current.next = tmp
            current = current.next

        return newHead

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res


if __name__ == '__main__':
    s = Solution()

    list = [ListNode(1), ListNode(2), ListNode(3), ListNode(4)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.swapPairs(list[0])) == [2,1,4,3]

    assert getVals(s.swapPairs(None)) == []

    list = [ListNode(1)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.swapPairs(list[0])) == [1]