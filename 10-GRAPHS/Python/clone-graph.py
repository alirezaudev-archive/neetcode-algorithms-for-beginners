from typing import Optional


# https://leetcode.com/problems/clone-graph/

class Node:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        if node is None:
            return None

        visited = dict()

        def dfs(n: Node):
            if n in visited:
                return visited[n]

            clone = Node(val=n.val)
            visited[n] = clone

            for nei in n.neighbors:
                clone.neighbors.append(dfs(nei))

            return clone

        return dfs(node)

