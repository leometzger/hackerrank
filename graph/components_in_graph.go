package graph

import "math"

type ComponentsInGraphSets struct {
	items map[int32]int32
}

func (s *ComponentsInGraphSets) add(key int32) {
	s.items[key] = key
}

func (s *ComponentsInGraphSets) find(i int32) int32 {
	if s.items[i] == i {
		return i
	} else {
		return s.find(s.items[i])
	}
}

func (s *ComponentsInGraphSets) union(i int32, j int32) {
	irep := s.find(i)
	jrep := s.find(j)

	s.items[irep] = jrep
}

func componentsInGraph(gb [][]int32) []int32 {
	set := &ComponentsInGraphSets{items: make(map[int32]int32)}
	counter := make(map[int32]int)

	min := math.MaxInt32
	max := 0

	for _, edge := range gb {
		i := edge[0]
		j := edge[1]

		if _, ok := set.items[i]; !ok {
			set.add(i)
		}

		if _, ok := set.items[j]; !ok {
			set.add(j)
		}

	}

	for _, edge := range gb {
		i := edge[0]
		j := edge[1]
		set.union(i, j)
	}

	for key, _ := range set.items {
		root := set.find(key)
		counter[root] = counter[root] + 1
	}

	for _, count := range counter {
		if count < min && count > 1 {
			min = count
		}

		if count > max {
			max = count
		}
	}

	return []int32{int32(min), int32(max)}
}
