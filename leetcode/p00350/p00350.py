from typing import List

class Solution:

    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def intersect_dict(self, nums1: List[int], nums2: List[int]) -> List[int]:
        dict1 = {n: nums1.count(n) for n in set(nums1)}
        dict2 = {n: nums2.count(n) for n in set(nums2)}

        result = []
        for num in dict1:
            if num in dict2:
                result += [num] * min(dict1[num], dict2[num])

        return result

    # Time Complexity: O(nlogn)
    # Space Complexity: O(n)
    def intersect_sort(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # using Python's buitin fucntion: timsort()
        # time complexity: O(nlogn)
        nums1.sort()
        nums2.sort()

        ptr1 = 0
        ptr2 = 0

        result = []
        while ptr1 < len(nums1) and ptr2 < len(nums2):
            if nums1[ptr1] == nums2[ptr2]:
                result.append(nums1[ptr1])
                ptr1 += 1
                ptr2 += 1
            elif nums1[ptr1] < nums2[ptr2]:
                ptr1 += 1
            else:
                ptr2 += 1
        
        return result

if __name__ == '__main__':
    s = Solution()

    assert sorted(s.intersect_dict([1,2,2,1], [2,2])) == [2,2]
    assert sorted(s.intersect_dict([4,9,5], [9,4,9,8,4])) == [4,9]

    assert sorted(s.intersect_sort([1,2,2,1], [2,2])) == [2,2]
    assert sorted(s.intersect_sort([4,9,5], [9,4,9,8,4])) == [4,9]