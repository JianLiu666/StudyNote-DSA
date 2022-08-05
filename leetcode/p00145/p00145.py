from typing import Optional, List
from collections import defaultdict

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        if not root:
            return res
        
        stack = [root]
        memo = defaultdict(bool)

        while len(stack) > 0:
            node = stack.pop()

            if (not node.left or memo[node.left]) and (not node.right or memo[node.right]):
                memo[node] = True
                res.append(node.val)
            else:
                stack.append(node)
                if node.left and memo[node.left] == False:
                    stack.append(node.left)
                elif node.right and memo[node.right] == False:
                    stack.append(node.right)
        
        return res