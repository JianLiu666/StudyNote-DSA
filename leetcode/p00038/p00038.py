class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(n)
    def countAndSay(self, n: int) -> str:
        s = "1"
        if n == 1:
            return s
        
        for i in range(1, n):
            tmp = ""
            
            keep = s[0]
            count = 1
            for ch in range(1, len(s)):
                if s[ch] == keep:
                    count += 1
                else:
                    tmp += str(count) + keep
                    keep = s[ch]
                    count = 1
            tmp += str(count) + keep
            s = tmp
        
        return s

if __name__ == '__main__':
    s = Solution()

    assert s.countAndSay(1) == "1"
    assert s.countAndSay(4) == "1211"