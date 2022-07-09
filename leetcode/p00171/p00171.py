class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def titleToNumber(self, columnTitle: str) -> int:
        idx = len(columnTitle)-1
        num = 0
        for ch in columnTitle:
            num += (ord(ch)-64) * pow(26, idx)
            idx -= 1

        return num

if __name__ == '__main__':
    s = Solution()

    assert s.titleToNumber("A") == 1
    assert s.titleToNumber("AB") == 28
    assert s.titleToNumber("ZY") == 701