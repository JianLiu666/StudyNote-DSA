from typing import List

class Solution:

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def validMountainArray(self, arr: List[int]) -> bool:
        head = 0
        tail = len(arr)-1
        while head != tail:
            if arr[head] >= arr[head+1] and arr[tail-1] <= arr[tail]:
                return False

            if arr[head] < arr[head+1]:
                head += 1
            
            if arr[tail] < arr[tail-1]:
                tail -= 1

        if head == 0 or tail == len(arr)-1:
            return False
        
        return True

if __name__ == '__main__':
    s = Solution()

    assert s.validMountainArray([2,1]) == False
    assert s.validMountainArray([3,5,5]) == False
    assert s.validMountainArray([0,3,2,1]) == True
