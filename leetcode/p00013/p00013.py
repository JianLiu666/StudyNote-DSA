class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def romanToInt(self, s: str) -> int:
        memo = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000,
        }

        ans = 0
        previous = 0
        for ch in s:
            num = memo[ch]
            if num <= previous:
                ans += num
            else:
                ans += num - (previous*2)
            
            previous = num

        return ans

if __name__ == '__main__':
    s = Solution()

    assert s.romanToInt("III") == 3
    assert s.romanToInt("LVIII") == 58
    assert s.romanToInt("MCMXCIV") == 1994