from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def findMin(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]

        left = 0
        right = len(nums)-1
        if nums[left] < nums[right]:
            return nums[0]
        
        while left != right:
            mid = (left + right) // 2
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        
        return nums[left]

if __name__ == '__main__':
    s = Solution()

    assert s.findMin([3,4,5,1,2]) == 1
    assert s.findMin([4,5,6,7,0,1,2]) == 0
    assert s.findMin([11,13,15,17]) == 11