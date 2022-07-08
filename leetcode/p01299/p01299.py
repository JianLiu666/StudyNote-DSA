from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def replaceElements(self, arr: List[int]) -> List[int]:
        maximum = arr[-1]
        arr[-1] = -1
        
        for n in range (len(arr)-2, -1, -1):
            if arr[n] > maximum:
                arr[n], maximum = maximum, arr[n]
            else:
                arr[n] = maximum
        
        return arr

if __name__ == '__main__':
    s = Solution()

    assert s.replaceElements([17,18,5,4,6,1]) == [18,6,6,6,1,-1]
    assert s.replaceElements([400]) == [-1]