from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def maxSubArray(self, nums: List[int]) -> int:
        res = -10001
        cur = 0

        for num in nums:
            if cur < 0:
                cur = 0
            
            cur += num
            if cur > res:
                res = cur
        
        return res

if __name__ == '__main__':
    s = Solution()

    assert s.maxSubArray([-2,1,-3,4,-1,2,1,-5,4]) == 6
    assert s.maxSubArray([1]) == 1
    assert s.maxSubArray([5,4,-1,7,8]) == 23