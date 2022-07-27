from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def flatten(self, root: Optional[TreeNode]) -> None:
        if not root:
            return 
        
        stack = []
        self.dfs(stack, root)

        current = stack.pop()
        while len(stack) > 0:
            current.left = None
            current.right = stack.pop()
            current = current.right
    
    def dfs(self, stack: List[TreeNode], node: Optional[TreeNode]):
        if node.right:
            self.dfs(stack, node.right)
        if node.left:
            self.dfs(stack, node.left)
        stack.append(node)