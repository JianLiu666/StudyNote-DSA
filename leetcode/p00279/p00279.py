from collections import deque
from math import isqrt

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def numSquares_bfs(self, n: int) -> int:
        if n <= 3:
            return n
        
        q = deque([(n, 1)])
        step = 0

        while q:
            size = len(q)
            step += 1
            for _ in range(size):
                remainder, minNum = q.popleft()

                num = isqrt(remainder)
                square = num * num
                while num >= minNum:
                    if remainder == square:
                        return step
                    
                    q.append((remainder-square, num))
                    num -= 1
                    square = num * num
        
        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.numSquares_bfs(12) == 3
    assert s.numSquares_bfs(13) == 2
