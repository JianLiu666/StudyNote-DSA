'''
Sort Array By Parity
    Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
    Return any array that satisfies this condition.

Example 1:
    Input: nums = [3,1,2,4]
    Output: [2,4,3,1]
    Explanation:
        The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Example 2:
    Input: nums = [0]
    Output: [0]

Constraints:
    1 <= nums.length <= 5000
    0 <= nums[i] <= 5000
'''

from typing import List

class Solution:
    # My Solution
    # Time Complexity: O(n)
    def sortArrayByParity_sol1(self, nums: List[int]) -> List[int]:
        head, tail = 0, len(nums)-1
        while head < tail:
            if nums[head] % 2 == 1 and nums[tail] % 2 == 0:
                nums[head], nums[tail] = nums[tail], nums[head]
                head += 1
                tail -= 1
            if nums[head] % 2 == 0:
                head += 1
            if nums[tail] % 2 == 1:
                tail -= 1
        
        return nums

    # LeetCode Solution
    # Time Complexity: O(n)
    def sortArrayByParity_sol2(self, nums: List[int]) -> List[int]:
        head, tail = 0, len(nums)-1
        while head < tail:
            if nums[head] % 2 == 1:
                if nums[tail] % 2 == 0:
                    nums[head], nums[tail] = nums[tail], nums[head]
                    head += 1
                tail -= 1
            else:
                if nums[tail] % 2 == 1:
                    tail -= 1
                head += 1
        
        return nums

if __name__ == '__main__':
    s = Solution()

    assert s.sortArrayByParity_sol1([3, 1, 2, 4]) == [4, 2, 1, 3]
    assert s.sortArrayByParity_sol1([0]) == [0]
    assert s.sortArrayByParity_sol2([3, 1, 2, 4]) == [4, 2, 1, 3]
    assert s.sortArrayByParity_sol2([0]) == [0]