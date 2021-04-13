/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func maxArea(height []int) int {
	x := 0
	y := len(height) - 1
	maxArea := 0
	for y > x {
		area := min(height[x], height[y]) * (y - x)
		if maxArea < area {
			maxArea = area
		}
		if height[x] >= height[y] {
			y--
		} else {
			x++
		}
	}
	return maxArea
}

// @lc code=end

