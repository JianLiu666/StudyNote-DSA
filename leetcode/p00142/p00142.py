from typing import Optional

class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def detectCycle_hashmap(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return None
        
        memo = {}
        current = head
        
        idx = 0
        while current.next != None:
            if current in memo:
                return current
            
            memo[current] = idx
            current = current.next
            idx += 1
    
        return None

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def detectCycle_pointers(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head == None:
            return None

        slow = head
        fast = head

        while fast.next != None and fast.next.next != None:
            slow = slow.next
            fast = fast.next.next
            
            if slow == fast:
                slow2 = head
                while slow != slow2:
                    slow = slow.next
                    slow2 = slow2.next
                return slow
        
        return None

if __name__ == '__main__':
    s = Solution()

    list = [ListNode(3), ListNode(2), ListNode(0), ListNode(-4)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    list[-1].next = list[1]
    assert s.detectCycle_hashmap(list[0]) == list[1]
    assert s.detectCycle_pointers(list[0]) == list[1]

    list = [ListNode(3), ListNode(2), ListNode(0), ListNode(-4)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    list[-1].next = list[0]
    assert s.detectCycle_hashmap(list[0]) == list[0]
    assert s.detectCycle_pointers(list[0]) == list[0]

    list = [ListNode(1)]
    assert s.detectCycle_hashmap(list[0]) == None
    assert s.detectCycle_pointers(list[0]) == None

    list = [ListNode(1), ListNode(1), ListNode(1), ListNode(1)]
    for i in range(0, len(list)-1):
        list[i].next = list[i+1]
    assert s.detectCycle_hashmap(list[0]) == None
    assert s.detectCycle_pointers(list[0]) == None