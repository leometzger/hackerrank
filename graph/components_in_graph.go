package graph

import "math"

type ComponentsInGraphSets struct {
	items map[int32]int32
	rank  map[int32]int32
}

func (s *ComponentsInGraphSets) add(key int32) {
	s.items[key] = key
}

func (s *ComponentsInGraphSets) find(i int32) int32 {
	if s.items[i] == i {
		return i
	}
	return s.find(s.items[i])
}

func (s *ComponentsInGraphSets) union(i int32, j int32) {
	irep := s.find(i)
	jrep := s.find(j)

	if irep == jrep {
		return
	}

	irank := s.rank[irep]
	jrank := s.rank[jrep]

	if irank < jrank {
		s.items[irep] = jrep
	} else if irank > jrank {
		s.items[jrep] = irep
	} else {
		s.items[irep] = jrep
		s.rank[jrep]++
	}
}

func componentsInGraph(gb [][]int32) []int32 {
	set := &ComponentsInGraphSets{
		items: make(map[int32]int32, len(gb)),
		rank:  make(map[int32]int32, len(gb)),
	}
	counter := make(map[int32]int, len(gb))

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
		set.union(i, j)
	}

	for key, _ := range set.items {
		root := set.find(key)
		counter[root]++
	}

	for _, count := range counter {
		if count < min {
			min = count
		}

		if count > max {
			max = count
		}
	}

	return []int32{int32(min), int32(max)}
}
