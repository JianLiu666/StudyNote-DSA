from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def singleNumber(self, nums: List[int]) -> int:
        ans = 0
        for num in nums:
            ans ^= num
        
        return ans
    
    # Time Compelxity: O(2n) -> O(n)
    # Time Complexity: O(n/2) -> O(n)
    def singleNumber_dsa(self, nums: List[int]) -> int:
        seen = {}

        for num in nums:
            if num in seen:
                seen[num] += 1
            else:
                seen[num] = 1
        
        for k, v in seen.items():
            if v == 1:
                return k

if __name__ == '__main__':
    s = Solution()

    assert s.singleNumber([2,2,1]) == 1
    assert s.singleNumber([4,1,2,1,2]) == 4
    assert s.singleNumber([1]) == 1

    assert s.singleNumber_dsa([2,2,1]) == 1
    assert s.singleNumber_dsa([4,1,2,1,2]) == 4
    assert s.singleNumber_dsa([1]) == 1