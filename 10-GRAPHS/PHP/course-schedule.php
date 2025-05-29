<?php

// https://leetcode.com/problems/course-schedule/

class Solution
{
    private array $graph = [];

    private array $visited = [];

    /**
     * @param Integer $numCourses
     * @param Integer[][] $prerequisites
     * @return Boolean
     */
    function canFinish(int $numCourses, array $prerequisites): bool
    {
        foreach ($prerequisites as $pair) {
            $course = $pair[0];
            $this->graph[$course][] = $pair[1];
        }

        for ($i = 0; $i < $numCourses; $i++) {
            if (!$this->dfs($i)) {
                return false;
            }
        }

        return true;
    }

    private function dfs(int $i): bool
    {
        if ($this->visited[$i] ?? false) {
            return false;
        }

        if (empty($this->graph[$i])) {
            return true;
        }

        $this->visited[$i] = true;

        foreach ($this->graph[$i] as $neighbour) {
            if (!$this->dfs($neighbour)) {
                return false;
            }
        }

        $this->visited[$i] = false;
        $this->graph[$i] = [];

        return true;
    }
}