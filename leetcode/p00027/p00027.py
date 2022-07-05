from typing import List

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        ptr = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[ptr] = nums[i]
                ptr += 1

        return ptr

if __name__ == '__main__':
    s = Solution()

    assert s.removeElement([3,2,2,3], 3) == 2
    assert s.removeElement([0,1,2,2,3,0,4,2], 2) == 5