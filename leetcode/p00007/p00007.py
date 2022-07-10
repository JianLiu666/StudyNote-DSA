class Solution:
    def reverse_recursion(self, x: int) -> int:
        sig = 1
        if x < 0:
            sig = -1
            x *= -1
        
        res = 0
        while x != 0:
            num = x % 10
            res = res*10 + num
            x //= 10
        
        if res > pow(2, 31)-1:
            return 0
        
        return res * sig

    def reverse_string(self, x: int) -> int:
        s = str(x)
        head = 0
        
        r = ""
        if s[head] == '-':
            head += 1
            r += '-'

        for i in range(len(s)-1, head-1, -1):
            r += s[i]
        
        if int(r) > pow(2, 31)-1:
            return 0
        
        return int(r)

if __name__ == '__main__':
    s = Solution()

    assert s.reverse_recursion(123) == 321
    assert s.reverse_recursion(-123) == -321
    assert s.reverse_recursion(120) == 21
    assert s.reverse_recursion(1234567899) == 0

    assert s.reverse_string(123) == 321
    assert s.reverse_string(-123) == -321
    assert s.reverse_string(120) == 21
    assert s.reverse_string(1234567899) == 0