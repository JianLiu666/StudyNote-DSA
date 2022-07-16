from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def dominantIndex(self, nums: List[int]) -> int:
        t1 = nums[0]
        t2 = nums[1]
        idx = 0
        if t2 > t1:
            t1, t2 = t2, t1
            idx = 1

        for i in range(2, len(nums)):
            if nums[i] > t1:
                t1, t2 = nums[i], t1
                idx = i
            elif nums[i] > t2:
                t2 = nums[i]
        
        if t1 >= t2 * 2:
            return idx
        
        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.dominantIndex([3,6,1,0]) == 1
    assert s.dominantIndex([1,2,3,4]) == -1