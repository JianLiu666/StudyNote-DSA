from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def search(self, nums: List[int], target: int) -> int:
        left = 0
        right = len(nums)-1

        while left <= right:
            mid = left + (right-left) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                right = mid - 1
            else:
                left = mid + 1

        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.search([-1,0,3,5,9,12], 9) == 4
    assert s.search([-1,0,3,5,9,12], 2) == -1