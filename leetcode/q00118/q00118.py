import time
from typing import List

class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        # 將 Pascal's Triangle 的值初始化為 1
        levels = [[1] * level for level in range(1, numRows+1, 1)]

        for row in range(2, numRows, 1):
            for col in range(1, row, 1):
                levels[row][col] = levels[row-1][col-1] + levels[row-1][col]

        print(levels)
        return levels

if __name__ == '__main__':
    s = Solution()
    assert s.generate(5) == [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
    assert s.generate(1) == [[1]]
