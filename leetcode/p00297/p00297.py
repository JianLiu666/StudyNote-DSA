from typing import Optional
from collections import deque

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root: Optional[TreeNode]):
        res = ""
        if not root:
            return res
        
        q = deque([root])
        while q:
            node = q.popleft()
            if not node:
                res += "null,"
                continue
            res += str(node.val) + ","
            q.extend([node.left, node.right])
        
        return res[:len(res)-1]

    def deserialize(self, data: str):
        if data == "":
            return None
        
        strs = data.split(",")
        root = TreeNode(int(strs[0]))

        q = deque([root])

        strIdx = 1
        while strIdx < len(strs):
            node = q.popleft()
            if strs[strIdx] != "null":
                node.left = TreeNode(int(strs[strIdx]))
                q.append(node.left)
            strIdx += 1
            if strs[strIdx] != "null":
                node.right = TreeNode(int(strs[strIdx]))
                q.append(node.right)
            strIdx += 1
        
        return root