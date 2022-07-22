class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        mapping = {}
        seen = {}

        for i in range(len(s)):
            if s[i] in mapping and mapping[s[i]] != t[i]:
                return False
            elif s[i] not in mapping and t[i] in seen:
                return False
            else:
                mapping[s[i]] = t[i]
                seen[t[i]] = True
        
        return True

if __name__ == '__main__':
    s = Solution()

    assert s.isIsomorphic("egg", "add") == True
    assert s.isIsomorphic("foo", "bar") == False
    assert s.isIsomorphic("paper", "title") == True
    assert s.isIsomorphic("abcdefghijklmnopqrstuvwxyzva", "abcdefghijklmnopqrstuvwxyzck") == False
    assert s.isIsomorphic("badc", "baba") == False

