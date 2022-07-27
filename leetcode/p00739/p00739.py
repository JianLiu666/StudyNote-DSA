from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res = [0] * len(temperatures)
        stack = []

        for cur, temp in enumerate(temperatures):
            # using stack[len(stack)-1] instead of stack[-1]
            while len(stack) > 0 and temp > temperatures[stack[len(stack)-1]]:
                prev = stack.pop()
                res[prev] = cur - prev
            
            stack.append(cur)
        
        return res

if __name__ == '__main__':
    s = Solution()

    assert s.dailyTemperatures([73,74,75,71,69,72,76,73]) == [1,1,4,2,1,1,0,0]
    assert s.dailyTemperatures([30,40,50,60]) == [1,1,1,0]
    assert s.dailyTemperatures([30,60,90]) == [1,1,0]