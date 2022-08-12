from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True
        
        def helper(node: Optional[TreeNode]) -> int:
            if not node:
                return 0
            
            ld = helper(node.left)
            rd = helper(node.right)
            if ld == -1 or rd == -1 or abs(ld-rd) > 1:
                return -1
            
            return max(ld, rd) + 1
        
        return helper(root) != -1