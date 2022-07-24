from typing import List

class Solution:
    # Time Complexity: O(m+n)
    # Space Complexity: O(1)
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        col = len(matrix)-1
        row = 0

        while col >= 0 and col <= len(matrix)-1 and row >= 0 and row <= len(matrix[0])-1:
            if matrix[col][row] == target:
                return True
            elif matrix[col][row] > target:
                col -= 1
            else:
                row += 1

        return False

if __name__ == '__main__':
    s = Solution()

    assert s.searchMatrix([
        [1,4,7,11,15],
        [2,5,8,12,19],
        [3,6,9,16,22],
        [10,13,14,17,24],
        [18,21,23,26,30]], 5) == True

    assert s.searchMatrix([
        [1,4,7,11,15],
        [2,5,8,12,19],
        [3,6,9,16,22],
        [10,13,14,17,24],
        [18,21,23,26,30]], 20) == False