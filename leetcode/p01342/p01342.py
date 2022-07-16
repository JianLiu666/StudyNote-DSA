class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def numberOfSteps(self, num: int) -> int:
        if num == 0:
            return num
        
        count = -1
        while num != 0:
            count += 1 if num & 1 == 0 else 2
            num >>= 1
        
        return count

if __name__ == '__main__':
    s = Solution()

    assert s.numberOfSteps(14) == 6
    assert s.numberOfSteps(8) == 4
    assert s.numberOfSteps(123) == 12