package main

import "fmt"

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var visited map[int]*Node = make(map[int]*Node)
	var needToVisit []*Node = make([]*Node, 0)
	ans := &Node{Val: node.Val, Neighbors: node.Neighbors}
	needToVisit = append(needToVisit, ans)
	visited[ans.Val] = ans
	for len(needToVisit) > 0 {
		visit := needToVisit[0]
		needToVisit = needToVisit[1:]
		if node.Neighbors == nil || len(node.Neighbors) == 0 {
			continue
		}
		neighbors := visit.Neighbors
		visit.Neighbors = make([]*Node, 0)
		for _, neighbor := range neighbors {
			if n, ok := visited[neighbor.Val]; ok {
				visit.Neighbors = append(visit.Neighbors, n)
			} else {
				n := &Node{Val: neighbor.Val, Neighbors: neighbor.Neighbors}
				needToVisit = append(needToVisit, n)
				visited[n.Val] = n
				visit.Neighbors = append(visit.Neighbors, n)
			}
		}
	}
	return ans
}

// @lc code=end

func buildGraph(neighborTable [][]int) *Node {
	nodes := make([]*Node, len(neighborTable))
	for i := 1; i <= len(neighborTable); i++ {
		node := &Node{Val: i}
		nodes[i-1] = node
	}
	for i, neighbors := range neighborTable {
		node := nodes[i]
		for _, neighbor := range neighbors {
			n := nodes[neighbor-1]
			node.Neighbors = append(node.Neighbors, n)
		}
	}
	return nodes[0]
}

func main() {
	ans := cloneGraph(buildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}))
	ans = cloneGraph(buildGraph([][]int{{2}, {1}}))
	ans = cloneGraph(buildGraph([][]int{}))
	ans = cloneGraph(buildGraph([][]int{{}}))
	fmt.Print(ans)
}
