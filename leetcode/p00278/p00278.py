# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def firstBadVersion(self, n: int) -> int:
        left = 1
        right = n

        while left != right:
            mid = left + (right-left) // 2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid + 1
        
        return left

badVersion = 0
def isBadVersion(version: int) -> bool:
    return version >= badVersion

if __name__ == '__main__':
    s = Solution()

    badVersion = 4
    assert s.firstBadVersion(5) == 4

    badVersion = 1
    assert s.firstBadVersion(1) == 1

    badVersion = 1702766719
    assert s.firstBadVersion(2126753390) == 1702766719