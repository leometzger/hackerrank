package datastructures

type CastleOnGridQueue struct {
	items []*CastleOnGridNode
}

func (q CastleOnGridQueue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *CastleOnGridQueue) enqueue(node *CastleOnGridNode) {
	q.items = append(q.items, node)
}

func (q *CastleOnGridQueue) dequeue() *CastleOnGridNode {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

type CastleOnGridNode struct {
	x     int32
	y     int32
	edges []*CastleOnGridNode
}

type CastleOnGridGraph struct {
	nodes []*CastleOnGridNode
}

func (g *CastleOnGridGraph) addNode(x, y int32) *CastleOnGridNode {
	node := &CastleOnGridNode{
		x: x,
		y: y,
	}
	g.nodes = append(g.nodes, node)
	return node
}

func (g *CastleOnGridGraph) addEdge(source *CastleOnGridNode, target *CastleOnGridNode) {
	source.edges = append(source.edges, target)
	target.edges = append(target.edges, source)
}

func createGraph(grid []string) [][]*CastleOnGridNode {
	N := len(grid)

	graph := &CastleOnGridGraph{
		nodes: make([]*CastleOnGridNode, N),
	}
	nodesMatrix := make([][]*CastleOnGridNode, N)

	for i := 0; i < N; i++ {
		nodesMatrix[i] = make([]*CastleOnGridNode, N)

		for j := 0; j < N; j++ {
			node := graph.addNode(int32(i), int32(j))
			nodesMatrix[i][j] = node
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			item := grid[i][j]

			if item == 'X' {
				break
			}

			if item == '.' {
				// to right
				for k := j + 1; k < N; k++ {
					if grid[i][k] == 'X' {
						break
					}
					graph.addEdge(nodesMatrix[i][j], nodesMatrix[i][k])
				}

				// to bottom
				for k := i + 1; k < N; k++ {
					if grid[k][j] == 'X' {
						break
					}
					graph.addEdge(nodesMatrix[i][j], nodesMatrix[k][j])
				}
			}
		}
	}
	return nodesMatrix
}

func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
	queue := &CastleOnGridQueue{}
	nodesMatrix := createGraph(grid)
	visited := make(map[*CastleOnGridNode]bool)

	start := nodesMatrix[startX][startY]
	queue.enqueue(start)

	minMoves := make([][]int32, len(grid))
	for i := 0; i < len(grid); i++ {
		minMoves[i] = make([]int32, len(grid))

		for j := 0; j < len(grid); j++ {
			minMoves[i][j] = int32(len(grid) + 1)
		}
	}

	for !queue.isEmpty() {
		node := queue.dequeue()
		if visited[node] {
			continue
		}

		for _, target := range node.edges {
			queue.enqueue(target)

			min := minMoves[target.x][target.y] + 1
			if min < minMoves[target.x][target.y] {
				minMoves[target.x][target.y] = min
			}
		}

		visited[node] = true
	}

	return minMoves[goalX][goalY]
}
