'''
Third Maximum Number
    Given an integer array nums, return the third distinct maximum number in this array.
    If the third maximum does not exist, return the maximum number.

Example 1:
    Input: nums = [3,2,1]
    Output: 1
    Explanation:
        The first distinct maximum is 3.
        The second distinct maximum is 2.
        The third distinct maximum is 1.

Example 2:
    Input: nums = [1,2]
    Output: 2
    Explanation:
        The first distinct maximum is 2.
        The second distinct maximum is 1.
        The third distinct maximum does not exist, so the maximum (2) is returned instead.

Example 3:
    Input: nums = [2,2,3,1]
    Output: 1
    Explanation:
        The first distinct maximum is 3.
        The second distinct maximum is 2 (both 2's are counted together since they have the same value).
        The third distinct maximum is 1.
 
Constraints:
    1 <= nums.length <= 10^4
    -2^31 <= nums[i] <= 2^31 - 1
'''

from typing import List

# LeetCode solution
# Time Complexity: O(n)
class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        tier1, tier2, tier3 = float("-inf"), float("-inf"), float("-inf")
        
        for num in nums:
            if num > tier1:
                tier1, tier2, tier3 = num, tier1, tier2
            if num > tier2 and num != tier1:
                tier2, tier3 = num, tier2
            if num > tier3 and num != tier1 and num != tier2:
                tier3 = num
        
        return tier3 if tier3 != float("-inf") else tier1

if __name__ == '__main__':
    s = Solution()

    assert s.thirdMax([3,2,1]) == 1
    assert s.thirdMax([1,2]) == 2
    assert s.thirdMax([2,2,3,1]) == 1