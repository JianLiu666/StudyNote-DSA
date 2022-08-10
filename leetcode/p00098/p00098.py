from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def helper(node: Optional[TreeNode]) -> List[int]:
            if not node:
                return []
            return helper(node.left) + [node.val] + helper(node.right)
        
        res = helper(root)
        for i in range(1, len(res)):
            if res[i] <= res[i-1]:
                return False
        return True