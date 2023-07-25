package graph

type RoadsAndLibrariesGraph struct {
	cities  [][]int32
	visited map[int32]bool
}

func NewRoadsAndLibrariesGraph(n int32, roads [][]int32) *RoadsAndLibrariesGraph {
	cities := [][]int32{}

	for i := 0; int32(i) < n; i++ {
		city := []int32{}
		cities = append(cities, city)
	}

	for _, road := range roads {
		sourceIndex := road[0] - 1
		targetIndex := road[1] - 1

		cities[sourceIndex] = append(cities[sourceIndex], targetIndex)
		cities[targetIndex] = append(cities[targetIndex], sourceIndex)
	}

	return &RoadsAndLibrariesGraph{
		cities:  cities,
		visited: make(map[int32]bool),
	}
}

func (g *RoadsAndLibrariesGraph) dfs(v int32) {
	g.visited[v] = true

	for _, index := range g.cities[v] {
		if !g.visited[index] {
			g.dfs(index)
		}
	}
}

func (g *RoadsAndLibrariesGraph) countSubgraphs(n int32) int32 {
	counter := 1
	g.dfs(0)

	for i := 1; int32(i) < n; i++ {
		if !g.visited[int32(i)] {
			g.dfs(int32(i))
			counter += 1
		}
	}

	return int32(counter)
}

// https://www.hackerrank.com/challenges/torque-and-development/problem?isFullScreen=true
func RoadsAndLibraries(numCities int32, libCost int32, roadCost int32, roads [][]int32) int64 {
	if roadCost >= libCost {
		return int64(numCities) * int64(libCost)
	}

	graph := NewRoadsAndLibrariesGraph(numCities, roads)
	subgraphCount := graph.countSubgraphs(numCities)

	return int64(subgraphCount)*int64(libCost) + (int64(numCities)-int64(subgraphCount))*int64(roadCost)
}
