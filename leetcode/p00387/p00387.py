from typing import Counter, OrderedDict

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def firstUniqChar(self, s: str) -> int:
        # Ordered Dict will save the characters it encounters in same sequence as the original string.
        for idx, cnt  in OrderedDict(Counter(s)).items():
            if cnt == 1:
                return s.index(idx)
        
        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.firstUniqChar("leetcode") == 0
    assert s.firstUniqChar("loveleetcode") == 2
    assert s.firstUniqChar("aabb") == -1