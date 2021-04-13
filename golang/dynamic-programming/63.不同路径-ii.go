package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	steps := make([]int, n)
	block := false
	for i := 0; i < n; i++ {
		if !block && obstacleGrid[0][i] == 1 {
			block = true
		}
		if !block {
			steps[i] = 1
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			steps[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				steps[j] = steps[j] + steps[j-1]
			} else {
				steps[j] = 0
			}
		}
	}
	return steps[n-1]
}

// @lc code=end

func test(obstacleGrid [][]int, except int) {
	s := time.Now()
	ans := uniquePathsWithObstacles(obstacleGrid)
	e := time.Now()
	if ans != except {
		log.Panicf("error ansower!! input: obstacleGrid: %s, except ans: %d, ans: %d", printInput(obstacleGrid), except, ans)
	} else {
		fmt.Printf("input: obstacleGrid: %s, except ans: %d, ans: %d", printInput(obstacleGrid), except, ans)
		fmt.Println(e.Sub(s))
	}
}
func printInput(obstacleGrid [][]int) string {
	result, err := json.Marshal(obstacleGrid)
	if err != nil {
		return "[]"
	} else {
		return string(result)
	}

}

func main() {
	test([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2)
	test([][]int{{0, 1}, {0, 0}}, 1)
	test([][]int{{0, 1}}, 0)
}
