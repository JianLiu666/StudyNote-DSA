from typing import Optional

class TreeNode:
    def __ini__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def searchBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        current = root

        while current:
            if current.val == val:
                return current
            elif current.val > val:
                current = current.left
            else:
                current = current.right
        
        return None
