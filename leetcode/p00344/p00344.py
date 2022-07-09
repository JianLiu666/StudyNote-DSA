from typing import List

class Solution:
    def reverseString(self, s: List[str]) -> None:
        head = 0
        tail = len(s)-1

        while head < tail:
            s[head], s[tail] = s[tail], s[head]
            head += 1
            tail -= 1

if __name__ == '__main__':
    s = Solution()

    list = ["h","e","l","l","o"]
    s.reverseString(list)
    assert list == ["o","l","l","e","h"]

    list = ["H","a","n","n","a","h"]
    s.reverseString(list)
    assert list == ["h","a","n","n","a","H"]
