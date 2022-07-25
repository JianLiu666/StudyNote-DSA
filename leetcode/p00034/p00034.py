from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        if len(nums) == 0:
            return [-1, -1]
        
        anchor = binary_search(nums, 0, len(nums)-1, target)
        if anchor == -1:
            return [-1, -1]
        
        res = []

        edge = anchor
        while edge != 0 and nums[edge-1] == target:
            edge = binary_search(nums, 0, edge-1, target)
        res.append(edge)

        edge = anchor
        while edge != len(nums)-1 and nums[edge+1] == target:
            edge = binary_search(nums, edge+1, len(nums)-1, target)
        res.append(edge)

        return res

def binary_search(nums: List[int], head: int, tail: int, target: int) -> int:
    while head <= tail:
        mid = head + (tail-head)//2
        if nums[mid] == target:
            return mid
        elif nums[mid] > target:
            tail = mid - 1
        else:
            head = mid + 1
    
    return -1

if __name__ == '__main__':
    s = Solution()

    assert s.searchRange([5,7,7,8,8,8,8,8,8,9,9,9], 8) == [3,8]
    assert s.searchRange([5,7,7,8,8,10], 8) == [3,4]
    assert s.searchRange([5,7,7,8,8,10], 6) == [-1,-1]
    assert s.searchRange([], 0) == [-1,-1]
    assert s.searchRange([2,2], 2) == [0,1]
