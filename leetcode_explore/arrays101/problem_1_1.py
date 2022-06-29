'''
Max Consecutive Ones
    Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:
    Input: nums = [1,1,0,1,1,1]
    Output: 3
    Explanation:
        The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

Example 2:
    Input: nums = [1,0,1,1,0,1]
    Output: 2
 
Constraints:
    1 <= nums.length <= 10^5
    nums[i] is either 0 or 1.
'''

'''
解題方向:
    - For Loop 一遍沒有異議 -> O(n)
    - 用兩個變數紀錄 local_maximum 與 global_maximum
    - 遇到 0 的時候重置 local_maximum
    - 每次計數後都確認一次 local_maximum 是否大於 global_maximum
'''

from typing import List

# My Solution
# Time Complexity: O(n)
class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        local_maximum, global_maximum = 0, 0
        for num in nums:
            if num == 0:
                local_maximum = 0
                continue

            local_maximum += 1
            global_maximum = max(local_maximum, global_maximum)
        
        return global_maximum

if __name__ == '__main__':
    s = Solution()
    
    s.findMaxConsecutiveOnes([1,1,0,1,1,1]) == 3
    s.findMaxConsecutiveOnes([1,0,1,1,0,1]) == 2