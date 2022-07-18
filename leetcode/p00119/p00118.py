from typing import List

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(1)
    def getRow(self, rowIndex: int) -> List[int]:
        row = [0] * (rowIndex+1)
        row[0] = 1

        for i in range(1, rowIndex+1):
            for j in range(i, 0, -1):
                row[j] += row[j-1]

        return row

if __name__ == '__main__':
    s = Solution()
    assert s.getRow(3) == [1,3,3,1]
    assert s.getRow(0) == [1]
    assert s.getRow(1) == [1,1]
