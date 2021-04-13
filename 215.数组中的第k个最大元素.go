package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	heapLen := len(nums)
	buildHeap(nums, heapLen)
	for i := 1; i < k; i++ {
		swap(nums, 0, heapLen-i)
		heapify(nums, heapLen-i, 0)
	}

	return nums[0]
}

func buildHeap(n []int, size int) {
	for i := (size - 1) / 2; i >= 0; i-- {
		heapify(n, size, i)
	}
}

func heapify(n []int, size, i int) {
	maxPos := i
	for {
		i = maxPos
		if i*2+1 < size && n[maxPos] < n[i*2+1] {
			maxPos = i*2 + 1
		}
		if i*2+2 < size && n[maxPos] < n[i*2+2] {
			maxPos = i*2 + 2
		}
		if i == maxPos {
			break
		}
		swap(n, i, maxPos)
	}
}

func swap(n []int, a, b int) {
	n[a], n[b] = n[b], n[a]
}

// @lc code=end
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))

}
