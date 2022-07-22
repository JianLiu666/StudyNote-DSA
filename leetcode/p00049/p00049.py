from typing import List

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(n)
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams_group = {}
        group_index = 0
        
        res = []
        for s in strs:
            key = counting_sort(s)
            if key in anagrams_group:
                res[anagrams_group[key]].append(s)
            else:
                res.append([s])
                anagrams_group[key] = group_index
                group_index += 1

        return res

# Time Complexity: O(n)
# Space Complexity: O(1)
def counting_sort(s: str) -> str:
    mapping = [0] * 26

    for ch in s:
        mapping[ord(ch)-97] += 1
    
    res = ""
    for idx, count in enumerate(mapping):
        if count != 0:
            res += chr(idx+97) * count

    return res

if __name__ == '__main__':
    s = Solution()

    assert s.groupAnagrams(["eat","tea","tan","ate","nat","bat"]) == [["eat","tea","ate"],["tan","nat"],["bat"]]
    assert s.groupAnagrams([""]) == [[""]]
    assert s.groupAnagrams(["a"]) == [["a"]]