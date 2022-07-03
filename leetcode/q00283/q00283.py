from typing import List

class Solution:

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def moveZeroes(self, nums: List[int]) -> None:
        nonzero_ptr = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[i], nums[nonzero_ptr] = nums[nonzero_ptr], nums[i]
                nonzero_ptr += 1

if __name__ == '__main__':
    s = Solution()

    list1 = [0,1,0,3,12]
    s.moveZeroes(list1)
    assert list1 == [1,3,12,0,0]

    list2 = [0]
    s.moveZeroes(list2)
    assert list2 == [0]
