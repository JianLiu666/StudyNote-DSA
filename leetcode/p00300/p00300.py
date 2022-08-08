from bisect import bisect_left
from typing import List

class Solution:
    # Time Complexity: O(nlogn)
    # Space Complexity: O(n)
    def lengthOfLIS(self, nums: List[int]) -> int:
        seq = [nums[0]]
        for i in range(1, len(nums)):
            if nums[i] > seq[-1]:
                seq.append(nums[i])
            elif nums[i] < seq[-1]:
                idx = bisect_left(seq, nums[i])
                seq[idx] = nums[i]
        
        return len(seq)

if __name__ == '__main__':
    s = Solution()

    assert s.lengthOfLIS([10,9,2,5,3,7,101,18]) == 4
    assert s.lengthOfLIS([0,1,0,3,2,3]) == 4
    assert s.lengthOfLIS([7,7,7,7,7,7,7,7]) == 1
