from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def rotate_reversion(self, nums: List[int], k: int) -> None:
        def reverse(nums: List[int], left: int, right: int) -> None:
            if left >= right:
                return 
            
            while left < right:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
                right -= 1
        
        reverse(nums, 0, len(nums)-1)
        reverse(nums, 0, k-1)
        reverse(nums, k, len(nums)-1)
    
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def rotate_cyclic(self, nums: List[int], k: int) -> None:
        k %= len(nums)
        count = 0
        start = 0
        while count < len(nums):
            current = start
            prev = nums[start] 

            while True:
                next = (current + k) % len(nums)
                tmp = nums[next]
                nums[next] = prev
                prev = tmp
                current = next
                count += 1
            
                if start == current:
                    break
            
            start += 1
        
if __name__ == '__main__':
    s = Solution()

    s.rotate_reversion([1,2,3,4,5,6,7], 3) == [5,6,7,1,2,3,4]
    s.rotate_reversion([-1,-100,3,99], 2) == [3,99,-1,-100]

    s.rotate_cyclic([1,2,3,4,5,6,7], 3) == [5,6,7,1,2,3,4]
    s.rotate_cyclic([-1,-100,3,99], 2) == [3,99,-1,-100]