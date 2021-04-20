package main

import (
	"fmt"
	"log"
	"time"
)

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	steps := make([]int, m)
	for i := 0; i < m; i++ {
		steps[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			steps[j] = steps[j] + steps[j-1]
		}
	}
	return steps[m-1]
}

// @lc code=end
func test(m int, n int, except int) {
	s := time.Now()
	ans := uniquePaths(m, n)
	e := time.Now()
	if ans != except {
		log.Panicf("error ansower!! input: m: %d, n: %d, except ans: %d, ans: %d", m, n, except, ans)
	} else {
		fmt.Printf("input: m: %d, n: %d, except ans: %d, ans: %d time:", m, n, except, ans)
		fmt.Println(e.Sub(s))
	}
}
func main() {
	test(3, 7, 28)
	test(3, 2, 3)
	test(7, 3, 28)
	test(3, 3, 6)
}
