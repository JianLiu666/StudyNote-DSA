'''
Find All Numbers Disappeared in an Array
    Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:
    Input: nums = [4,3,2,7,8,2,3,1]
    Output: [5,6]

Example 2:
    Input: nums = [1,1]
    Output: [2]
 
Constraints:
    n == nums.length
    1 <= n <= 10^5
    1 <= nums[i] <= n
'''

from typing import List

class Solution:
    # My Solution
    # Time Complexity: O(3n)
    def findDisappearedNumbers_sol1(self, nums: List[int]) -> List[int]:
        for i in range(len(nums)):
            if nums[i] == i+1:
                continue

            current, value = i, nums[i]
            while current+1 != value:
                current = value-1
                nums[value-1], value = value, nums[value-1]
                # value, nums[value-1] = nums[value-1], value

        result = []
        for i in range(len(nums)):
            if nums[i] == i+1:
                continue
            result.append(i+1)
        
        return result

    # My Solution
    # Time Complexity: O(n)
    def findDisappearedNumbers_sol2(self, nums: List[int]) -> List[int]:
        for num in nums:
            offset = abs(num) - 1
            if nums[offset] > 0:
                nums[offset] *= -1
        
        result = []
        for i in range(len(nums)):
            if nums[i] > 0:
                result.append(i+1)

        return result

if __name__ == '__main__':
    s = Solution()

    assert s.findDisappearedNumbers_sol1([4,3,2,7,8,2,3,1]) == [5,6]
    assert s.findDisappearedNumbers_sol1([1,1]) == [2]

    assert s.findDisappearedNumbers_sol2([4,3,2,7,8,2,3,1]) == [5,6]
    assert s.findDisappearedNumbers_sol2([1,1]) == [2]
