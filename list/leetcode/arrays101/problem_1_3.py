'''
Squares of a Sorted Array
    Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
    Input: nums = [-4,-1,0,3,10]
    Output: [0,1,9,16,100]
    Explanation: After squaring, the array becomes [16,1,0,9,100].
        After sorting, it becomes [0,1,9,16,100].

Example 2:
    Input: nums = [-7,-3,2,3,11]
    Output: [4,9,9,49,121]

Constraints:
    1 <= nums.length <= 104
    -104 <= nums[i] <= 104
    nums is sorted in non-decreasing order.
'''

from typing import List

# My Solution
# Time Complexity: O(n)
def sortedSquares(nums: List[int]) -> List[int]:
    head, tail = 0, len(nums)-1
    offset = len(nums)-1

    result = [None] * len(nums)
    while head != tail:
        if abs(nums[head]) > abs(nums[tail]):
            result[offset] = nums[head] ** 2
            head += 1
        else:
            result[offset] = nums[tail] ** 2
            tail -= 1
        offset -= 1

    result[offset] = nums[head] ** 2

    return result

print(sortedSquares([-4,-1,0,3,10]))