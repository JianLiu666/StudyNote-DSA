class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def strStr(self, haystack: str, needle: str) -> int:
        haystack_size = len(haystack)
        needle_size = len(needle)

        if needle_size == 0:
            return 0
        if haystack_size < needle_size:
            return -1

        for i in range(haystack_size):
            if haystack[i] == needle[0]:
                if i+needle_size <= haystack_size:
                    if haystack[i:i+needle_size] == needle:
                        return i
                else:
                    return -1
        
        return -1

if __name__ == '__main__':
    s = Solution()

    assert s.strStr("hello", "ll") == 2
    assert s.strStr("aaaaa", "bba") == -1
    assert s.strStr("mississippi", "issip") == 4