package main

import "fmt"

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var n uint32 = 0
	n = n | (num & 0x1)
	for i := 1; i < 32; i++ {
		n = n << 1
		num = num >> 1
		n = n | (num & 0x1)
	}
	return n
}

// @lc code=end
func main() {
	fmt.Print(reverseBits(692074872))
}
