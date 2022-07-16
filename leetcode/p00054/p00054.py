from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        m_min, m_cur, m_max = -1, 0, len(matrix)
        n_min, n_cur, n_max = -1, 0, len(matrix[0])

        res = [0] * m_max * n_max
        count = 0

        # 0: right, 1: down, 2: left, 3: up
        direction = 0
        while count < len(res):
            res[count] = matrix[m_cur][n_cur]
            count += 1

            if direction == 0:
                if n_cur + 1 < n_max:
                    n_cur += 1
                else:
                    direction = 1
                    m_cur += 1
                    m_min += 1
            elif direction == 1:
                if m_cur + 1 < m_max:
                    m_cur += 1
                else:
                    direction = 2
                    n_cur -= 1
                    n_max -= 1
            elif direction == 2:
                if n_cur - 1 > n_min:
                    n_cur -= 1
                else:
                    direction = 3
                    m_cur -= 1
                    m_max -= 1
            else:
                if m_cur - 1 > m_min:
                    m_cur -= 1
                else:
                    direction = 0
                    n_cur += 1
                    n_min += 1
        
        return res

if __name__ == '__main__':
    s = Solution()

    assert s.spiralOrder([[1,2,3],[4,5,6],[7,8,9]]) == [1,2,3,6,9,8,7,4,5]
    assert s.spiralOrder([[1,2,3,4],[5,6,7,8],[9,10,11,12]]) == [1,2,3,4,8,12,11,10,9,5,6,7]