from re import L
from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        left = 0
        right = len(arr)-k

        while left < right:
            mid = (left+right) // 2
            if x - arr[mid] > arr[mid+k] - x:
                left = mid + 1
            else:
                right = mid
        

        return arr[left:left+k]

if __name__ == '__main__':
    s = Solution()

    assert s.findClosestElements([1,2,3,4,5], 4, 3) == [1,2,3,4]
    assert s.findClosestElements([1,2,3,4,5], 4, -1) == [1,2,3,4]