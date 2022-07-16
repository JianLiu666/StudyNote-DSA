from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = head
        fast = head

        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        return slow


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
    assert getVals(s.middleNode(list[0])) == [3,4,5]

    list = [ListNode(1), ListNode(2), ListNode(3), ListNode(4), ListNode(5), ListNode(6)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.middleNode(list[0])) == [4,5,6]