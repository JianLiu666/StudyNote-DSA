'''
Valid Mountain Array
    Given an array of integers arr, return true if and only if it is a valid mountain array.
    Recall that arr is a mountain array if and only if:
        arr.length >= 3
        There exists some i with 0 < i < arr.length - 1 such that:
            arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
            arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

Example 1:
    Input: arr = [2,1]
    Output: false

Example 2:
    Input: arr = [3,5,5]
    Output: false

Example 3:
    Input: arr = [0,3,2,1]
    Output: true
 
Constraints:
    1 <= arr.length <= 10^4
    0 <= arr[i] <= 10^4
'''

from typing import List

# My Solution
# Time Complexity: O(n)
class Solution:
    def validMountainArray(self, arr: List[int]) -> bool:
        head, tail = 0, len(arr)-1
        
        while head < tail and arr[head] < arr[head+1]:
            head += 1
        if head == 0:
            return False
        
        while tail > head and arr[tail-1] > arr[tail]:
            tail -= 1
        if tail == len(arr)-1:
            return False
        
        return head == tail

if __name__ == '__main__':
    s = Solution()

    assert s.validMountainArray([2,1]) == False
    assert s.validMountainArray([3,5,5]) == False
    assert s.validMountainArray([0,3,2,1]) == True