from typing import List

class Solution:
    # Time Complexity: O(n^2)
    # Space Complexity: O(n)
    def numFactoredBinaryTrees(self, arr: List[int]) -> int:
        arr.sort()
        
        memo = {val: idx for idx, val in enumerate(arr)}
        
        counter = [1 for _ in range(len(arr))]
        mod = 1000000007
        
        for current in range(len(arr)):
            for left in range(current):
                if arr[current] % arr[left] != 0:
                    continue
                
                if arr[current]//arr[left] in memo:
                    right = memo[arr[current]//arr[left]]
                    counter[current] += (counter[left] * counter[right]) % mod
        
        return sum(counter) % mod


if __name__ == '__main__':
    s = Solution()

    assert s.numFactoredBinaryTrees([2,4]) == 3
    assert s.numFactoredBinaryTrees([2,4,5,10]) == 7