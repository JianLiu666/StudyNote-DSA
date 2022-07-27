from collections import deque
from typing import List

class Solution:
    # Time Complexity: O(m+n)
    # Space Complexity: O(m)
    def numIslands_bfs(self, grid: List[List[str]]) -> int:
        # traverse whole grid
        count = 0
        for row in range(len(grid)):
            for col in range(len(grid[0])):
                if grid[row][col] == "1":
                    # traveres whole island
                    self.bfs(grid, row, col)
                    count += 1

        return count

    # Time Complexity: O(m)
    # Space Complexity: O(m)
    def bfs(self, grid: List[List[str]], row: int, col: int):
        gridHeight = len(grid)-1
        gridWidth = len(grid[0])-1

        q = deque([(row, col)])
        while q:
            row, col = q.popleft()
            if 0 <= row <= gridHeight and 0 <= col <= gridWidth and grid[row][col] == "1":
                grid[row][col] = "2"
                q.extend([
                    (row-1, col),
                    (row+1, col),
                    (row, col-1),
                    (row, col+1)])
    
    # Time Complexity: O(m+n)
    # Space Complexity: O(m)
    def numIslands_dfs(self, grid: List[List[str]]) -> int:
        count = 0

        for row in range(len(grid)):
            for col in range(len(grid[0])):
                if grid[row][col] == "1":
                    self.dfs(grid, row, col)
                    count += 1
        
        return count

    def dfs(self, grid: List[List[str]], row: int, col: int):
        grid[row][col] = "2"

        if row-1 >= 0 and grid[row-1][col] == "1":
            self.dfs(grid, row-1, col)
        
        if col+1 <= len(grid[0])-1 and grid[row][col+1] == "1":
            self.dfs(grid, row, col+1)
        
        if row+1 <= len(grid)-1 and grid[row+1][col] == "1":
            self.dfs(grid, row+1, col)

        if col-1 >= 0 and grid[row][col-1] == "1":
            self.dfs(grid, row, col-1)

if __name__ == '__main__':
    s = Solution()

    assert s.numIslands_bfs([
        ["1","1","1","1","0"],
        ["1","1","0","1","0"],
        ["1","1","0","0","0"],
        ["0","0","0","0","0"]]) == 1

    assert s.numIslands_bfs([
        ["1","1","0","0","0"],
        ["1","1","0","0","0"],
        ["0","0","1","0","0"],
        ["0","0","0","1","1"]]) == 3

    assert s.numIslands_dfs([
        ["1","1","1","1","0"],
        ["1","1","0","1","0"],
        ["1","1","0","0","0"],
        ["0","0","0","0","0"]]) == 1

    assert s.numIslands_dfs([
        ["1","1","0","0","0"],
        ["1","1","0","0","0"],
        ["0","0","1","0","0"],
        ["0","0","0","1","1"]]) == 3
