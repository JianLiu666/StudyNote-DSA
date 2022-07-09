class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        memo = {}
        for ch in s:
            if ch not in memo:
                memo[ch] = 1
            else:
                memo[ch] += 1

        for ch in t:
            if ch not in memo:
                return False
            if memo[ch] == 0:
                return False
            memo[ch] -= 1
        
        return True

if __name__ == '__main__':
    s = Solution()

    assert s.isAnagram("anagram", "nagaram") == True
    assert s.isAnagram("fat", "car") == False