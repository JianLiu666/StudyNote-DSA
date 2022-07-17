from typing import List

class Solution:
    # Time Complexity: O(nlogn)
    # Space Complexity: O(n)
    def arrayPairSum(self, nums: List[int]) -> int:
        return sum(sorted(nums)[::2])

if __name__ == '__main__':
    s = Solution()

    assert s.arrayPairSum([1,4,2,3]) == 4
    assert s.arrayPairSum([6,2,6,5,1,2]) == 9