from typing import List

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(1)
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if len(strs) == 1:
            return strs[0]

        size = len(strs[0])
        for s in strs:
            if len(s) < size:
                size = len(s)
        
        res = ""
        for ch in range(size):
            for i in range(1, len(strs)):
                if strs[i][ch] != strs[0][ch]:
                    return res
            res += strs[0][ch]
        
        return res

if __name__ == '__main__':
    s = Solution()

    assert s.longestCommonPrefix(["flower","flow","flight"]) == "fl"
    assert s.longestCommonPrefix(["dog","racecar","car"]) == ""