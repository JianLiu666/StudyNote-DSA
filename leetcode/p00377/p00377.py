from math import comb
from typing import List

class Solution:
    # Time Complexity: O(nm), where n is the length of nums, and m is the number of target
    # Space Complexity: O(m), where m is the number of target
    def combinationSum4(self, nums: List[int], target: int) -> int:
        memo = [0 for _ in range(target+1)]
        memo[0] = 1

        for i in range(1, target+1, 1):
            for n in nums:
                if i-n >= 0:
                    memo[i] += memo[i-n]
        
        return memo[target]

if __name__ == '__main__':
    s = Solution()

    assert s.combinationSum4([1,2,3], 4) == 7
    assert s.combinationSum4([9], 3) == 0