package datastructures

type MergingCommunitiesSet struct {
	items   map[int32]int32
	rank    map[int32]int32
	counter map[int32]int32
}

func (s *MergingCommunitiesSet) add(item int32) {
	s.items[item] = item
	s.rank[item] = 1
	s.counter[item] = 1
}

func (s *MergingCommunitiesSet) find(item int32) int32 {
	if s.items[item] == item {
		return item
	}
	return s.find(s.items[item])
}

func (s *MergingCommunitiesSet) union(i, j int32) {
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
		s.items[jrep] = irep
		s.rank[irep]++
	}

	counter := s.counter[irep] + s.counter[jrep]
	s.counter[jrep] = counter
	s.counter[irep] = counter
}

func (s *MergingCommunitiesSet) count(item int32) int32 {
	irep := s.find(item)
	return s.counter[irep]
}

func MergingCommunities(n int32, commands []string, params [][]int32) []int32 {
	set := &MergingCommunitiesSet{
		items:   make(map[int32]int32, n),
		rank:    make(map[int32]int32, n),
		counter: make(map[int32]int32, n),
	}
	result := []int32{}

	for i := 1; i <= int(n); i++ {
		set.add(int32(i))
	}

	for index, command := range commands {
		param := params[index]

		if command == "M" {
			set.union(param[0], param[1])
		}

		if command == "Q" {
			result = append(result, set.count(param[0]))
		}
	}

	return result
}
