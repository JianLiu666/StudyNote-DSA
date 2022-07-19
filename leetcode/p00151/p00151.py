class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def reverseWords(self, s: str) -> str:
        res = ""
        tail = 0

        # 0: search word tail, 1: search word head
        state = 0
        for i in range(len(s)-1, -1, -1):
            if state == 0 and s[i] != ' ' and (i == len(s)-1 or s[i+1] == ' '):
                tail = i
                state = 1
            if state == 1 and s[i] != ' ' and (i == 0 or s[i-1] == ' '):
                if res != "":
                    res += " " + s[i:tail+1]
                else:
                    res = s[i:tail+1]
                state = 0
        
        return res
    
    def reverseWords_pythonic(self, s: str) -> str:
        print(' '.join(s.split()[::-1]))

if __name__ == '__main__':
    s = Solution()

    assert s.reverseWords("the sky is blue") == "blue is sky the"
    assert s.reverseWords("  hello world  ") == "world hello"
    assert s.reverseWords("a good   example") == "example good a"

    assert s.reverseWords_pythonic("the sky is blue") == "blue is sky the"
    assert s.reverseWords_pythonic("  hello world  ") == "world hello"
    assert s.reverseWords_pythonic("a good   example") == "example good a"