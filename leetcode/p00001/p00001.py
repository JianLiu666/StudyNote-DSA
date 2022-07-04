from typing import List

class Solution:

    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}

        for idx, val in enumerate(nums):
            remaining = target - val
            if remaining in seen:
                return [seen[remaining], idx]

            seen[val] = idx

        return None

if __name__ == '__main__':
    s = Solution()

    assert s.twoSum([2, 7, 11, 15], 9) == [0, 1]
    assert s.twoSum([3, 2, 4], 6) == [1, 2]
    assert s.twoSum([3, 3], 6) == [0, 1]
