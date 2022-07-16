class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        if len(ransomNote) > len(magazine):
            return False
        
        memo = {}
        for ch in magazine:
            if ch in memo:
                memo[ch] += 1
            else:
                memo[ch] = 1
        
        for ch in ransomNote:
            if ch in memo and memo[ch] > 0:
                memo[ch] -= 1
            else:
                return False
        
        return True

if __name__ == '__main__':
    s = Solution()

    assert s.canConstruct("a", "b") == False
    assert s.canConstruct("aa", "ab") == False
    assert s.canConstruct("aa", "aab") == True