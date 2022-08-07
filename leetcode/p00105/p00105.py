import enum
from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        self.rootIdx, memo = 0, {val: idx for idx, val in enumerate(inorder)}

        def helper(startIdx: int, endIdx: int) -> TreeNode:
            if startIdx > endIdx:
                 return None
            
            root = TreeNode(preorder[self.rootIdx])
            self.rootIdx += 1
            root.left = helper(startIdx, memo[root.val]-1)
            root.right = helper(memo[root.val]+1, endIdx)
            
            return root
        
        return helper(0, len(inorder)-1)