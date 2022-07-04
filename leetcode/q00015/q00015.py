from typing import List

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(1)
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return []
        
        nums.sort()

        result = []
        for i in range(0, len(nums)-2, 1):
            if nums[i] > 0:
                break
            
            if i > 0 and nums[i] == nums[i-1]:
                continue

            head = i+1
            tail = len(nums)-1
            while head < tail:
                threeSum = nums[i] + nums[head] + nums[tail]
                if threeSum < 0:
                    head += 1
                elif threeSum > 0:
                    tail -= 1
                else:
                    result.append([nums[i], nums[head], nums[tail]])
                    head += 1

                    while nums[head] == nums[head-1] and head < tail:
                        head += 1
        
        return result

if __name__ == '__main__':
    s = Solution()

    assert s.threeSum([-4,-1,-1,0,1,2]) == [[-1,-1,2],[-1,0,1]]
    assert s.threeSum([-4,-3,-2,6,7]) == [[-4,-3,7],[-4,-2,6]]
    assert s.threeSum([]) == []
    assert s.threeSum([0]) == []
    assert s.threeSum([0,0,0,0]) == [[0,0,0]]