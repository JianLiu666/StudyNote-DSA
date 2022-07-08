from typing import List
import math

class Solution:
    def findNumbers(self, nums: List[int]) -> int:
        count = 0
        for n in nums:
            count += int(math.log10(n)) & 1
                
        return count

if __name__ == '__main__':
    s = Solution()

    assert s.findNumbers([12,345,2,6,7896]) == 2
    assert s.findNumbers([555,901,482,1771]) == 1