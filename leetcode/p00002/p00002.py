from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = current = ListNode(0)

        carried = 0
        while l1 or l2:
            if l1:
                carried += l1.val
                l1 = l1.next
            if l2:
                carried += l2.val
                l2 = l2.next
            
            carried, val = divmod(carried, 10)
            current.next = ListNode(val)
            current = current.next
        
        if carried > 0:
            current.next = ListNode(carried)
        
        return dummy.next

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res


if __name__ == '__main__':
    s = Solution()

    # Example 1
    list1 = [ListNode(2), ListNode(4), ListNode(3)]
    for i in range(0, len(list1)-1):
        list1[i].next = list1[i+1]
    list2 = [ListNode(5), ListNode(6), ListNode(4)]
    for i in range(0, len(list2)-1):
        list2[i].next = list2[i+1]
    assert getVals(s.addTwoNumbers(list1[0], list2[0])) == [7,0,8]

    # Example 2
    list1 = [ListNode(0)]
    for i in range(0, len(list1)-1):
        list1[i].next = list1[i+1]
    list2 = [ListNode(0)]
    for i in range(0, len(list2)-1):
        list2[i].next = list2[i+1]
    assert getVals(s.addTwoNumbers(list1[0], list2[0])) == [0]

    # Example 3
    list1 = [ListNode(9), ListNode(9), ListNode(9), ListNode(9), ListNode(9), ListNode(9), ListNode(9)]
    for i in range(0, len(list1)-1):
        list1[i].next = list1[i+1]
    list2 = [ListNode(9), ListNode(9), ListNode(9), ListNode(9)]
    for i in range(0, len(list2)-1):
        list2[i].next = list2[i+1]
    assert getVals(s.addTwoNumbers(list1[0], list2[0])) == [8,9,9,9,0,0,0,1]