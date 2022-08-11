from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class BSTIterator:
    def __init__(self, root: Optional[TreeNode]):
        self.stack = []
        self.helper(root)

    def next(self) -> int:
        node = self.stack.pop()
        self.helper(node.right)
        return node.val

    def hasNext(self) -> bool:
        return self.stack
    
    def helper(self, node: Optional[TreeNode]):
        while node:
            self.stack.append(node)
            node = node.left
