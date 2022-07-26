from typing import List
from collections import deque

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def openLock(self, deadends: List[str], target: str) -> int:
        if target == "0000":
            return 0
        seen = set(deadends)
        if "0000" in seen:
            return -1
        
        q = deque(["0000"])
        step = 0

        while q:
            size = len(q)
            for _ in range(size):
                s = q.popleft()
                for i, ch in enumerate(s):
                    down, up = str(int(ch)-1), str(int(ch)+1)
                    if ch == '0':
                        down, up = '9', '1'
                    elif ch == '9':
                        down, up = '8', '0'
                    
                    wheels = s[:i] + down + s[i+1:]
                    if wheels not in seen:
                        seen.add(wheels)
                        q.append(wheels)
                    
                    wheels = s[:i] + up + s[i+1:]
                    if wheels not in seen:
                        seen.add(wheels)
                        q.append(wheels)
                
            step += 1
            if target in seen:
                return step
        
        return -1


if __name__ == '__main__':
    s = Solution()

    assert s.openLock(["0201","0101","0102","1212","2002"], "0202") == 6
    assert s.openLock(["8888"], "0009") == 1
    assert s.openLock(["8887","8889","8878","8898","8788","8988","7888","9888"], "8888") == -1
    assert s.openLock(["0000"], "8888") == -1
    assert s.openLock(["0201","0101","0102","1212","2002"], "0000") == 0
