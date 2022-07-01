from typing import Dict, List

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(1) -> Assuming output list is not considered
    def generate(self, numRows: int) -> List[List[int]]:
        result = []
        
        result.append([1])
        if numRows == 1:
            return result
        
        result.append([1,1])
        if numRows == 2:
            return result
        
        currentLayer = [1,1]
        for i in range(2, numRows, 1):
            currentLayer = self.generateNextLayer(currentLayer)
            result.append(currentLayer)
            
        return result
    
    def generateNextLayer(self, arr: List[int]) -> List[int]:
        result = [1]
        
        right = 1
        while right < len(arr):
            result += [arr[right-1] + arr[right]]
            right += 1
        
        return result + [1]

class Solution2:
    # Time Complexity: O(n)
    # Space Complexity: O(n^2)
    def generate(self, numRows: int) -> List[List[int]]:
        # 將 Pascal's Triangle 的值初始化為 1
        levels = [[1] * level for level in range(1, numRows+1)]
        memo = {}
        for col in range(1, numRows-1):
            self.recurse(numRows-1, col, levels)

        return levels

    def recurse(self, row: int, col: int, levels: List[List[int]], memo: Dict[List[int], int]) -> int:
        if col == 0 or row <= 1:
            return 1
        if (row, col) in memo:
            return memo[(row, col)]

        value = self.recurse(row-1, col-1) + self.recurse(row-1, col)
        levels[row][col] = value
        memo[(row, col)] = value


if __name__ == '__main__':
    s = Solution()
    assert s.generate(5) == [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
    assert s.generate(1) == [[1]]

    s2 = Solution()
    assert s2.generate(5) == [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
    assert s2.generate(1) == [[1]]