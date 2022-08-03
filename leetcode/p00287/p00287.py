from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def findDuplicate(self, nums: List[int]) -> int:
        slow, fast = nums[0], nums[nums[0]]
        
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]
        
        anchor = 0
        while slow != anchor:
            slow = nums[slow]
            anchor = nums[anchor]
        
        return anchor

if __name__ == "__main__":
    s = Solution()

    assert s.findDuplicate([1,3,4,2,2]) == 2
    assert s.findDuplicate([3,1,3,4,2]) == 3