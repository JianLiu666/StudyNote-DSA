from typing import List

class Solution:
    
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def heightChecker(self, heights: List[int]) -> int:
        counter = [0] * 100

        for n in heights:
            counter[n-1] += 1

        mismatch = 0
        current = 0
        for i, v in enumerate(counter):
            for j in range(v):
                if heights[current] != i+1:
                    mismatch += 1
                current += 1
        
        return mismatch

if __name__ == '__main__':
    s = Solution()

    assert s.heightChecker([1,1,4,2,1,3]) == 3
    assert s.heightChecker([5,1,2,3,4]) == 5