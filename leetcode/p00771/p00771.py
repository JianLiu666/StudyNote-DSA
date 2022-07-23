class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        mapping = set(jewels)

        count = 0
        for ch in stones:
            if ch in mapping:
                count += 1

        return count
    
if __name__ == '__main__':
    s = Solution()

    assert s.numJewelsInStones("aA", "aAAbbbb") == 3
    assert s.numJewelsInStones("z", "ZZ") == 0