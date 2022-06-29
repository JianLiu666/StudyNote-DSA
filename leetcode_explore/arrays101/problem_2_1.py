'''
Duplicate Zeros
    Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.
    Note that elements beyond the length of the original array are not written.
    Do the above modifications to the input array in place and do not return anything.

Example 1:
    Input: arr = [1,0,2,3,0,4,5,0]
    Output: [1,0,0,2,3,0,0,4]
    Explanation: 
        After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:
    Input: arr = [1,2,3]
    Output: [1,2,3]
    Explanation:
        After calling your function, the input array is modified to: [1,2,3]
 
Constraints:
    1 <= arr.length <= 10^4
    0 <= arr[i] <= 9
'''

'''
解題方向:
    - Brute-force 就是遇到零的時候按照規則往後位移 -> O(n^2)
    - Two-pointers: fixed width -> left pointer and right pointer
        - 第一次先掃過一遍確定有多少個0, 定位出新陣列的終點 (決定 left pointer and right pointer 的位置)
        - 逐一往前複製, 遇到0時要特別處理, right pointer 要多補一個0
'''

from typing import List

class Solution:
        
    # My Solution
    # Time Complexity: O(n^2)
    def duplicateZeros_sol1(self, arr: List[int]) -> None:
        offset = 0
        while offset < len(arr)-1:
            if arr[offset] != 0:
                offset += 1
                continue
            
            # 從尾巴開始將資料依序往後位移
            for idx in range(len(arr)-2, offset, -1):
                arr[idx+1] = arr[idx]
            
            # 在指定位置補0
            arr[offset+1] = 0
            offset += 2

    # LeetCode Solution
    # Time Complexity: O(n^2)
    #  - built-in pop operation at the last element -> O(1)
    #  - built-in insert operation -> O(n)
    def duplicateZeros_sol2(self, arr: List[int]) -> None:
        offset = 0
        while offset < len(arr):
            if arr[offset] == 0:
                arr.pop()
                arr.insert(offset, 0)
                offset += 1
            offset += 1

    # LeetCode Solution
    # Time Complexity: O(2n)
    def duplicateZeros_sol3(self, arr: List[int]) -> None:
        left = -1
        capacity = len(arr)
        right = capacity -1

        # 定位出要複製的最末端
        while capacity > 0:
            left += 1
            capacity -= 1

            if arr[left] == 0:
                capacity -= 1

        # 最後一個位置是 0 的時候
        if arr[left] == 0 and capacity < 0:
            arr[right] = arr[left]
            left -= 1
            right -= 1

        # 由後往前逐一複製
        while left > -1:
            arr[right] = arr[left]
            right -= 1

            if arr[left] == 0 and left <= right:
                arr[right] = arr[left]
                right -= 1

            left -= 1

if __name__ == '__main__':
    s = Solution()

    list1 = [1,0,2,3,0,4,5,0]
    list2 = [1,2,3]

    s.duplicateZeros_sol1(list1)
    assert list1 == [1,0,0,2,3,0,0,4]

    s.duplicateZeros_sol1(list2)
    assert list2 == [1,2,3]

    list1 = [1,0,2,3,0,4,5,0]
    list2 = [1,2,3]

    s.duplicateZeros_sol2(list1)
    assert list1 == [1,0,0,2,3,0,0,4]

    s.duplicateZeros_sol2(list2)
    assert list2 == [1,2,3]

    list1 = [1,0,2,3,0,4,5,0]
    list2 = [1,2,3]

    s.duplicateZeros_sol3(list1)
    assert list1 == [1,0,0,2,3,0,0,4]

    s.duplicateZeros_sol3(list2)
    assert list2 == [1,2,3]