class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        if n <= 0:
            return False
        
        while n % 3 == 0:
            n /= 3
        
        return n == 1

if __name__ == '__main__':
    s = Solution()

    assert s.isPowerOfThree(27) == True
    assert s.isPowerOfThree(0) == False
    assert s.isPowerOfThree(9) == True