from typing import List

class Solution:
    # Time Complexity: O(nlogn)
    # Space Complexity: O(1)
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        maxRow, maxCol = len(matrix)-1, len(matrix[0])-1
        low, high = matrix[0][0], matrix[maxRow][maxCol]

        while low < high:
            mid = low + (high-low)//2
            count = 0
            col = maxCol
            for row in range(maxRow+1):
                while col >= 0 and matrix[row][col] > mid:
                    col -= 1
                count += col+1
            
            if count < k:
                low = mid + 1
            else:
                high = mid
        
        return low

if __name__ == '__main__':
    s = Solution()

    assert s.kthSmallest([[1,5,9],[10,11,13],[12,13,15]], 8) == 13
    assert s.kthSmallest([[-5]], 1) == -5