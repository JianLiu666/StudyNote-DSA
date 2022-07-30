from typing import List

class Solution:
    # Time Complexity: O(n+k), where k is the length of words2
    # Space Complexity: O(1)
    def wordSubsets(self, words1: List[str], words2: List[str]) -> List[str]:
        def counter(word: str) -> List[int]:
            frequency = [0] * 26
            for ch in word:
                frequency[ord(ch)-97] += 1
            return frequency
        
        subsetMap = [0] * 26
        for word in words2:
            letters = counter(word)
            for i in range(26):
                subsetMap[i] = max(letters[i], subsetMap[i])
        
        res = []
        for word in words1:
            supersetMap = counter(word)
            isSubset = True
            for i in range(26):
                if subsetMap[i] > supersetMap[i]:
                    isSubset = False
                    break
            if isSubset:
                res.append(word)
        
        return res

if __name__ == "__main__":
    s = Solution()

    assert s.wordSubsets(
        ["amazon", "apple", "facebook", "google", "leetcode"],
        ["e", "o"]) == ["facebook", "google", "leetcode"]
    
    assert s.wordSubsets(
        ["amazon", "apple", "facebook", "google", "leetcode"],
        ["l", "e"]) == ["apple", "google", "leetcode"]