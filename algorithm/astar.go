// Package algorithm 提供算法实现
// 该包包含A*寻路等算法
package algorithm

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

// AStar A*寻路算法实现
// 用于在地图上寻找最短路径
type AStar struct {
	mapWidth  int           // 地图宽度
	mapHeight int           // 地图高度
	walls     map[int]bool  // 障碍物集合
}

// Point 表示二维空间中的一个点
type Point struct {
	X int
	Y int
}

// String 将点转换为字符串
// 返回值:
//   string: 点的字符串表示
func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

// Equals 比较两个点是否相等
// 参数:
//   other: 另一个点
// 返回值:
//   bool: 如果相等返回true，否则返回false
func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// Node A*算法中的节点
// 包含位置信息和路径成本
type Node struct {
	Point
	GCost   int     // 从起点到当前节点的实际成本
	HCost   int     // 从当前节点到目标的估计成本（启发式）
	FCost   int     // 总成本 (G + H)
	Parent  *Node   // 父节点
}

// F 计算F成本
// 返回值:
//   int: F成本 (G + H)
func (n *Node) F() int {
	return n.GCost + n.HCost
}

// PriorityQueue 优先队列（最小堆）
// 用于A*算法中按F成本排序
type PriorityQueue []*Node

// Len 返回队列长度
// 返回值:
//   int: 队列长度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 比较两个节点的F成本
// 参数:
//   i, j: 节点索引
// 返回值:
//   bool: 如果i的F成本小于j返回true
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F() < pq[j].F()
}

// Swap 交换两个节点
// 参数:
//   i, j: 节点索引
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 添加节点到队列
// 参数:
//   x: 要添加的节点
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

// Pop 从队列中弹出节点
// 返回值:
//   interface{}: 弹出的节点
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// NewAStar 创建A*算法实例
// 返回值:
//   *AStar: A*算法实例
func NewAStar() *AStar {
	return &AStar{
		walls: make(map[int]bool),
	}
}

// SetMap 设置地图
// 参数:
//   width: 地图宽度
//   height: 地图高度
//   walls: 障碍物坐标列表
func (a *AStar) SetMap(width, height int, walls []Point) {
	a.mapWidth = width
	a.mapHeight = height
	a.walls = make(map[int]bool)
	for _, wall := range walls {
		a.walls[a.pointToKey(wall)] = true
	}
}

// FindPath 寻找路径
// 参数:
//   startX, startY: 起点坐标
//   endX, endY: 终点坐标
// 返回值:
//   []Point: 路径点列表
//   error: 错误信息
func (a *AStar) FindPath(startX, startY, endX, endY int) ([]Point, error) {
	start := Point{X: startX, Y: startY}
	end := Point{X: endX, Y: endY}

	// 检查起点和终点是否有效
	if !a.isValidPoint(start) || !a.isValidPoint(end) {
		return nil, fmt.Errorf("invalid start or end point")
	}

	// 检查起点和终点是否是障碍物
	if a.isWall(start) || a.isWall(end) {
		return nil, fmt.Errorf("start or end point is a wall")
	}

	// 初始化开放列表和关闭列表
	openList := &PriorityQueue{}
	heap.Init(openList)
	closedList := make(map[int]bool)

	// 创建起始节点
	startNode := &Node{
		Point: start,
		GCost: 0,
		HCost: heuristic(start, end),
	}
	heap.Push(openList, startNode)

	// 主循环
	for openList.Len() > 0 {
		// 获取F成本最低的节点
		current := heap.Pop(openList).(*Node)

		// 检查是否到达终点
		if current.Point.Equals(end) {
			return a.reconstructPath(current), nil
		}

		// 标记为已访问
		closedList[a.pointToKey(current.Point)] = true

		// 检查邻居节点
		neighbors := a.getNeighbors(current.Point)
		for _, neighbor := range neighbors {
			// 跳过已访问的节点
			if closedList[a.pointToKey(neighbor)] {
				continue
			}

			// 跳过障碍物
			if a.isWall(neighbor) {
				continue
			}

			// 计算新的G成本
			newGCost := current.GCost + 1

			// 创建新节点
			newNode := &Node{
				Point:   neighbor,
				GCost:   newGCost,
				HCost:   heuristic(neighbor, end),
				Parent:  current,
			}

			// 添加到开放列表
			heap.Push(openList, newNode)
		}
	}

	// 未找到路径
	return nil, fmt.Errorf("no path found")
}

// isValidPoint 检查点是否在地图范围内
// 参数:
//   p: 要检查的点
// 返回值:
//   bool: 如果有效返回true，否则返回false
func (a *AStar) isValidPoint(p Point) bool {
	return p.X >= 0 && p.X < a.mapWidth && p.Y >= 0 && p.Y < a.mapHeight
}

// isWall 检查点是否是障碍物
// 参数:
//   p: 要检查的点
// 返回值:
//   bool: 如果是障碍物返回true，否则返回false
func (a *AStar) isWall(p Point) bool {
	return a.walls[a.pointToKey(p)]
}

// getNeighbors 获取邻居节点
// 参数:
//   p: 当前点
// 返回值:
//   []Point: 邻居点列表
func (a *AStar) getNeighbors(p Point) []Point {
	neighbors := []Point{
		{X: p.X + 1, Y: p.Y},     // 右
		{X: p.X - 1, Y: p.Y},     // 左
		{X: p.X, Y: p.Y + 1},     // 下
		{X: p.X, Y: p.Y - 1},     // 上
		{X: p.X + 1, Y: p.Y + 1}, // 右下
		{X: p.X - 1, Y: p.Y - 1}, // 左上
		{X: p.X + 1, Y: p.Y - 1}, // 右上
		{X: p.X - 1, Y: p.Y + 1}, // 左下
	}

	// 过滤无效点
	validNeighbors := []Point{}
	for _, neighbor := range neighbors {
		if a.isValidPoint(neighbor) {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}
	return validNeighbors
}

// reconstructPath 重建路径
// 参数:
//   node: 终点节点
// 返回值:
//   []Point: 路径点列表
func (a *AStar) reconstructPath(node *Node) []Point {
	path := []Point{}
	current := node
	for current != nil {
		path = append([]Point{current.Point}, path...)
		current = current.Parent
	}
	return path
}

// pointToKey 将点转换为唯一键
// 参数:
//   p: 点
// 返回值:
//   int: 唯一键
func (a *AStar) pointToKey(p Point) int {
	return p.Y*a.mapWidth + p.X
}

// heuristic 启发式函数（曼哈顿距离）
// 参数:
//   a, b: 两个点
// 返回值:
//   int: 曼哈顿距离
func heuristic(a, b Point) int {
	dx := int(math.Abs(float64(a.X - b.X)))
	dy := int(math.Abs(float64(a.Y - b.Y)))
	return dx + dy
}

// FindNearestPos 查找最近的位置
// 参数:
//   allPos: 所有位置字符串，格式: "x1,y1|x2,y2|..." 或 "name1,x1,y1|name2,x2,y2|..."
//   type_: 类型 (1=纯坐标, 2=带名称的坐标)
//   x, y: 目标坐标
// 返回值:
//   string: 最近的位置
func FindNearestPos(allPos string, type_ int, x, y int) string {
	parts := strings.Split(allPos, "|")
	var nearestPos string
	minDist := math.MaxFloat64

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		var px, py int
		var name string

		if type_ == 1 {
			// 纯坐标格式: "x,y"
			_, err := fmt.Sscanf(part, "%d,%d", &px, &py)
			if err != nil {
				continue
			}
		} else {
			// 带名称的坐标格式: "name,x,y"
			_, err := fmt.Sscanf(part, "%s,%d,%d", &name, &px, &py)
			if err != nil {
				continue
			}
		}

		// 计算距离
		dist := math.Sqrt(float64((x-px)*(x-px) + (y-py)*(y-py)))
		if dist < minDist {
			minDist = dist
			if type_ == 1 {
				nearestPos = fmt.Sprintf("%d,%d", px, py)
			} else {
				nearestPos = fmt.Sprintf("%s,%d,%d", name, px, py)
			}
		}
	}

	return nearestPos
}

// ParsePositions 解析位置字符串
// 参数:
//   s: 位置字符串，格式: "x1,y1|x2,y2|..."
// 返回值:
//   []Point: 位置列表
//   error: 错误信息
func ParsePositions(s string) ([]Point, error) {
	parts := strings.Split(s, "|")
	points := make([]Point, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		var x, y int
		_, err := fmt.Sscanf(part, "%d,%d", &x, &y)
		if err != nil {
			return nil, err
		}
		points = append(points, Point{X: x, Y: y})
	}

	return points, nil
}
