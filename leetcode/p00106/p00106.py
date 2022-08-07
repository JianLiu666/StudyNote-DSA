from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        self.rootIdx, memo = len(postorder)-1, {val: idx for idx, val in enumerate(inorder)}

        def helper(startIdx: int, endIdx: int) -> TreeNode:
            if startIdx > endIdx:
                return None
            
            root = TreeNode(postorder[self.rootIdx])
            self.rootIdx -= 1
            root.right = helper(memo[root.val]+1, endIdx)
            root.left = helper(startIdx, memo[root.val]-1)

            return root
        
        return helper(0, len(postorder)-1)
