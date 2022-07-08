from typing import List

class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        result = []
        for i in range(1, n+1, 1):
            if i % 15 == 0:
                result.append("FizzBuzz")
            elif i % 3 == 0:
                result.append("Fizz")
            elif i % 5 == 0:
                result.append("Buzz")
            else:
                result.append(str(i))
                
        return result

if __name__ == '__main__':
    s = Solution()

    assert s.fizzBuzz(3) == ["1","2","Fizz"]
    assert s.fizzBuzz(5) == ["1","2","Fizz","4","Buzz"]
    assert s.fizzBuzz(15) == ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]