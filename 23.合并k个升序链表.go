package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	heap := make([]*ListNode, len(lists))
	heapSize := 0
	if len(lists) == 0 {
		return nil
	}
	for _, list := range lists {
		node := list
		if node == nil {
			continue
		}
		heap[heapSize] = node
		heapSize++
		// fmt.Print((heapSize-2)/2, heapSize, "\t")
		// printHeep(heap, heapSize)
		heapifyFromBottom(heap, (heapSize-2)/2, heapSize)
	}
	// printHeep(heap, heapSize)
	root := new(ListNode)
	if root == nil {
		return nil
	}
	current := root
	for {
		if heap[0] == nil {
			heapSize--
			if heapSize <= 0 {
				break
			}
			heap[0] = heap[heapSize]
		}
		heapifyFromTop(heap, 0, heapSize)
		// printHeep(heap, heapSize)
		current.Next = heap[0]
		current = current.Next
		heap[0] = heap[0].Next
		current.Next = nil

	}
	return root.Next
}

// i is not leaf node index
func heapifyFromBottom(h []*ListNode, i, size int) {
	for {
		minNode := i
		if i*2+2 < size && h[minNode].Val > h[i*2+2].Val {
			minNode = i*2 + 2
		}
		if i*2+1 < size && h[minNode].Val > h[i*2+1].Val {
			minNode = i*2 + 1
		}
		if minNode == i {
			break
		}
		swap(h, i, minNode)
		i = (i - 1) / 2
	}
}

func heapifyFromTop(h []*ListNode, i, size int) {
	minNode := i
	for {
		if i*2+1 < size && h[minNode].Val > h[i*2+1].Val {
			minNode = i*2 + 1
		}
		if i*2+2 < size && h[minNode].Val > h[i*2+2].Val {
			minNode = i*2 + 2
		}
		if minNode == i {
			break
		}
		swap(h, i, minNode)
		i = minNode
	}
}

func swap(h []*ListNode, a, b int) {
	h[a], h[b] = h[b], h[a]
}

// @lc code=end
func printHeep(h []*ListNode, size int) {
	fmt.Print("heep: \t")
	for i := 0; i < size; i++ {
		fmt.Printf(" %d", h[i].Val)
	}
	fmt.Print("\n")
}
func printvalue(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf(" %d", p.Val)
		p = p.Next
	}
	fmt.Print("\n")
}
func l(item ...*ListNode) []*ListNode {
	return item
}
func n(values ...int) *ListNode {
	l := new(ListNode)
	root := l
	for _, v := range values {
		l.Next = new(ListNode)
		l = l.Next
		l.Val = v
	}
	return root.Next
}
func main() {
	// printvalue(mergeKLists(l()))
	// printvalue(mergeKLists(l(n())))
	// printvalue(mergeKLists(l(n(-6, -3, -1, 1, 2, 2, 2), n(-10, -8, -6, -2, 4), n(-2), n(-8, -4, -3, -3, -2, -1, 1, 2, 3), n(-8, -6, -5, -4, -2, -2, 2, 4))))
	printvalue(mergeKLists(l(n(-10, -9, -9, -3, -1, -1, 0), n(-5), n(4), n(-8), n(), n(-9, -6, -5, -4, -2, 2, 3), n(-3, -3, -2, -1, 0))))
	// printvalue(mergeKLists(l(n(), n())))
	// printvalue(mergeKLists(l(n(-8, -7, -6, -3, -2, -2, 0, 3), n(-10, -6, -4, -4, -4, -2, -1, 4), n(-10, -9, -8, -8, -6), n(-10, 0, 4))))
	// printvalue(mergeKLists(l(n(1), n(0))))
	// printvalue(mergeKLists(l(n(-1, 1), n(-3, 1, 4), n(-2, -1, 0, 2))))
	// printvalue(mergeKLists(l(n(-2, -1, -1, -1), n())))
	// printvalue(mergeKLists(l(n(1, 3, 4, 6, 8, 9, 12), n(1, 2, 5, 7, 11, 21, 24), n(-4, 0, 4, 7, 10, 14, 22, 29))))
}
