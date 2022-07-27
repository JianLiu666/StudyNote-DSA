"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Node:
    def __init__(self, val=0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def cloneGraph(self, node: 'Node') -> 'Node':
        visited = {}
        head = Node()
        self.dfs(visited, node, head)

        return head
    
    def dfs(self, visited: dict, src: 'Node', dst: 'Node'):
        dst.val = src.val
        visited[src] = dst

        for node in src.neighbors:
            if node in visited:
                dst.neighbors.append(visited[node])
                continue
            
            new_node = Node()
            dst.neighbors.append(new_node)
            self.dfs(visited, node, new_node)
