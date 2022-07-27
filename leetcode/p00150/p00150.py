from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for symbol in tokens:
            if symbol in ["+", "-", "*", "/"]:
                stack.append(operate(symbol, stack.pop(), stack.pop()))
            else:
                stack.append(int(symbol))

        return stack.pop()

def operate(symbol: str, right: int, left: int) -> int:
    if symbol == "+":
        return left + right   
    if symbol == "-":
        return left - right
    if symbol == "*":
        return left * right

    return int(left / right)

if __name__ == '__main__':
    s = Solution()

    assert s.evalRPN(["2","1","+","3","*"]) == 9
    assert s.evalRPN(["4","13","5","/","+"]) == 6
    assert s.evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"]) == 22