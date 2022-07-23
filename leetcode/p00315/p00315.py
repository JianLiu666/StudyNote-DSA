from typing import List

class Solution:
    # Time Complexity: O(nlogn)
    # Space Complexity: O(n)
    def countSmaller(self, nums: List[int]) -> List[int]:

        def merge_and_count(nums: List[List[int]]) -> List[List[int]]:
            if len(nums) == 1:
                return nums
            
            mid = len(nums)//2
            l_arr = merge_and_count(nums[:mid])
            r_arr = merge_and_count(nums[mid:])

            res = []
            l_cur = 0
            r_cur = 0

            accumulator = 0
            while l_cur < len(l_arr) and r_cur < len(r_arr):
                if l_arr[l_cur][0] > r_arr[r_cur][0]:
                    accumulator += 1
                    res.append(r_arr[r_cur])
                    r_cur += 1
                else:
                    # 先將目前為止的累加結果寫回 left element
                    l_arr[l_cur][2] += accumulator
                    res.append(l_arr[l_cur])
                    l_cur += 1
            
            while l_cur < len(l_arr):
                # 將最後的累加結果寫回所有還在 left array 的 elements
                l_arr[l_cur][2] += accumulator
                res.append(l_arr[l_cur])
                l_cur += 1
            
            while r_cur < len(r_arr):
                res.append(r_arr[r_cur])
                r_cur += 1
            
            return res

        counter = merge_and_count([[value, original_index, 0] for original_index, value in enumerate(nums)])
        res = [0] * len(counter)
        for info in counter:
            res[info[1]] = info[2]
        
        return res


if __name__ == '__main__':
    s = Solution()

    assert s.countSmaller([5,2,6,1]) == [2,1,1,0]
    assert s.countSmaller([-1]) == [0]
    assert s.countSmaller([-1,-1]) == [0,0]