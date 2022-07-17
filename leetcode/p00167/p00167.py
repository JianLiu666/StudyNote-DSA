from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        l, r = 0, len(numbers)-1

        while l < r:
            num = numbers[l] + numbers[r]
            if num == target:
                return [l+1, r+1]
            if num > target:
                r -= 1
            else:
                l += 1
        
        return [-1, -1]