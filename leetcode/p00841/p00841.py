from typing import List
from collections import deque

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        def visitRoom(roomId: int) -> int:
            visitedCount = 0
            for key in rooms[roomId]:
                if visited[key] == False:
                    visited[key] = True
                    visitedCount += 1
                    q.append(key)
            
            return visitedCount

        visited = [False] * len(rooms)
        visited[0] = True
        visitedCount = 1

        q = deque()
        visitedCount += visitRoom(0)

        while q:
            roomId = q.popleft()
            visitedCount += visitRoom(roomId)

        return visitedCount == len(rooms)

if __name__ == '__main__':
    s = Solution()

    assert s.canVisitAllRooms([[1],[2],[3],[]]) == True
    assert s.canVisitAllRooms([[1,3],[3,0,1],[2],[0]]) == False