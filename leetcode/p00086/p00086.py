from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        left_head = ListNode(0)
        left_cursor = left_head

        right_head = ListNode(0)
        right_cursor = right_head

        while head:
            if head.val < x:
                left_cursor.next = head
                left_cursor = left_cursor.next
            else:
                right_cursor.next = head
                right_cursor = right_cursor.next
            
            head = head.next
        
        left_cursor.next = right_head.next
        right_cursor.next = None

        return left_head.next

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res


if __name__ == '__main__':
    s = Solution()

    list = [ListNode(1), ListNode(4), ListNode(3), ListNode(2), ListNode(5), ListNode(2)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.partition(list[0], 3)) == [1,2,2,4,3,5]

    list = [ListNode(2), ListNode(1)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert getVals(s.partition(list[0], 2)) == [1,2]