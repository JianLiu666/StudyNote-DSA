from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def findPeakElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0
        
        left = 0
        right = len(nums)-1

        while True:
            mid = (left+right) // 2

            if nums[mid] > nums[mid+1]:
                if mid == 0 or nums[mid] > nums[mid-1]:
                    return mid
                else:
                    right = mid - 1
            else:
                if mid + 1 == len(nums)-1:
                    return mid + 1
                else:
                    left = mid + 1

if __name__ == '__main__':
    s = Solution()

    assert s.findPeakElement([1,2,3,1]) == 2
    assert s.findPeakElement([1,2,1,3,5,6,4]) == 5