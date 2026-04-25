// Package algorithm 提供算法实现
// 该包包含A*路径查找算法和其他工具算法
package algorithm

import (
	"container/heap"
	"fmt"
	"strings"
)

// Point 表示二维坐标
type Point struct {
	X int
	Y int
}

// AStar 实现A*路径查找算法
// 用于在网格地图中查找最短路径
type AStar struct {
	mapWidth  int              // 地图宽度
	mapHeight int              // 地图高度
	obstacles map[Point]bool   // 障碍点
}

// NewAStar 创建新的AStar实例
// 返回值:
//   *AStar: AStar实例
func NewAStar() *AStar {
	return &AStar{
		obstacles: make(map[Point]bool),
	}
}

// AStarFindPath 使用A*算法查找路径
// 参数:
//   mapWidth: 地图宽度
//   mapHeight: 地图高度
//   disablePoints: 障碍点字符串, 格式: "x1,y1|x2,y2|..."
//   beginX: 起点X坐标
//   beginY: 起点Y坐标
//   endX: 终点X坐标
//   endY: 终点Y坐标
// 返回值:
//   string: 路径结果, 格式: "x1,y1|x2,y2|..."
func (a *AStar) AStarFindPath(mapWidth, mapHeight int, disablePoints string, beginX, beginY, endX, endY int) string {
	a.mapWidth = mapWidth
	a.mapHeight = mapHeight

	// 解析障碍点
	a.parseObstacles(disablePoints)

	// 创建起点和终点
	start := Point{X: beginX, Y: beginY}
	end := Point{X: endX, Y: endY}

	// 验证点
	if !a.isValid(start) || !a.isValid(end) {
		return ""
	}

	// 检查起点或终点是否为障碍
	if a.obstacles[start] || a.obstacles[end] {
		return ""
	}

	// 运行A*算法
	path := a.findPath(start, end)
	if len(path) == 0 {
		return ""
	}

	// 转换路径为字符串
	return a.pathToString(path)
}

// FindNearestPos 从位置列表中查找到目标点最近的位置
// 参数:
//   allPos: 所有位置字符串, 格式: "x1,y1|x2,y2|..."
//   posType: 位置类型 (0=欧几里得距离, 1=曼哈顿距离)
//   x: 目标X坐标
//   y: 目标Y坐标
// 返回值:
//   string: 最近位置, 格式: "x,y"
func (a *AStar) FindNearestPos(allPos string, posType, x, y int) string {
	if allPos == "" {
		return ""
	}

	positions := strings.Split(allPos, "|")
	if len(positions) == 0 {
		return ""
	}

	var nearestPos string
	minDist := -1

	for _, pos := range positions {
		var px, py int
		if _, err := fmt.Sscanf(pos, "%d,%d", &px, &py); err != nil {
			continue
		}

		// 根据类型计算距离
		var dist int
		if posType == 0 {
			// 欧几里得距离(平方以避免sqrt)
			dx := px - x
			dy := py - y
			dist = dx*dx + dy*dy
		} else {
			// 曼哈顿距离
			dx := px - x
			dy := py - y
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			dist = dx + dy
		}

		if minDist == -1 || dist < minDist {
			minDist = dist
			nearestPos = pos
		}
	}

	return nearestPos
}

// parseObstacles 从字符串解析障碍点
// 参数:
//   disablePoints: 障碍点字符串, 格式: "x1,y1|x2,y2|..."
func (a *AStar) parseObstacles(disablePoints string) {
	a.obstacles = make(map[Point]bool)
	if disablePoints == "" {
		return
	}

	points := strings.Split(disablePoints, "|")
	for _, point := range points {
		var x, y int
		if _, err := fmt.Sscanf(point, "%d,%d", &x, &y); err == nil {
			a.obstacles[Point{X: x, Y: y}] = true
		}
	}
}

// isValid 检查点是否在地图范围内
// 参数:
//   p: 要检查的点
// 返回值:
//   bool: 有效返回true, 否则返回false
func (a *AStar) isValid(p Point) bool {
	return p.X >= 0 && p.X < a.mapWidth && p.Y >= 0 && p.Y < a.mapHeight
}

// findPath 使用A*算法查找从起点到终点的最短路径
// 参数:
//   start: 起点
//   end: 终点
// 返回值:
//   []Point: 路径点
func (a *AStar) findPath(start, end Point) []Point {
	// 创建开放和关闭集合
	openSet := &PriorityQueue{}
	heap.Init(openSet)

	// 跟踪已访问节点
	closedSet := make(map[Point]bool)

	// 跟踪g分数(从起点开始的成本)
	gScore := make(map[Point]int)
	gScore[start] = 0

	// 跟踪父节点用于路径重建
	parent := make(map[Point]Point)

	// 将起点添加到开放集合
	heap.Push(openSet, &AStarNode{
		Point:    start,
		GScore:   0,
		FScore:   a.heuristic(start, end),
	})

	// 主循环
	for openSet.Len() > 0 {
		// 获取f分数最低的节点
		current := heap.Pop(openSet).(*AStarNode)

		// 检查是否到达终点
		if current.Point == end {
			return a.reconstructPath(parent, current.Point)
		}

		// 如果已处理则跳过
		if closedSet[current.Point] {
			continue
		}
		closedSet[current.Point] = true

		// 检查邻居
		for _, neighbor := range a.getNeighbors(current.Point) {
			// 如果在关闭集合中或是障碍点则跳过
			if closedSet[neighbor] || a.obstacles[neighbor] {
				continue
			}

			// 计算临时g分数
			tentativeG := gScore[current.Point] + 1

			// 如果找到更好的路径则更新
			if _, exists := gScore[neighbor]; !exists || tentativeG < gScore[neighbor] {
				parent[neighbor] = current.Point
				gScore[neighbor] = tentativeG
				fScore := tentativeG + a.heuristic(neighbor, end)

				heap.Push(openSet, &AStarNode{
					Point:    neighbor,
					GScore:   tentativeG,
					FScore:   fScore,
				})
			}
		}
	}

	// 未找到路径
	return nil
}

// getNeighbors 获取点的有效邻居
// 参数:
//   p: 当前点
// 返回值:
//   []Point: 有效邻居点
func (a *AStar) getNeighbors(p Point) []Point {
	neighbors := []Point{
		{X: p.X + 1, Y: p.Y},     // 右
		{X: p.X - 1, Y: p.Y},     // 左
		{X: p.X, Y: p.Y + 1},     // 下
		{X: p.X, Y: p.Y - 1},     // 上
		{X: p.X + 1, Y: p.Y + 1}, // 右下
		{X: p.X - 1, Y: p.Y + 1}, // 左下
		{X: p.X + 1, Y: p.Y - 1}, // 右上
		{X: p.X - 1, Y: p.Y - 1}, // 左上
	}

	var valid []Point
	for _, n := range neighbors {
		if a.isValid(n) {
			valid = append(valid, n)
		}
	}
	return valid
}

// heuristic 计算两点之间的启发式距离
// 参数:
//   a: 第一个点
//   b: 第二个点
// 返回值:
//   int: 启发式距离
func (a *AStar) heuristic(p1, p2 Point) int {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	// 使用切比雪夫距离用于8方向移动
	if dx > dy {
		return dx
	}
	return dy
}

// reconstructPath 从父指针重建路径
// 参数:
//   parent: 父映射
//   end: 终点
// 返回值:
//   []Point: 路径点
func (a *AStar) reconstructPath(parent map[Point]Point, end Point) []Point {
	var path []Point
	current := end

	for current != (Point{}) {
		path = append([]Point{current}, path...)
		current = parent[current]
	}

	return path
}

// pathToString 将路径点转换为字符串格式
// 参数:
//   path: 路径点
// 返回值:
//   string: 路径字符串, 格式: "x1,y1|x2,y2|..."
func (a *AStar) pathToString(path []Point) string {
	if len(path) == 0 {
		return ""
	}

	var parts []string
	for _, p := range path {
		parts = append(parts, fmt.Sprintf("%d,%d", p.X, p.Y))
	}
	return strings.Join(parts, "|")
}

// AStarNode 表示A*算法中的节点
type AStarNode struct {
	Point    Point
	GScore   int
	FScore   int
}

// PriorityQueue 实现A*算法的优先队列
type PriorityQueue []*AStarNode

// Len 返回优先队列的长度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 通过f分数比较两个节点
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].FScore < pq[j].FScore
}

// Swap 交换两个节点
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 向优先队列添加节点
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*AStarNode)
	*pq = append(*pq, node)
}

// Pop 移除并返回最高优先级的节点
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}
