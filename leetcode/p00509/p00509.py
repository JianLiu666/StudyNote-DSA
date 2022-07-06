from math import sqrt

class Solution:
    # Time Complexity: O(2^n)
    # Space Complexity: O(n)
    def fib_recursion(self, n: int) -> int:
        if n < 2:
            return n
        
        return self.fib_recursion(n-1) + self.fib_recursion(n-2)

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def fib_dp(self, n: int) -> int:
        if n < 2:
            return n
        
        fn, fn_1, fn_2 = 0, 1, 0
        for i in range(1, n):
            fn = fn_1 + fn_2
            fn_2 = fn_1
            fn_1 = fn
        
        return fn

    # Time Complexity: O(1)
    # Space Complexity: O(1)
    def fib_math(self, n: int) -> int:
        phi = (sqrt(5)+1) / 2

        return round(pow(phi, n) / sqrt(5))


if __name__ == '__main__':
    s = Solution()

    assert s.fib_recursion(2) == 1
    assert s.fib_recursion(3) == 2
    assert s.fib_recursion(4) == 3

    assert s.fib_dp(2) == 1
    assert s.fib_dp(3) == 2
    assert s.fib_dp(4) == 3

    assert s.fib_math(2) == 1
    assert s.fib_math(3) == 2
    assert s.fib_math(4) == 3