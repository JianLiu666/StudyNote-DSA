class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        
        seen = {}
        max_distance = 0
        
        start = 0
        for cursor, ch in enumerate(s):
            if ch in seen and seen[ch] >= start:
                start = seen[ch] + 1
            elif cursor - start > max_distance:
                max_distance = cursor - start
            
            seen[ch] = cursor
        
        return max_distance + 1

if __name__ == '__main__':
    s = Solution()

    assert s.lengthOfLongestSubstring("abcabcbb") == 3
    assert s.lengthOfLongestSubstring("abcadefg") == 7
    assert s.lengthOfLongestSubstring("") == 0
    assert s.lengthOfLongestSubstring("a") == 1
    assert s.lengthOfLongestSubstring("tmmzuxtabc") == 8
    assert s.lengthOfLongestSubstring("tmmzuxt") == 5
