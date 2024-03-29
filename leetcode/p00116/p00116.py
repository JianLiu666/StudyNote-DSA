from typing import Optional
from collections import deque

class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def connect_bfs(self, root: 'Optional[Node]') -> 'Optional[Node]':

        def putNodes(q: deque, node: 'Optional[Node]'):
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        
        if not root:
            return root
        
        q = deque()
        putNodes(q, root)
        
        while q:
            size = len(q)
            left = q.popleft()
            putNodes(q, left)
            
            for _ in range(size-1):
                right = q.popleft()
                left.next = right
                putNodes(q, right)
                left = right
        
        return root

    # Time Complexity: O(n)
    # Space Complexity: O(n), where n is the implicit call stack
    def connect_recursion(self, root: 'Optional[Node]') -> 'Optional[Node]':
        def helper(left: 'Optional[Node]', right: 'Optional[Node]'):
            if not left or not right:
                return
            
            left.next = right
            helper(left.left, left.right)
            helper(left.right, right.left)
            helper(right.left, right.right)

        if not root:
            return root
        
        helper(root.left, root.right)
        return root