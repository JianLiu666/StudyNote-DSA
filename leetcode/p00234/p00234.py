from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next =next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        if not head.next:
            return True
        
        slow = head
        fast = head

        while fast and fast.next:
            fast = fast.next.next
            if not fast:
                tmp = slow.next
                slow.next = None
                slow = tmp
            else:
                slow = slow.next
        
        previous = slow
        current = previous.next
        previous.next = None

        while current:
            tmp = current.next
            current.next = previous
            previous = current
            current = tmp
        
        tail = previous
        while head and tail:
            if head.val != tail.val:
                return False
            
            head = head.next
            tail = tail.next
        
        return True

if __name__ == '__main__':
    s = Solution()

    list = [ListNode(1), ListNode(2), ListNode(2), ListNode(1)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert s.isPalindrome(list[0]) == True

    list = [ListNode(1), ListNode(2)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert s.isPalindrome(list[0]) == False