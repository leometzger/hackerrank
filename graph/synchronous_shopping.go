package graph

import (
	"container/heap"
	"log"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/synchronous-shopping/problem?isFullScreen=true

type SynchronousShoppingGraph struct {
	nodes []*SynchronousShoppingNode
}

type ShortestPathResultPair struct {
	weight    int32
	fishTypes []int32
}

func (g *SynchronousShoppingGraph) addEdge(source, target *SynchronousShoppingNode, weight int32) {
	source.addEdge(target, weight)
	target.addEdge(source, weight)
}

func (g *SynchronousShoppingGraph) addNode(node *SynchronousShoppingNode) {
	g.nodes = append(g.nodes, node)
}

func (g *SynchronousShoppingGraph) shortestsPath() {}

type SynchronousShoppingNode struct {
	key    int
	fishes []int32
	edges  []*SynchronousShoppingEdge
}

func (n *SynchronousShoppingNode) addEdge(node *SynchronousShoppingNode, weight int32) {
	n.edges = append(n.edges, &SynchronousShoppingEdge{
		weight: weight,
		node:   node,
	})
}

type SynchronousShoppingEdge struct {
	weight int32
	node   *SynchronousShoppingNode
}

type SynchronousShoppingPQ []*SyncrhonousShoppingPQEntry
type SyncrhonousShoppingPQEntry struct {
	priority int32
	node     *SynchronousShoppingNode
}

func (pq SynchronousShoppingPQ) Len() int           { return len(pq) }
func (pq SynchronousShoppingPQ) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq SynchronousShoppingPQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq SynchronousShoppingPQ) IsEmpty() bool      { return pq.Len() == 0 }

func (pq *SynchronousShoppingPQ) Push(x any) {
	item := x.(*SyncrhonousShoppingPQEntry)
	*pq = append(*pq, item)
}

func (pq *SynchronousShoppingPQ) Pop() any {
	n := len(*pq)
	old := *pq
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func createSynchronousGraph(n int32, k int32, centers []string, roads [][]int32) *SynchronousShoppingGraph {
	graph := &SynchronousShoppingGraph{}

	for i := 0; i < int(n); i++ {
		node := &SynchronousShoppingNode{key: i + 1}
		fishes := strings.Split(centers[i], " ")[1:]

		for j := 0; j < len(fishes); j++ {
			fish, err := strconv.Atoi(fishes[j])
			if err != nil {
				log.Fatalln(err)
			}

			node.fishes = append(node.fishes, int32(fish))
		}

		graph.addNode(node)
	}

	for i := 0; i < len(roads); i++ {
		weight := roads[i][2]
		source := graph.nodes[roads[i][0]-1]
		target := graph.nodes[roads[i][1]-1]
		graph.addEdge(source, target, weight)
	}

	return graph
}

func shop(n int32, k int32, centers []string, roads [][]int32) int32 {
	graph := createSynchronousGraph(n, k, centers, roads)

	pq := SynchronousShoppingPQ{}
	heap.Init(&pq)
	heap.Push(&pq, SyncrhonousShoppingPQEntry{
		priority: 0,
		node:     graph.nodes[0],
	})

	for !pq.IsEmpty() {

	}

	return 0
}
