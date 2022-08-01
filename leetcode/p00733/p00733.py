from typing import List
from collections import deque

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        if image[sr][sc] == color:
            return image
        
        target = image[sr][sc]

        q = deque([(sr, sc)])
        image[sr][sc] = color

        while q:
            size = len(q)
            for _ in range(size):
                row, col = q.popleft()

                if row-1 >= 0 and image[row-1][col] == target:
                    q.append((row-1, col))
                    image[row-1][col] = color
                if row+1 < len(image) and image[row+1][col] == target:
                    q.append((row+1, col))
                    image[row+1][col] = color
                if col-1 >= 0 and image[row][col-1] == target:
                    q.append((row, col-1))
                    image[row][col-1] = color
                if col+1 < len(image[0]) and image[row][col+1] == target:
                    q.append((row, col+1))
                    image[row][col+1] = color
        
        return image

if __name__ == '__main__':
    s = Solution()

    assert s.floodFill([[1,1,1],[1,1,0],[1,0,1]], 1, 1, 2) == [[2,2,2],[2,2,0],[2,0,1]]
    assert s.floodFill([[0,0,0],[0,0,0]], 0, 0, 0) == [[0,0,0],[0,0,0]]
