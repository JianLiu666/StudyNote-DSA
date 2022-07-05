from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def search(self, nums: List[int], target: int) -> int:
        left = 0
        right = len(nums)-1

        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            
            if nums[mid] >= nums[left]:
                if target < nums[left] or target > nums[mid]:
                    left = mid + 1
                else:
                    right = mid - 1
            else:
                if target < nums[mid] or target > nums[right]:
                    right = mid - 1
                else:
                    left = mid + 1

        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.search([4,5,6,7,0,1,2], 0) == 4
    assert s.search([4,5,6,7,0,1,2], 3) == -1
    assert s.search([1], 0) == -1
    assert s.search([3,1], 1) == 1