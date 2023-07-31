package graph

import "container/heap"

// https://www.hackerrank.com/challenges/primsmstsub/problem?isFullScreen=true

type PqItem struct {
	value    *SpecialTreeNode
	priority int
	index    int
}

type PriorityQueue []*PqItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) IsEmpty() bool {
	return len(pq) == 0
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type SpecialTreeEdge struct {
	target *SpecialTreeNode
	weight int32
}

type SpecialTreeNode struct {
	edges []*SpecialTreeEdge
}

type SpecialTreeGraph struct {
	nodes []*SpecialTreeNode
}

func createSpecialTreeGraph(n int32, edges [][]int32) *SpecialTreeGraph {
	graph := &SpecialTreeGraph{}

	for i := 0; i < int(n); i++ {
		node := &SpecialTreeNode{}
		graph.nodes = append(graph.nodes, node)
	}

	for i := 0; i < len(edges); i++ {
		sourceIndex := edges[i][0] - 1
		targetIndex := edges[i][1] - 1
		weight := edges[i][2]

		source := graph.nodes[sourceIndex]
		target := graph.nodes[targetIndex]

		source.edges = append(source.edges, &SpecialTreeEdge{target: target, weight: weight})
		target.edges = append(target.edges, &SpecialTreeEdge{target: source, weight: weight})
	}

	return graph
}

func prims(n int32, edges [][]int32, start int32) int32 {
	graph := createSpecialTreeGraph(n, edges)
	visited := make(map[*SpecialTreeNode]bool, n)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startingNode := graph.nodes[start-1]
	heap.Push(&pq, &PqItem{value: startingNode, priority: 0})

	var result int32 = 0

	for !pq.IsEmpty() {
		item := heap.Pop(&pq).(*PqItem)
		node := item.value

		if visited[node] {
			continue
		}

		result += int32(item.priority)
		visited[node] = true

		for _, edge := range node.edges {
			if !visited[edge.target] {
				heap.Push(&pq, &PqItem{
					value:    edge.target,
					priority: int(edge.weight),
				})
			}
		}
	}

	return result
}
