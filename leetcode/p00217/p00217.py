from typing import List

class Solution:
    # Time: O(nlogn)
    # Space: O(n)
    def containsDuplicate_sort(self, nums: List[int]) -> bool:
        merge_sort(nums)
        for i in range(1, len(nums), 1):
            if nums[i-1] == nums[i]:
                return True
        
        return False

    # Time: O(n)
    # Space: O(n)
    def containsDuplicate_dict(self, nums: List[int]) -> bool:
        seen = {}
        for num in nums:
            if num in seen:
                return True
            else:
                seen[num] = 1
        
        return False

def merge_sort(nums: List[int]):
    if len(nums) == 1:
        return
    
    mid = len(nums) // 2

    left_arr = nums[:mid]
    merge_sort(left_arr)
    
    right_arr = nums[mid:]
    merge_sort(right_arr)

    left_idx = 0
    right_idx = 0
    nums_idx = 0
    
    while left_idx < len(left_arr) and right_idx < len(right_arr):
        if left_arr[left_idx] <= right_arr[right_idx]:
            nums[nums_idx] = left_arr[left_idx]
            left_idx += 1
        else:
            nums[nums_idx] = right_arr[right_idx]
            right_idx += 1
        nums_idx += 1
    
    while left_idx < len(left_arr):
        nums[nums_idx] = left_arr[left_idx]
        left_idx += 1
        nums_idx += 1
    
    while right_idx < len(right_arr):
        nums[nums_idx] = right_arr[right_idx]
        right_idx += 1
        nums_idx += 1

if __name__ == '__main__':
    s = Solution()

    assert s.containsDuplicate_sort([1,2,3,1]) == True
    assert s.containsDuplicate_sort([1,2,3,4]) == False
    assert s.containsDuplicate_sort([1,1,1,3,3,4,3,2,4,2]) == True

    assert s.containsDuplicate_dict([1,2,3,1]) == True
    assert s.containsDuplicate_dict([1,2,3,4]) == False
    assert s.containsDuplicate_dict([1,1,1,3,3,4,3,2,4,2]) == True