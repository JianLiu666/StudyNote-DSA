class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def climbStairs(self, n: int) -> int:
        if n == 1:
            return 1
        
        nxt = 2
        pre = 1

        for i in range(3, n+1, 1):
            nxt, pre = nxt+pre, nxt
        
        return nxt

if __name__ == '__main__':
    s = Solution()

    s.climbStairs(2) == 2
    s.climbStairs(3) == 3