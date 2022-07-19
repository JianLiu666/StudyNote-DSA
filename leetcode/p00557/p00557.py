class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(n)
    def reverseWords(self, s: str) -> str:
        return ' '.join([word[::-1] for word in s.split()])

if __name__ == '__main__':
    s = Solution()

    assert s.reverseWords("Let's take LeetCode contest") == "s'teL ekat edoCteeL tsetnoc"
    assert s.reverseWords("God Ding") == "doG gniD"