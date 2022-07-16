from typing import List

class Solution:
    # Time Cimplexity: O(n)
    # Space Cimplexity: O(1)
    def runningSum(self, nums: List[int]) -> List[int]:
        for i in range (1, len(nums)):
            nums[i] += nums[i-1]
        
        return nums

if __name__ == '__main__':
    s = Solution()

    assert s.runningSum([1,2,3,4]) == [1,3,6,10]
    assert s.runningSum([1,1,1,1,1]) == [1,2,3,4,5]
    assert s.runningSum([3,1,2,10,1]) == [3,4,6,16,17]