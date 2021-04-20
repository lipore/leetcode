package main

import "fmt"

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	queue := make([]int, 0)
	neighborTable := make([][]int, numCourses)
	unLearnCount := numCourses
	for _, prerequisite := range prerequisites {
		inDegree[prerequisite[0]] += 1
		if neighborTable[prerequisite[1]] == nil {
			neighborTable[prerequisite[1]] = make([]int, 0)
		}
		neighborTable[prerequisite[1]] = append(neighborTable[prerequisite[1]], prerequisite[0])
	}
	for i, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, i)
			unLearnCount--
		}
	}
	for len(queue) != 0 {
		i := queue[0]
		queue = queue[1:]
		for _, neighbor := range neighborTable[i] {
			inDegree[neighbor] -= 1
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
				unLearnCount--
			}
		}
	}

	ans := unLearnCount == 0

	return ans
}

// @lc code=end
func main() {
	fmt.Print(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}
