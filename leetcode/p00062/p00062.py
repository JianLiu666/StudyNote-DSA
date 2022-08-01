class Solution:
    # Time Complexity: O(mn)
    # Space Complexity: O(mn)
    def uniquePaths(self, m: int, n: int) -> int:
        matrix = [[1 for _ in range(n)] for _ in range(m)]

        for row in range(m):
            for col in range(n):
                if row > 0 and col > 0:
                    matrix[row][col] = matrix[row-1][col] + matrix[row][col-1]
        
        return matrix[m-1][n-1]

if __name__ == '__main__':
    s = Solution()

    assert s.uniquePaths(3, 7) == 28
    assert s.uniquePaths(3, 2) == 3