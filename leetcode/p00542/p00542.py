from typing import List
from collections import deque

class Solution:
    # Time Complexity: O(mn), where m is the length of rows, and n is the length of columns
    # Space Complexity: O(mn), where m is the length of rows, and n is the length of columns
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        res = [[float('inf') for _ in range(len(mat[0]))] for _ in range(len(mat))]

        q = deque()
        for row in range(len(mat)):
            for col in range(len(mat[0])):
                if mat[row][col] == 0:
                    res[row][col] = 0
                    q.append((row, col))

        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        while q:
            row, col = q.popleft()

            for i in range(4):
                newRow = row + directions[i][0]
                newCol = col + directions[i][1]
                if 0 <= newRow < len(res) and 0 <= newCol < len(res[0]) and res[newRow][newCol] > res[row][col] + 1:
                    res[newRow][newCol] = res[row][col] + 1
                    q.append((newRow, newCol))
        
        return res

