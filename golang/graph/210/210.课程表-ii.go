package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	// start := time.Now()
	inDegree := make([]int, numCourses)
	queue := list.New()
	neighborTable := make([][]int, numCourses)
	for i := range neighborTable {
		neighborTable[i] = make([]int, 0)
	}
	ans := make([]int, 0)
	for _, prerequisite := range prerequisites {
		inDegree[prerequisite[0]] += 1
		neighborTable[prerequisite[1]] = append(neighborTable[prerequisite[1]], prerequisite[0])
	}
	// fmt.Println(time.Now().Sub(start))
	for i, degree := range inDegree {
		if degree == 0 {
			queue.PushBack(i)
		}
	}
	// fmt.Println(time.Now().Sub(start))
	for queue.Len() != 0 {
		e := queue.Front()
		queue.Remove(e)
		i := e.Value.(int)
		ans = append(ans, i)
		for _, neighbor := range neighborTable[i] {
			inDegree[neighbor] -= 1
			if inDegree[neighbor] == 0 {
				queue.PushBack(neighbor)
			}
		}
	}
	// fmt.Println(time.Now().Sub(start))
	if len(ans) == numCourses {
		return ans
	} else {
		return []int{}
	}

}

// @lc code=end

func main() {
	fmt.Print(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
