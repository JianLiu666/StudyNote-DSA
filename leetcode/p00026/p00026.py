from typing import List

class Solution:

    # Time Complexity: O(n)
    # Space Cpmplexity: O(1)
    def removeDuplicates(self, nums: List[int]) -> int:
        count = 0
        for i in range(1, len(nums)):
            if nums[count] != nums[i]:
                count += 1
                nums[count] = nums[i]
        
        return count+1

if __name__ == '__main__':
    s = Solution()

    assert s.removeDuplicates([1,1,2]) == 2
    assert s.removeDuplicates([0,0,1,1,1,2,2,3,3,4]) == 5
