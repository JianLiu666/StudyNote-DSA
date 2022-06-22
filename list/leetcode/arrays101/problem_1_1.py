'''
Max Consecutive Ones
    Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:
    Input: nums = [1,1,0,1,1,1]
    Output: 3
    Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

Example 2:
    Input: nums = [1,0,1,1,0,1]
    Output: 2
 
Constraints:
    1 <= nums.length <= 10^5
    nums[i] is either 0 or 1.
'''

from typing import List

# My Solution
# Time Complexity: O(n)
def findMaxConsecutiveOnes_sol1(nums:List[int]) -> int:
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

print(findMaxConsecutiveOnes_sol1([1,1,0,1,1,1]))