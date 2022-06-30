from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def plusOne(self, digits: List[int]) -> List[int]:
        result = []
        
        digits[-1] += 1
        carried = 0
        for i in range(len(digits)-1, -1, -1):
            num = digits[i]
            if digits[i] == 10:
                carried = 1
                num = 0
            else:
                carried = 0
            result = [num] + result
        
        if carried == 1:
            result = [1] + result
        
        return result


if __name__ == '__main__':
    s = Solution()

    assert s.plusOne([1,2,3]) == [1,2,4]
    assert s.plusOne([4,3,2,1]) == [4,3,2,2]
    assert s.plusOne([9]) == [1,0]