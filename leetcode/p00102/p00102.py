from typing import Optional, List
from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        
        q = deque([root])

        res = []
        while q:
            size = len(q)
            levels = []
            for _ in range(size):
                node = q.popleft()
                levels.append(node.val)
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            res.append(levels)

        return res
