from typing import List

class Solution:

    # Time Complexity: O(m+n)
    # Space Complexity: O(1)
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        current = len(nums1)-1
        while n > 0 and m > 0:
            if nums1[m-1] > nums2[n-1]:
                nums1[current] = nums1[m-1]
                m -= 1
            else:
                nums1[current] = nums2[n-1]
                n -= 1
            current -= 1

        while n > 0:
            nums1[current] = nums2[n-1]
            n -= 1
            current -= 1


if __name__ == '__main__':
    s = Solution()

    list1 = [1,2,3,0,0,0]
    s.merge(list1, 3, [2,5,6], 3)
    assert list1 == [1,2,2,3,5,6]

    list2 = [1]
    s.merge(list2, 1, [], 0)
    assert list2 == [1]

    list3 = [0]
    s.merge(list3, 0, [1], 1)
    assert list3 == [1]