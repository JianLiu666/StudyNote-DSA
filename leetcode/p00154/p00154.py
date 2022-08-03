from typing import List

class Solution:
    # Time Complexity: average is O(logn), the worst case is O(n)
    # Space Complexity: O(1)
    def findMin(self, nums: List[int]) -> int:
        if nums[-1] > nums[0]:
            return nums[0]

        low, high = 0, len(nums)-1
        while low < high:
            mid = low + (high-low) // 2
            if nums[mid] > nums[high]:
                low = mid + 1
            elif nums[mid] == nums[high]:
                while mid < high and nums[mid] == nums[high]:
                    high -= 1
            else:
                high = mid
        
        return nums[low]

if __name__ == '__main__':
    s = Solution()

    assert s.findMin([3,4,5,1,2]) == 1
    assert s.findMin([4,5,6,7,0,1,2]) == 0
    assert s.findMin([11,13,15,17]) == 11
    assert s.findMin([3,1,3]) == 1
    assert s.findMin([3,3,3,4,1,3,3,3]) == 1
    assert s.findMin([3,3,3,3,3,3,3,3,1,3,3,3]) == 1
    assert s.findMin([3,1,3,3,3,3,3,3,3,3,3,3]) == 1
