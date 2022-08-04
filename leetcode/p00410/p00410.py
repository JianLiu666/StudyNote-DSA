from typing import List

class Solution:
    # Time Complexity: O(nlogn)
    # Space Complexity: O(1)
    def splitArray(self, nums: List[int], m: int) -> int:
        
        def cut(cuts: int, maximum: int) -> bool:
            acc = 0
            for num in nums:
                if acc+num <= maximum:
                    acc += num
                else:
                    acc = num
                    cuts -= 1
                    if cuts < 0:
                        return False
            
            return True
        
        low, high = 0, 0
        for num in nums:
            low = max(low, num)
            high += num
        
        while low < high:
            mid = low + (high-low)//2
            if cut(m-1, mid):
                high = mid
            else:
                low = mid + 1
        
        return low

if __name__ == '__main__':
    s = Solution()

    s.splitArray([7,2,5,10,8], 2) == 18
    s.splitArray([1,2,3,4,5], 2) == 9
    s.splitArray([1,4,4], 3) == 4