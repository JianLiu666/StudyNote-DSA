from typing import Optional
from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        
        q = deque([root])
        depth = 0
        while q:
            size = len(q)
            depth += 1
            for _ in range(size):
                node = q.popleft()
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
        
        return depth
