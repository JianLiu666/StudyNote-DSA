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

from typing import List

# My Solution
# Time Complexity: O(n^2)
def duplicateZeros_sol1(arr: List[int]) -> None:
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

list1 = [1,0,2,3,0,4,5,0]
duplicateZeros_sol1(list1)
print(list1)

# LeetCode Solution
# Time Complexity: O(n^2)
#  - built-in pop operation at the last element -> O(1)
#  - built-in insert operation -> O(n)
def duplicateZeros_sol2(arr: List[int]) -> None:
    offset = 0
    while offset < len(arr):
        if arr[offset] == 0:
            arr.pop()
            arr.insert(offset, 0)
            offset += 1
        offset += 1

list2 = [1,0,2,3,0,4,5,0]
duplicateZeros_sol2(list2)
print(list2)