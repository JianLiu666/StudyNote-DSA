from typing import Optional, List

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # Time Complexity: O(1)
    # space Complexity: O(1)
    def deleteNode(self, node):
        node.val = node.next.val
        node.next = node.next.next

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res


if __name__ == '__main__':
    s = Solution()

    list = [ListNode(4), ListNode(5), ListNode(1), ListNode(9)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    s.deleteNode(list[1])
    assert getVals(list[0]) == [4,1,9]

    list = [ListNode(4), ListNode(5), ListNode(1), ListNode(9)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    s.deleteNode(list[2])
    assert getVals(list[0]) == [4,5,9]