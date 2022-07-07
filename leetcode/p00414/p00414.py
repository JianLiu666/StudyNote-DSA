from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def thirdMax(self, nums: List[int]) -> int:
        first, second, third = float("-inf"), float("-inf"), float("-inf")

        for num in nums:
            if num > first:
                first, second, third = num, first, second
            elif num > second and num != first:
                second, third = num, second
            elif num > third and num != second and num != first:
                third = num
        
        return third if third != float("-inf") else first

if __name__ == '__main__':
    s = Solution()

    assert s.thirdMax([3,2,1]) == 1
    assert s.thirdMax([1,2]) == 2
    assert s.thirdMax([2,2,3,1]) == 1