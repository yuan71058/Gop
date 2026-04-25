// Package algorithm 提供算法实现
// 该包包含A*寻路算法等算法实现
package algorithm

import (
	"container/heap"
	"math"
)

// Point 二维坐标点
// 用于A*算法中的节点
type Point struct {
	X int
	Y int
}

// Equals 判断两个点是否相等
// 参数:
//   other: 另一个点
// 返回值:
//   bool: 是否相等
func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// Node A*算法节点
// 包含位置信息和路径代价
type Node struct {
	Point
	GCost   int     // 从起点到当前节点的实际代价
	HCost   int     // 从当前节点到终点的估计代价
	Parent  *Node   // 父节点
}

// FCost 获取总代价 (G + H)
// 返回值:
//   int: 总代价
func (n *Node) FCost() int {
	return n.GCost + n.HCost
}

// PriorityQueue 优先队列
// 用于A*算法中的开放列表
type PriorityQueue []*Node

// Len 获取队列长度
// 返回值:
//   int: 队列长度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 比较两个节点的优先级
// 参数:
//   i, j: 节点索引
// 返回值:
//   bool: i是否小于j
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].FCost() < pq[j].FCost()
}

// Swap 交换两个节点
// 参数:
//   i, j: 节点索引
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 向队列中添加节点
// 参数:
//   x: 节点指针
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

// Pop 从队列中移除并返回节点
// 返回值:
//   interface{}: 节点指针
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// AStar A*寻路算法
// 参数:
//   start: 起点坐标
//   end: 终点坐标
//   isWalkable: 判断某点是否可通行的函数
// 返回值:
//   []Point: 路径点列表，如果找不到路径返回nil
func AStar(start, end Point, isWalkable func(Point) bool) []Point {
	openList := &PriorityQueue{}
	heap.Init(openList)

	startNode := &Node{
		Point: start,
		GCost: 0,
		HCost: heuristic(start, end),
	}
	heap.Push(openList, startNode)

	closedSet := make(map[Point]bool)

	for openList.Len() > 0 {
		current := heap.Pop(openList).(*Node)

		if current.Point.Equals(end) {
			return reconstructPath(current)
		}

		closedSet[current.Point] = true

		for _, neighbor := range getNeighbors(current.Point) {
			if closedSet[neighbor] || !isWalkable(neighbor) {
				continue
			}

			gCost := current.GCost + 1
			hCost := heuristic(neighbor, end)

			newNode := &Node{
				Point:  neighbor,
				GCost:  gCost,
				HCost:  hCost,
				Parent: current,
			}

			heap.Push(openList, newNode)
		}
	}

	return nil
}

// heuristic 启发式函数 (曼哈顿距离)
// 参数:
//   a, b: 两个点
// 返回值:
//   int: 曼哈顿距离
func heuristic(a, b Point) int {
	dx := abs(a.X - b.X)
	dy := abs(a.Y - b.Y)
	return dx + dy
}

// abs 计算绝对值
// 参数:
//   x: 整数
// 返回值:
//   int: 绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// getNeighbors 获取相邻的可行走点
// 参数:
//   p: 当前点
// 返回值:
//   []Point: 相邻点列表
func getNeighbors(p Point) []Point {
	return []Point{
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}
}

// reconstructPath 重建路径
// 参数:
//   node: 终点节点
// 返回值:
//   []Point: 路径点列表
func reconstructPath(node *Node) []Point {
	var path []Point
	for node != nil {
		path = append([]Point{node.Point}, path...)
		node = node.Parent
	}
	return path
}

// Distance 计算两点之间的欧几里得距离
// 参数:
//   p1, p2: 两个点
// 返回值:
//   float64: 距离
func Distance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// ManhattanDistance 计算两点之间的曼哈顿距离
// 参数:
//   p1, p2: 两个点
// 返回值:
//   int: 曼哈顿距离
func ManhattanDistance(p1, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}
