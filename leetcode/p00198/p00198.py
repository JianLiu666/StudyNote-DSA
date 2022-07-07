from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return nums[0]
        
        pre = nums[0]
        nxt = max(nums[1], nums[0])

        for i in range(2, len(nums)):
            nxt, pre = max(nums[i]+pre, nxt), nxt
        
        return max(nxt, pre)

if __name__ == '__main__':
    s = Solution()

    assert s.rob([1,2,3,1]) == 4
    assert s.rob([2,7,9,3,1]) == 12
