from typing import List
from collections import defaultdict

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(n^2) 
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        lookup = defaultdict(int)
        for n1 in nums1: 
            for n2 in nums2:
                lookup[n1+n2] += 1
        
        count = 0
        for n3 in nums3:
            for n4 in nums4:
                count += lookup[-(n3+n4)]
        
        return count

if __name__ == '__main__':
    s = Solution()

    assert s.fourSumCount([1,2],[-2,-1],[-1,2],[0,2]) == 2
    assert s.fourSumCount([-1,-1],[-1,1],[-1,1],[1,-1]) == 6
    