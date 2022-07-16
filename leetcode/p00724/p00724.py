from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def pivotIndex(self, nums: List[int]) -> int:
        left = 0
        right = 0
        for n in nums:
            right += n
        
        for i in range(len(nums)):
            right -= nums[i]
            if left == right:
                return i
            
            left += nums[i]
        
        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.pivotIndex([1,7,3,6,5,6]) == 3
    assert s.pivotIndex([1,2,3]) == -1
    assert s.pivotIndex([2,1,-1]) == 0