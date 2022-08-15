from typing import List

class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def postorder(self, root: 'Node') -> List[int]:
        res = []
        if not root:
            return res
        
        stack = [root]
        while len(stack) > 0:
            node = stack.pop()
            res.insert(0, node.val)
            for child in node.children:
                stack.append(child)
        
        return res
                