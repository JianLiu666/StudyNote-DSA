from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def firstMissingPositive_swap(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            while nums[i] > 0 and nums[i] <= len(nums) and nums[i] != nums[nums[i]-1]:
                tmp = nums[nums[i]-1]
                nums[nums[i]-1] = nums[i]
                nums[i] = tmp

        for i in range(len(nums)):
            if nums[i] != i+1:
                return i+1

        return len(nums)+1

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def firstMissingPositive_hash(self, nums: List[int]) -> int:
        nums.append(0)
        size = len(nums)
        
        for i in range(len(nums)):
            if nums[i] < 0 or nums[i] >= size:
                nums[i] = 0

        for i in range(len(nums)):
            nums[nums[i]%size] += size
        
        for i in range(1, len(nums)):
            if nums[i] // size == 0:
                return i
        
        return size


if __name__ == '__main__':
    s = Solution()

    assert s.firstMissingPositive_swap([1, 2, 0]) == 3
    assert s.firstMissingPositive_swap([1, 2, 3]) == 4
    assert s.firstMissingPositive_swap([3, 4, -1, 1]) == 2

    assert s.firstMissingPositive_swap([1, 2, 0]) == 3
    assert s.firstMissingPositive_swap([1, 2, 3]) == 4
    assert s.firstMissingPositive_swap([3, 4, -1, 1]) == 2
