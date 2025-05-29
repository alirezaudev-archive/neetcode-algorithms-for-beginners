from typing import List

# https://leetcode.com/problems/course-schedule/

class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = {i: [] for i in range(numCourses)}
        for course, prereq in prerequisites:
            graph[course].append(prereq)

        visited = [False] * numCourses

        def dfs(node: int) -> bool:
            if visited[node]:
                return False
            if not graph[node]:
                return True

            visited[node] = True

            for neighbor in graph[node]:
                if not dfs(neighbor):
                    return False

            visited[node] = False
            graph[node] = []
            return True

        for i in range(numCourses):
            if not dfs(i):
                return False

        return True
