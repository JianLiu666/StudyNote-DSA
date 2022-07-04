'''
Evaluate Postfix Expression Using a Stack

Problem Statement:
    The usual convention followed in mathematics is the infix expression. 
    Operators like + and * appear between the two numbers involved in the calculation:
        6 + 3 * 8 - 4
   
    Another convention is the postfix expression where the operators appear after the two numbers involved in the expression. 
    In postfix, the expression written above will be presented as:
        6 3 8 * + 4 -

    The two digits preceding an operator will be used with that operator
        1. From the first block of digits 6 3 8, we pick the last two which are 3 and 8.
        2. Reading the operators from left to right, the first one is *. The expression now becomes 3 * 8
        3. The next number is 6 while the next operator is +, so we have 6 + 8 * 3.
        4. The value of this expression is followed by 4, which is right before -. Hence we have 6 + 8 * 3 - 4.

    Implement a function called evaluatePostFix() that will compute a postfix expression given to it as a string.

Input:
    A string containing a postfix mathematic expression. Each digit is considered to be a separate number, i.e., there are no double digit numbers.

Sample Input:
    exp = "921 * - 8 - 4 +" # 9 - 2 * 1 - 8 + 4

Output:
    A result of the given postfix expression.

Sample Output:
    3
'''

# My Solution
# Time Complexity: O(n)
class Stack:
    def __init__(self):
        self.main = []
        self.length = 0

    def is_empty(self):
        return self.length == 0

    def peek(self):
        if self.is_empty():
            return None
        
        return self.main[-1]
    
    def size(self):
        return self.length

    def push(self, data):
        self.length += 1
        self.main.append(data)

    def pop(self):
        if self.is_empty():
            return None
        
        data = self.main.pop()
        self.length -= 1
        return data

def evaluate_post_fix_sol1(exp):
    stack = Stack()
    try:
        for char in exp:
            if char.isdigit():
                stack.push(char)
            else:
                right = stack.pop()
                left = stack.pop()
                stack.push(str(eval(left + char + right)))
        
        return int(float(stack.pop()))
    except TypeError:
        return "Invalid Sequence"

print(evaluate_post_fix_sol1("921*-8-4+"))
print(evaluate_post_fix_sol1("921*-8--4+"))