from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        if len(nums) <= 1:
            return False
        
        seen = {}

        for idx, num in enumerate(nums):
            if num in seen and idx - seen[num] <= k:
                return True
            seen[num] = idx
        
        return False

if __name__ == '__main__':
    s = Solution()

    assert s.containsNearbyDuplicate([1,2,3,1], 3) == True
    assert s.containsNearbyDuplicate([1,0,1,1], 1) == True
    assert s.containsNearbyDuplicate([1,2,3,1,2,3], 2) == False