from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        mid = len(nums)//2
        root = TreeNode(nums[mid])
        if mid > 0:
            root.left = self.sortedArrayToBST(nums[:mid])
        if mid+1 < len(nums):
            root.right = self.sortedArrayToBST(nums[mid+1:])
        
        return root