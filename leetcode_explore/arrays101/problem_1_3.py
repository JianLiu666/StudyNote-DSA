'''
Squares of a Sorted Array
    Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
    Input: nums = [-4,-1,0,3,10]
    Output: [0,1,9,16,100]
    Explanation:
        After squaring, the array becomes [16,1,0,9,100].
        After sorting, it becomes [0,1,9,16,100].

Example 2:
    Input: nums = [-7,-3,2,3,11]
    Output: [4,9,9,49,121]

Constraints:
    1 <= nums.length <= 10^4
    -10^4 <= nums[i] <= 10^4
    nums is sorted in non-decreasing order.
'''

'''
解題方向:
    - 題目已經明確表示為 non-descreasing order, 只需考慮如何解決負數問題
    - Two-pointers 解法: head and tail pointer
        - head pointer 從頭開始檢查
        - tail pointer 從尾巴開始檢查
    - 持續比較頭尾數字平方後的大小, 將大的那邊塞進新的 list (從尾巴開始往前塞), 然後往對方前進
    - 直到 head pointer 與 tail pointer 相遇 -> O(n)
'''

from typing import List

# My Solution
# Time Complexity: O(n)
class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        head, tail = 0, len(nums)-1
        offset = len(nums)-1

        result = [None] * len(nums)
        while head != tail:
            if abs(nums[head]) > abs(nums[tail]):
                result[offset] = nums[head] ** 2
                head += 1
            else:
                result[offset] = nums[tail] ** 2
                tail -= 1
            offset -= 1

        result[offset] = nums[head] ** 2

        return result

if __name__ == '__main__':
    s = Solution()

    assert s.sortedSquares([-4,-1,0,3,10]) == [0,1,9,16,100]
    assert s.sortedSquares([-7,-3,2,3,11]) == [4,9,9,49,121]