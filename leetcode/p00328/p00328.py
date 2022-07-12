from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next =next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        while not head or not head.next:
            return head
        
        previous = head
        current = head.next
        
        while current and current.next:
            tmp = previous.next
            previous.next = current.next
            current.next = current.next.next
            previous.next.next = tmp
            previous = previous.next
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

    list = [ListNode(1), ListNode(2), ListNode(3), ListNode(4), ListNode(5)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.oddEvenList(list[0])) == [1,3,5,2,4]

    list = [ListNode(2), ListNode(1), ListNode(3), ListNode(5), ListNode(6), ListNode(4), ListNode(7)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.oddEvenList(list[0])) == [2,3,6,7,1,5,4]