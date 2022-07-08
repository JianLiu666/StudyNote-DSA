
answer = 0

class Solution:
    def guessNumber(self, n: int) -> int:
        head = 1
        tail = n

        while True:
            num = head + (tail-head) // 2
            
            res = guess(num)
            if res == 0:
                return num
            elif res == 1:
                head = num + 1
            else:
                tail = num - 1
            
def guess(n: int) -> int:
    if n == answer:
        return 0
    elif n < answer:
        return 1
    else:
        return -1

if __name__ == '__main__':
    s = Solution()

    answer = 6
    assert s.guessNumber(10) == answer

    answer = 1
    assert s.guessNumber(1) == answer

    answer = 1
    assert s.guessNumber(2) == answer