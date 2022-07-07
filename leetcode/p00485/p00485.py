from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        maximum = 0
        counter = 0

        for num in nums:
            if num == 1:
                counter += 1
                if counter > maximum:
                    maximum = counter
            else:
                counter = 0
        
        return maximum

if __name__ == '__main__':
    s = Solution()

    assert s.findMaxConsecutiveOnes([1,1,0,1,1,1]) == 3
    assert s.findMaxConsecutiveOnes([1,0,1,1,0,1]) == 2