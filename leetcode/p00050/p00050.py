class Solution:
    def myPow(self, x: float, n: int) -> float:
        if x == 0:
            return 0
        if n == 0:
            return 1
        
        neg = False
        if n < 0:
            n = -n
            neg = True
        
        ans = 1
        while n != 0:
            if n%2 == 1:
                ans *= x
            x *= x
            n >>= 1
        
        if neg:
            ans = 1/ans
        return round(ans*100000)/100000

if __name__ == '__main__':
    s = Solution()

    assert s.myPow(2.00000, 4) == 16.00000
    assert s.myPow(2.00000, 10) == 1024.00000
    assert s.myPow(2.10000, 3) == 9.26100