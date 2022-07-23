from re import L
from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [[0]*9 for _ in range(9)]
        cols = [[0]*9 for _ in range(9)]
        recs = [[0]*9 for _ in range(9)]

        for row, columns in enumerate(board):
            for col, ch in enumerate(columns):
                if ch[0] == '.':
                    continue

                num = ord(ch[0])-48-1
                rows[row][num] += 1
                if rows[row][num] > 1:
                    return False
                
                cols[col][num] += 1
                if cols[col][num] > 1:
                    return False
                
                rec = (row//3)*3 + (col//3)
                recs[rec][num] += 1
                if recs[rec][num] > 1:
                    return False
            
        return True

if __name__ == '__main__':
    s = Solution()

    assert s.isValidSudoku([
        ["5","3",".",".","7",".",".",".","."],
        ["6",".",".","1","9","5",".",".","."],
        [".","9","8",".",".",".",".","6","."],
        ["8",".",".",".","6",".",".",".","3"],
        ["4",".",".","8",".","3",".",".","1"],
        ["7",".",".",".","2",".",".",".","6"],
        [".","6",".",".",".",".","2","8","."],
        [".",".",".","4","1","9",".",".","5"],
        [".",".",".",".","8",".",".","7","9"]]
    ) == True

    assert s.isValidSudoku([
        ["8","3",".",".","7",".",".",".","."],
        ["6",".",".","1","9","5",".",".","."],
        [".","9","8",".",".",".",".","6","."],
        ["8",".",".",".","6",".",".",".","3"],
        ["4",".",".","8",".","3",".",".","1"],
        ["7",".",".",".","2",".",".",".","6"],
        [".","6",".",".",".",".","2","8","."],
        [".",".",".","4","1","9",".",".","5"],
        [".",".",".",".","8",".",".","7","9"]]
    ) == False