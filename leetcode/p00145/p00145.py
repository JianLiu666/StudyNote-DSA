from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        if not root:
            return res
        
        stack = []
        while root or len(stack) > 0:
            while root:
                res.insert(0, root.val)
                stack.append(root)
                root = root.right
            root = stack.pop()
            root = root.left
        
        return res