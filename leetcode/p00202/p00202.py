class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def isHappy(self, n: int) -> bool:
        slow = n
        fast = n
        while fast != 1 and getNext(fast) != 1:
            slow = getNext(slow)
            fast = getNext(getNext(fast))

            if slow == fast:
                return False
        
        return True
    
def getNext(n: int) -> int:
    memo = {
        0:0,
        1:1,
        2:4,
        3:9,
        4:16,
        5:25,
        6:36,
        7:49,
        8:64,
        9:81,
    }

    sum = 0
    while n != 0:
        n, remainder = divmod(n, 10)
        sum += memo[remainder]
    
    return sum

if __name__ == '__main__':
    s = Solution()

    assert s.isHappy(19) == True
    assert s.isHappy(2) == False
    assert s.isHappy(10) == True