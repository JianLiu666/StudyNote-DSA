'''
Move Zeroes
    Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
    Note that you must do this in-place without making a copy of the array.

Example 1:
    Input: nums = [0,1,0,3,12]
    Output: [1,3,12,0,0]

Example 2:
    Input: nums = [0]
    Output: [0]
 
Constraints:
    1 <= nums.length <= 10^4
    -2^31 <= nums[i] <= 2^31 - 1
'''

from typing import List

# My Solution
# Time Complexity: O(n)
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        offset = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[offset] = nums[i]
                offset += 1
        
        while offset < len(nums):
            nums[offset] = 0
            offset += 1

if __name__ == '__main__':
    s = Solution()

    list1 = [0, 1, 0, 3, 12]
    s.moveZeroes(list1)
    print(list1)
    
    list2 = [0]
    s.moveZeroes(list2)
    print(list2)