from typing import List

class Solution:
    # Time Complexity: O(nk), wherer k is the length of the word
    # Space Complexity: O(n+2k), where k is the length of the word
    def findAndReplacePattern(self, words: List[str], pattern: str) -> List[str]:
        def match(word) -> bool:
            mapping = {}
            for i, j in zip(pattern, word):
                if mapping.setdefault(i, j) != j:
                    return False
            return len(set(mapping.values())) == len(mapping.values())

        return filter(match, words)

if __name__ == '__main__':
    s = Solution()

    assert s.findAndReplacePattern(["abc", "deq", "mee", "aqq", "dkd", "ccc"], "abb") == ["mee", "aqq"]
    assert s.findAndReplacePattern(["a", "b", "c"], "a") == ["a", "b", "c"]