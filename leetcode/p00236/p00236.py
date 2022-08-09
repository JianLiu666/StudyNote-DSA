class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.ancestor = None
    
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        def helper(node: 'TreeNode') -> bool:
            if not node:
                return False
            
            left = helper(node.left)
            right = helper(node.right)
            mid = (node == p or node == q)

            if (mid and (left or right)) or (left and right):
                self.ancestor = node
            
            return mid or left or right

        helper(root)
        return self.ancestor
