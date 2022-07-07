from typing import List

# Time Complexity: O(n)
# Space Complexity: O(1)
class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        result = []

        size = len(nums)+1
        for i in range(len(nums)):
            nums[(nums[i]%size)-1] += size

        print(nums)
        for i in range(len(nums)):
            if nums[i] // size == 0:
                result.append(i+1)
        
        return result

if __name__ == '__main__':
    s = Solution()

    assert s.findDisappearedNumbers([4,3,2,7,8,2,3,1]) == [5,6]
    assert s.findDisappearedNumbers([1,1]) == [2]