from typing import List

class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        head = 0
        tail = len(nums)-1

        while head < tail:
            if nums[head]%2 == 1 and nums[tail]%2 == 0:
                nums[head], nums[tail] = nums[tail], nums[head]
                head += 1
                tail -= 1
            elif nums[head]%2 == 0:
                head += 1
            elif nums[tail]%2 == 1:
                tail -= 1
        
        return nums

if __name__ == '__main__':
    s = Solution()

    assert s.sortArrayByParity([3,1,2,4]) == [4,2,1,3]
    assert s.sortArrayByParity([0]) == [0]