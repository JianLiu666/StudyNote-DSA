'''
Find Numbers with Even Number of Digits
    Given an array nums of integers, return how many of them contain an even number of digits.

Example 1:
    Input: nums = [12,345,2,6,7896]
    Output: 2
    Explanation: 
        12 contains 2 digits (even number of digits). 
        345 contains 3 digits (odd number of digits). 
        2 contains 1 digit (odd number of digits). 
        6 contains 1 digit (odd number of digits). 
        7896 contains 4 digits (even number of digits). 
        Therefore only 12 and 7896 contain an even number of digits.

Example 2:
    Input: nums = [555,901,482,1771]
    Output: 1 
    Explanation: 
        Only 1771 contains an even number of digits.
 
Constraints:
    1 <= nums.length <= 500
    1 <= nums[i] <= 10^5
'''

from typing import List

class Solution:

    # My Solution
    # Time Complexity: O(n)
    def findNumbers_sol1(self, nums: List[int]) -> int:
        count = 0
        for num in nums:
            count_digits = 0
            while num != 0:
                num = num//10
                count_digits += 1
                
            if count_digits % 2 == 0:
                count += 1
            
        return count

    # LeetCode Solution
    # Time Complexity: O(n)
    def findNumbers_sol2(self, nums: List[int]) -> int:
        num_count = 0
        for num in nums:
            if len(str(num))%2 == 0:
                num_count += 1

        return num_count

if __name__ == '__main__':
    s = Solution()
    
    assert s.findNumbers_sol1([12, 345, 2, 6, 7896]) == 2
    assert s.findNumbers_sol1([555,901,482,1771]) == 1

    assert s.findNumbers_sol2([12, 345, 2, 6, 7896]) == 2
    assert s.findNumbers_sol2([555,901,482,1771]) == 1