from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> List[Optional[TreeNode]]:
        res = []
        seen = {}

        def dfs(node: Optional[TreeNode], s1: str) -> str:
            if node.left:
                s1 = dfs(node.left, s1)
            else:
                s1 += "n"
            
            s2 = ""
            if node.right:
                s2 = dfs(node.right, s2)
            else:
                s2 += "n"
            
            s = s1 + s2 + "0" + str(node.val)
            if s not in seen:
                seen[s] = 0
            elif s in seen and seen[s] == 0:
                res.append(node)
                seen[s] = 1
            
            return s

        dfs(root, "")
        return res

if __name__ == '__main__':
    s = Solution()

    # Example 1
    l20 = TreeNode(4, None, None)
    l10 = TreeNode(2, l20, None)
    l32 = TreeNode(4, None, None)
    l22 = TreeNode(2, l32, None)
    l23 = TreeNode(4, None, None)
    l11 = TreeNode(3, l22, l23)
    root = TreeNode(1, l10, l11)

    assert s.findDuplicateSubtrees(root) == [l32, l22]
