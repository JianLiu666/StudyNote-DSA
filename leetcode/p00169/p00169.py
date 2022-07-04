
from typing import List

class Solution:
    
    # Boyer-Moore Voting Algorithm
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def majorityElement(self, nums: List[int]) -> int:
        answer = 0
        count = 0

        for num in nums:
            if count == 0:
                answer = num
                count = 1
            elif answer == num:
                count += 1
            else:
                count -= 1
        
        return answer

    # Divide and Conquer
    # Time Complexity: O(nlogn)
    # Space Complexity: O(logn)
    def majorityElement_dac(self, nums: List[int]) -> int:
        def recursive(head: int, tail: int):
            if head == tail:
                return nums[head]
            
            mid = (tail - head) // 2 + head
            left = recursive(head, mid)
            right = recursive(mid+1, tail)

            if left == right:
                return left

            left_count = sum(1 for i in range(head, tail+1, 1) if nums[i] == left)
            right_count = sum(1 for i in range(head, tail+1, 1) if nums[i] == right)

            return left if left_count > right_count else right

        return recursive(0, len(nums)-1)

    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def majorityElement_dict(self, nums: List[int]) -> int:
        seen = {}

        max_num = 0
        count = 0
        for num in nums:
            if num in seen:
                seen[num] += 1
            else:
                seen[num] = 1

            if seen[num] > count:
                max_num = num
                count = seen[num]
        
        return max_num


if __name__ == '__main__':
    s = Solution()

    assert s.majorityElement([3,2,3]) == 3
    assert s.majorityElement([2,2,1,1,1,2,2]) == 2

    assert s.majorityElement_dac([3,2,3]) == 3
    assert s.majorityElement_dac([2,2,1,1,1,2,2]) == 2

    assert s.majorityElement_dict([3,2,3]) == 3
    assert s.majorityElement_dict([2,2,1,1,1,2,2]) == 2