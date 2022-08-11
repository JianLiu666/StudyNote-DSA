from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        node = TreeNode(val)
        if not root:
            return node
        
        current = root
        while True:
            if current.val > val:
                if current.left:
                    current = current.left
                else:
                    current.left = node
                    break
            else:
                if current.right:
                    current = current.right
                else:
                    current.right = node
                    break
        
        return root