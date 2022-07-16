from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        m, mx = 0, len(mat)
        n, nx = 0, len(mat[0])

        # 0: up-right, 1:down-left
        direction = 0

        res = []
        while m < mx and n < nx:
            res.append(mat[m][n])

            if direction == 0:
                if m - 1 >= 0 and n + 1 < nx:
                    m -= 1
                    n += 1
                else:
                    direction = 1
                    if n + 1 < nx:
                        n += 1
                    else:
                        m += 1
            else:
                if m + 1 < mx and n - 1 >= 0:
                    m += 1
                    n -= 1
                else:
                    direction = 0
                    if m + 1 < mx:
                        m += 1
                    else:
                        n += 1

        return res

if __name__ == '__main__':
    s = Solution()

    assert s.findDiagonalOrder([[1,2,3],[4,5,6],[7,8,9]]) == [1,2,4,7,5,3,6,8,9]
    assert s.findDiagonalOrder([[1,2],[3,4]]) == [1,2,3,4]
    assert s.findDiagonalOrder([[2,3]]) == [2,3]