from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        head = 0
        summation = 0
        length = len(nums)

        for current in range(len(nums)):
            summation += nums[current]

            while summation >= target:
                if length > current-head+1:
                    length = current-head+1
                
                summation -= nums[head]
                head += 1

        if head == 0 and summation < target:
            return 0
        
        return length

if __name__ == '__main__':
    s = Solution()

    assert s.minSubArrayLen(7, [2,3,1,2,4,3]) == 2
    assert s.minSubArrayLen(4, [1,4,4]) == 1
    assert s.minSubArrayLen(11, [1,1,1,1,1,1,1,1]) == 0
    assert s.minSubArrayLen(15, [1,2,3,4,5]) == 5