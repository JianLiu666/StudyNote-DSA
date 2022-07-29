class Solution:
    def decodeString(self, s: str) -> str:
        result = ""
        
        stack = []
        
        head = 0
        status = 0
        for i in range(len(s)):
            # processing number
            if status == 0 and '0' <= s[i] <= '9':
                head = i
                status = 1
            if status == 1 and (s[i+1] < '0' or s[i+1] > '9'):
                stack.append(s[head:i+1])
                status = 0
            
            # processing letters
            if status == 0 and 'a' <= s[i] <= 'z':
                head = i
                status = 2
            if status == 2 and (i+1 == len(s) or (s[i+1] < 'a' or s[i+1] > 'z')):
                if len(stack) == 0:
                    result += s[head:i+1]
                else:
                    stack.append(s[head:i+1])
                status = 0
            
            # processing right square bracket
            if s[i] == ']':
                letters = stack.pop()
                while 'a' <= stack[len(stack)-1] <= 'z':
                    letters = stack.pop() + letters
                
                count = int(stack.pop())
                summed = ""
                for _ in range(count):
                    summed += letters
                
                if len(stack) == 0:
                    result += summed
                else:
                    stack.append(summed)
            
        return result
        


if __name__ == '__main__':
    s = Solution()

    assert s.decodeString("abc") == "abc"
    assert s.decodeString("3[a]2[bc]") == "aaabcbc"
    assert s.decodeString("3[a]10[bc]") == "aaabcbcbcbcbcbcbcbcbcbc"
