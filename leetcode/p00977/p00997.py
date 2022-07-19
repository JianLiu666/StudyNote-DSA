from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def sortedSquares(self, nums: List[int]) -> List[int]:
        head = 0
        tail = len(nums)-1

        result = [None] * len(nums)
        current = len(result)-1
        while current >= 0:
            if abs(nums[head]) >= abs(nums[tail]):
                result[current] = pow(nums[head], 2)
                head += 1
            else:
                result[current] = pow(nums[tail], 2)
                tail -= 1
            current -= 1
        
        return result

if __name__ == '__main__':
    s = Solution()

    s.sortedSquares([-4,-1,0,3,10]) == [0,1,9,16,100]
    s.sortedSquares([-7,-3,2,3,11]) == [4,9,9,49,121]

