class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def countVowelPermutation(self, n: int) -> int:
        a, e, i, o, u = 1, 1, 1, 1, 1
        mod = 1000000007
        for _ in range(2, n+1, 1):
            na = (e+i+u) % mod
            ne = (a+i) % mod
            ni = (e+o) % mod
            no = i % mod
            nu = (i+o) % mod

            a, e, i, o, u = na, ne, ni, no, nu
        
        return (a+e+i+o+u) % mod

if __name__ == "__main__":
    s = Solution()

    assert s.countVowelPermutation(1) == 5
    assert s.countVowelPermutation(2) == 10
    assert s.countVowelPermutation(5) == 68