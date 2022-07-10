from typing import Optional

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if head == None:
            return False
        
        slow = head
        fast = head
        while fast.next != None:
            slow = slow.next
            fast = fast.next
            if fast.next == None:
                return False
            
            fast = fast.next
            if slow.next == fast.next:
                return True

        return False

if __name__ == '__main__':
    s = Solution()

    list = [ListNode(3), ListNode(2), ListNode(0), ListNode(-4)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    list[-1].next = list[1]
    assert s.hasCycle(list[0]) == True

    list = [ListNode(3), ListNode(2), ListNode(0), ListNode(-4)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    list[-1].next = list[0]
    assert s.hasCycle(list[0]) == True

    list = [ListNode(1)]
    assert s.hasCycle(list[0]) == False

    list = [ListNode(1), ListNode(1), ListNode(1), ListNode(1)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert s.hasCycle(list[0]) == False