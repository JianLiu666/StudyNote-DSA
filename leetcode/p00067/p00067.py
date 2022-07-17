class Solution:
    
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def addBinary_bruteforce(self, a: str, b: str) -> str:
        ia = len(a)-1
        ib = len(b)-1
        carried = 0

        res = ""
        while ia > -1 or ib > -1 or carried > 0:
            if ia > -1:
                carried += int(a[ia])
                ia -= 1
            if ib > -1:
                carried += int(b[ib])
                ib -= 1
            
            res = str(carried%2) + res
            carried //= 2
        
        return res

    # Time Complexity: O()
    # Space Complexity: O()
    def addBinary_builtin(self, a: str, b: str) -> str:
        return bin(int(a,2)+int(b,2))[2:]

if __name__ == '__main__':
    s = Solution()

    assert s.addBinary_bruteforce("11", "1") == "100"
    assert s.addBinary_bruteforce("1010", "1011") == "10101"

    assert s.addBinary_builtin("11", "1") == "100"
    assert s.addBinary_builtin("1010", "1011") == "10101"