from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        memo = [{} for _ in range(len(nums))]
        return self.dfs(memo, nums, 0, 0, target)

    def dfs(self, memo: dict, nums: List[int], idx: int, sum: int, target:int) -> int:
        if idx == len(nums):
            return 1 if sum == target else 0
        
        if sum in memo[idx]:
            return memo[idx][sum]

        sum1 = self.dfs(memo, nums, idx+1, sum+nums[idx], target)
        sum2 = self.dfs(memo, nums, idx+1, sum-nums[idx], target)

        memo[idx][sum] = sum1 + sum2
        return sum1 + sum2

if __name__ == '__main__':
    s = Solution()

    assert s.findTargetSumWays([1,1,1,1,1], 3) == 5
    assert s.findTargetSumWays([1], 1) == 1