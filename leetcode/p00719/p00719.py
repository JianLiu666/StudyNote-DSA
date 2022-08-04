from typing import List

class Solution:
    # Time Complexity: O(nlogn)
    # Space Complexity: O(n)
    def smallestDistancePair(self, nums: List[int], k: int) -> int:
        
        def count(distance: int) -> bool:
            acc, cur1, cur2, size = 0, 0, 0, len(nums)
            while cur1 < size:
                while cur2 < size and nums[cur2]-nums[cur1] <= distance:
                    cur2 += 1
                acc += cur2 - cur1 - 1
                cur1 += 1
                if acc >= k:
                    return True
            
            return False
        
        nums.sort()
        low, high = 0, nums[-1] - nums[0]
        while low < high:
            mid = low + (high-low) // 2
            if count(mid):
                high = mid
            else:
                low = mid + 1
        
        return low

if __name__ == '__main__':
    s = Solution()

    assert s.smallestDistancePair([1,3,1], 1) == 0
    assert s.smallestDistancePair([1,1,1], 2) == 0
    assert s.smallestDistancePair([1,6,1], 3) == 5