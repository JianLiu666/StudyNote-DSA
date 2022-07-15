from typing import Optional

class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None
        
        memo = {}
        
        anchor = head
        root = Node(head.val)
        current = root

        memo[anchor] = root
        anchor = anchor.next

        while anchor:
            node = Node(anchor.val)
            current.next = node
            current = current.next

            memo[anchor] = current

            anchor = anchor.next
        
        anchor = head
        current = root

        while anchor:
            if anchor.random in memo:
                current.random = memo[anchor.random]
            
            anchor = anchor.next
            current = current.next
        
        return root
