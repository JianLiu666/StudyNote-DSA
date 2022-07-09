class Solution:
    def mySqrt(self, x: int) -> int:
        if x <= 1:
            return x
        
        head = 1
        tail = x

        while head < tail:
            mid = head + (tail-head) // 2
            ans = mid * mid

            if ans == x:
                return mid
            elif ans > x:
                tail = mid - 1
            else:
                head = mid + 1
        
        if head * head > x:
            return head - 1
        return head

if __name__ == '__main__':
    s = Solution()

    assert s.mySqrt(4) == 2
    assert s.mySqrt(8) == 2
    assert s.mySqrt(808909662) == 28441