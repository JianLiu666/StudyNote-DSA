from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        if not root:
            return res
        
        visited = set()
        stack = [root]
        while len(stack) > 0:
            node = stack.pop()

            if node.left and node.left not in visited:
                stack.append(node)
                stack.append(node.left)
            else:
                res.append(node.val)
                visited.add(node)
                if node.right and node.right not in visited:
                    stack.append(node.right)
        
        return res

