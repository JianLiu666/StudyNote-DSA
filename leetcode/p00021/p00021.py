from typing import Optional, List

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if not list1:
            return list2
        if not list2:
            return list1

        head = None
        tail = None

        if list1.val <= list2.val:
            head = list1
            tail = head
            list1 = list1.next
        else:
            head = list2
            tail = head
            list2 = list2.next
        
        while list1 and list2:
            if list1.val <= list2.val:
                tail.next = list1
                tail = tail.next
                list1 = list1.next
            else:
                tail.next = list2
                tail = tail.next
                list2 = list2.next
        
        if list1:
            tail.next = list1
        else:
            tail.next = list2
        
        return head

def getVals(head: Optional[ListNode]) -> List[int]:
    res = []

    while head:
        res.append(head.val)
        head = head.next

    return res

if __name__ == '__main__':
    s = Solution()

    # Exmpale 1
    list1 = [ListNode(1), ListNode(2), ListNode(4)]
    for i in range(0, len(list1)-1):
        list1[i].next = list1[i+1]
    list2 = [ListNode(1), ListNode(3), ListNode(4)]
    for i in range(0, len(list2)-1):
        list2[i].next = list2[i+1]
    assert getVals(s.mergeTwoLists(list1[0], list2[0])) == [1,1,2,3,4,4]

    # Example 2
    assert getVals(s.mergeTwoLists(None, None)) == []

    # Example 3s
    list2 = [ListNode(0)]
    for i in range(0, len(list2)-1):
        list2[i].next = list2[i+1]
    assert getVals(s.mergeTwoLists([], list2[0])) == [0]