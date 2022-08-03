class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def isPerfectSquare(self, num: int) -> bool:
        if num == 1:
            return True

        low, high = 1, num
        while low <= high:
            mid = low + (high-low)//2
            square = mid * mid
            if square == num:
                return True
            elif square > num:
                high = mid + 1
            else:
                low = mid + 1

        return False

if __name__ == '__main__':
    s = Solution()

    assert s.isPerfectSquare(16) == True
    assert s.isPerfectSquare(14) == False