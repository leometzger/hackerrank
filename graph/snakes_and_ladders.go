package graph

import (
	"math"
)

// https://www.hackerrank.com/challenges/the-quickest-way-up/problem?isFullScreen=false

type SnakesAndLaddersEdge struct {
	node          *SnakesAndLaddersNode
	ladderOrSnake bool
}

type SnakesAndLaddersNode struct {
	key   int32
	edges []*SnakesAndLaddersEdge
}

type SnakesAndLaddersGraph struct {
	nodes []*SnakesAndLaddersNode
}

func createSnakesAndLaddersGraph(ladders [][]int32, snakes [][]int32) *SnakesAndLaddersGraph {
	graph := &SnakesAndLaddersGraph{
		nodes: make([]*SnakesAndLaddersNode, 100),
	}

	for i := 1; i <= 100; i++ {
		graph.nodes[i-1] = &SnakesAndLaddersNode{
			key: int32(i),
		}
	}

	for i := 0; i < len(snakes); i++ {
		sourceIndex := snakes[i][0] - 1
		targetIndex := snakes[i][1] - 1
		graph.nodes[sourceIndex].edges = append(graph.nodes[sourceIndex].edges, &SnakesAndLaddersEdge{
			node:          graph.nodes[targetIndex],
			ladderOrSnake: true,
		})
	}

	for i := 0; i < len(ladders); i++ {
		sourceIndex := ladders[i][0] - 1
		targetIndex := ladders[i][1] - 1
		graph.nodes[sourceIndex].edges = append(graph.nodes[sourceIndex].edges, &SnakesAndLaddersEdge{
			node:          graph.nodes[targetIndex],
			ladderOrSnake: true,
		})
	}

	for i := 0; i < 100; i++ {
		hasSnakeOrLadder := len(graph.nodes[i].edges) > 0
		if hasSnakeOrLadder {
			continue
		}

		for j := 1; j <= 6; j++ {
			if i+j >= 100 {
				continue
			}

			graph.nodes[i].edges = append(graph.nodes[i].edges, &SnakesAndLaddersEdge{
				node:          graph.nodes[i+j],
				ladderOrSnake: false,
			})
		}
	}

	return graph
}

type SnakesAndLaddersQueueEntry struct {
	node *SnakesAndLaddersNode
	dist int
}

type SnakesAndLaddersQueue struct {
	items []*SnakesAndLaddersQueueEntry
}

func (q *SnakesAndLaddersQueue) enqueue(node *SnakesAndLaddersNode, dist int) {
	q.items = append(q.items, &SnakesAndLaddersQueueEntry{
		node: node,
		dist: dist,
	})
}

func (q *SnakesAndLaddersQueue) dequeue() *SnakesAndLaddersQueueEntry {
	n := len(q.items)
	item := q.items[0]
	q.items = q.items[1:n]

	return item
}

func (q *SnakesAndLaddersQueue) isEmpty() bool {
	return len(q.items) == 0
}

func quickestWayUp(ladders [][]int32, snakes [][]int32) int32 {
	graph := createSnakesAndLaddersGraph(ladders, snakes)

	queue := &SnakesAndLaddersQueue{}
	queue.enqueue(graph.nodes[0], 0)

	visited := make(map[*SnakesAndLaddersNode]bool, 100)
	var min int32 = math.MaxInt32

	for !queue.isEmpty() {
		entry := queue.dequeue()
		if visited[entry.node] {
			continue
		}

		visited[entry.node] = true
		for _, edge := range entry.node.edges {
			increment := 1
			if edge.ladderOrSnake {
				increment = 0
			}

			if edge.node.key == 100 {
				if int32(entry.dist) < min {
					min = int32(entry.dist + increment)
				}
			}

			if !visited[edge.node] {
				queue.enqueue(edge.node, entry.dist+increment)
			}
		}
	}

	if min == math.MaxInt32 {
		return -1
	}

	return min
}
