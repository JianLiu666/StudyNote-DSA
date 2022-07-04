from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def missingNumber_xor(self, nums: List[int]) -> int:
        full_xor = 0
        nums_xor = 0

        for i in range(len(nums)):
            full_xor ^= i
            nums_xor ^= nums[i]
        full_xor ^= len(nums)

        return full_xor ^ nums_xor

    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def missingNumber_math(self, nums: List[int]) -> int:
        l = len(nums)

        actual_sum = (l*(l+1)) // 2
        real_sum = sum(nums)

        return actual_sum - real_sum

    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def missingNumber_dict(self, nums: List[int]) -> int:
        seen = set(nums)

        for i in range(len(nums)):
            if i not in seen:
                return i

        return len(nums)

    # Time Complexity: O(nlogn)
    # Space Complexity: O(n)
    def missingNumber_sort(self, nums: List[int]) -> int:
        sorted_arr = merge_sort(nums)

        for i in range(len(sorted_arr)):
            if i != sorted_arr[i]:
                return i
        
        return len(sorted_arr)

def merge_sort(nums: List[int]) -> List[int]:
    if len(nums) == 1:
        return nums

    mid = len(nums) // 2
    l_arr = merge_sort(nums[:mid])
    r_arr = merge_sort(nums[mid:])

    l_idx = 0
    r_idx = 0
    current = 0

    while l_idx < len(l_arr) and r_idx < len(r_arr):
        if l_arr[l_idx] <= r_arr[r_idx]:
            nums[current] = l_arr[l_idx]
            l_idx += 1
        else:
            nums[current] = r_arr[r_idx]
            r_idx += 1
        current += 1
    
    while l_idx < len(l_arr):
        nums[current] = l_arr[l_idx]
        l_idx += 1
        current += 1
    
    while r_idx < len(r_arr):
        nums[current] = r_arr[r_idx]
        r_idx += 1
        current += 1
    
    return nums

if __name__ == '__main__':
    s = Solution()

    assert s.missingNumber_xor([3,0,1]) == 2
    assert s.missingNumber_xor([0,1]) == 2
    assert s.missingNumber_xor([9,6,4,2,3,5,7,0,1]) == 8

    assert s.missingNumber_math([3,0,1]) == 2
    assert s.missingNumber_math([0,1]) == 2
    assert s.missingNumber_math([9,6,4,2,3,5,7,0,1]) == 8

    assert s.missingNumber_dict([3,0,1]) == 2
    assert s.missingNumber_dict([0,1]) == 2
    assert s.missingNumber_dict([9,6,4,2,3,5,7,0,1]) == 8

    assert s.missingNumber_sort([3,0,1]) == 2
    assert s.missingNumber_sort([0,1]) == 2
    assert s.missingNumber_sort([9,6,4,2,3,5,7,0,1]) == 8