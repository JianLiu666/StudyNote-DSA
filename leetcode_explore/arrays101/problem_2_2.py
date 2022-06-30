'''
Merge Sorted Array
    You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
    Merge nums1 and nums2 into a single array sorted in non-decreasing order.
    The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
    To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:
    Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
    Output: [1,2,2,3,5,6]
    Explanation: 
        The arrays we are merging are [1,2,3] and [2,5,6].
        The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:
    Input: nums1 = [1], m = 1, nums2 = [], n = 0
    Output: [1]
    Explanation:
        The arrays we are merging are [1] and [].
        The result of the merge is [1].

Example 3:
    Input: nums1 = [0], m = 0, nums2 = [1], n = 1
    Output: [1]
    Explanation:
        The arrays we are merging are [] and [1].
        The result of the merge is [1].
        Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
 
Constraints:
    nums1.length == m + n
    nums2.length == n
    0 <= m, n <= 200
    1 <= m + n <= 200
    -10^9 <= nums1[i], nums2[j] <= 10^9
'''

'''
解題方向
    - nums1 已經預留好合併後的 array size 了, 所以我們可以給定一個 offset = len(nums)-1, 從尾巴往前位移
    - 比較 nums1[m] 跟 nums2[n] 當前位置的大小, 將較大的數字塞進 nums1[offset] 後再將位置往前為宜
        - e.g.
            nums1[m] > nums2[n]
            nums1[current] = nums1[m]
            offset -= 1
            m -= 1

'''

from typing import List

class Solution:
    # My Solution
    # Time Complexity: O(2(m+n))
    def merge_sol1(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        result = [None] * (m+n)
        current, mIdx, nIdx = 0, 0, 0

        while mIdx < m or nIdx < n:
            if mIdx >= m:
                result[current] = nums2[nIdx]
                nIdx += 1
            elif nIdx >= n:
                result[current] = nums1[mIdx]
                mIdx += 1
            else:
                if nums1[mIdx] <= nums2[nIdx]:
                    result[current] = nums1[mIdx]
                    mIdx += 1
                else:
                    result[current] = nums2[nIdx]
                    nIdx += 1
            
            current += 1

        for i in range(len(nums1)):
            nums1[i] = result[i]

    # LeetCode Solution
    # Time Complexity: O(m+n)
    def merge_sol2(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        current = len(nums1)-1
        while m > 0 and n > 0:
            if nums1[m-1] >= nums2[n-1]:
                nums1[current] = nums1[m-1]
                m -= 1
            else:
                nums1[current] = nums2[n-1]
                n -= 1
            
            current -= 1
        
        while n > 0:
            nums1[current] = nums2[n-1]
            current -= 1
            n -= 1

if __name__ == '__main__':
    s = Solution()

    list1 = [1,2,3,0,0,0]
    list2 = [1]
    list3 = [0]

    s.merge_sol1(list1, 3, [2,5,6], 3)
    assert list1 == [1,2,2,3,5,6]
    
    s.merge_sol1(list2, 1, [], 0)
    assert list2 == [1]

    s.merge_sol1(list3, 0, [1], 1)
    assert list3 == [1]

    list1 = [1,2,3,0,0,0]
    list2 = [1]
    list3 = [0]

    s.merge_sol2(list1, 3, [2,5,6], 3)
    assert list1 == [1,2,2,3,5,6]
    
    s.merge_sol2(list2, 1, [], 0)
    assert list2 == [1]

    s.merge_sol2(list3, 0, [1], 1)
    assert list3 == [1]