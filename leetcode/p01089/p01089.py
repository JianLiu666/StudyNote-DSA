from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def duplicateZeros(self, arr: List[int]) -> None:
        head = len(arr)-1
        tail = len(arr)-1
        current = 0
        
        while current < head:
            if arr[current] == 0:
                head -= 1
            current += 1

            if arr[current] == 0 and current == head:
                arr[tail] = 0
                head -= 1
                tail -= 1
        
        while head > -1 and head != tail:
            arr[tail] = arr[head]
            if arr[head] == 0:
                tail -= 1
                arr[tail] = arr[head]
            head -= 1
            tail -= 1
 
if __name__ == '__main__':
    s = Solution()

    list = [1, 0, 2, 3, 0, 4, 5, 0]
    s.duplicateZeros(list) == [1, 0, 0, 2, 3, 0, 0, 4]

    list = [1, 2, 3]
    s.duplicateZeros(list) == [1, 2, 3]

    list = [1, 5, 2, 0, 6, 8, 0, 6, 0]
    s.duplicateZeros(list) == [1, 5, 2, 0, 0, 6, 8, 0, 0]

    list = [1, 5, 2, 0, 6, 8, 1, 0, 0]
    s.duplicateZeros(list) == [1, 5, 2, 0, 0, 6, 8, 1, 0]