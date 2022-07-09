class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        for ch in s:
            if ch in '([{':
                stack.append(ch)
            elif ch in ')]}' and len(stack) == 0:
                return False
            elif ch == ')' and stack[-1] != '(':
                return False
            elif ch == ']' and stack[-1] != '[':
                return False
            elif ch == '}' and stack[-1] != '{':
                return False
            else:
                stack.pop()
        
        return len(stack) == 0

if __name__ == '__main__':
    s = Solution()

    assert s.isValid("()") == True
    assert s.isValid("()[]{}") == True
    assert s.isValid("(]") == False
    assert s.isValid("([)]") == False