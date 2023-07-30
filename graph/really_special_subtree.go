package graph

import (
	"sort"
)

// https://www.hackerrank.com/challenges/kruskalmstrsub/problem?isFullScreen=true

type byWeight []ReallySpecialSubtreeEdge

func (e byWeight) Len() int           { return len(e) }
func (e byWeight) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e byWeight) Less(i, j int) bool { return e[i].weight < e[j].weight }

type ReallySpecialSubtreeEdge struct {
	source int32
	target int32
	weight int32
}

type ReallySpeciallSubtreeGraph struct {
	nodes []int32
	edges []ReallySpecialSubtreeEdge
}

// Disjoint Set - auxiliary datastructure
type ReallySpeciallSubtreeSet struct {
	items map[int32]int32
	rank  map[int32]int32
}

func (s *ReallySpeciallSubtreeSet) add(item int32) {
	s.items[item] = item
}

func (s *ReallySpeciallSubtreeSet) find(item int32) int32 {
	if s.items[item] == item {
		return item
	}

	return s.find(s.items[item])
}

func (s *ReallySpeciallSubtreeSet) union(i int32, j int32) {
	irank := s.rank[i]
	jrank := s.rank[j]
	irep := s.find(i)
	jrep := s.find(j)

	if jrep == irep {
		return
	}

	if irank < jrank {
		s.items[irep] = jrep
	} else if irank > jrank {
		s.items[jrep] = irep
	} else {
		s.items[irep] = jrep
		s.rank[jrep]++
	}
}

// Return the minimum spanning tree for this graph
func (graph *ReallySpeciallSubtreeGraph) minimumSpanningTree() int32 {
	var min int32 = 0
	set := &ReallySpeciallSubtreeSet{
		items: make(map[int32]int32),
		rank:  make(map[int32]int32),
	}

	for i := 0; i < len(graph.nodes); i++ {
		set.add(graph.nodes[i])
	}

	for _, edge := range graph.edges {
		if set.find(edge.source) != set.find(edge.target) {
			set.union(edge.source, edge.target)
			min += edge.weight
		}
	}

	return min
}

func createGraph(gNodes int32, gFrom []int32, gTo []int32, gWeight []int32) *ReallySpeciallSubtreeGraph {
	graph := &ReallySpeciallSubtreeGraph{}

	for i := 1; i <= int(gNodes); i++ {
		graph.nodes = append(graph.nodes, int32(i))
	}

	for i := 0; i < len(gFrom); i++ {
		graph.edges = append(graph.edges, ReallySpecialSubtreeEdge{
			source: gFrom[i],
			target: gTo[i],
			weight: gWeight[i],
		})
	}
	sort.Sort(byWeight(graph.edges))
	return graph
}

func kruskals(gNodes int32, gFrom []int32, gTo []int32, gWeight []int32) int32 {
	graph := createGraph(gNodes, gFrom, gTo, gWeight)
	return graph.minimumSpanningTree()
}
