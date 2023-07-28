package datastructures

// https://www.hackerrank.com/challenges/equal-stacks/problem?isFullScreen=false

type EqualStacksStack struct {
	items  []int32
	height int32
}

func NewEqualStacksStack(items []int32) *EqualStacksStack {
	var stackItems []int32
	height := int32(0)

	for i := len(items) - 1; i >= 0; i-- {
		height += items[i]
		stackItems = append(stackItems, items[i])
	}

	return &EqualStacksStack{
		items:  stackItems,
		height: height,
	}
}

func (s *EqualStacksStack) Pop() int32 {
	length := len(s.items)
	tmp := s.items[length-1]
	s.items = s.items[:length-1]
	s.height -= tmp
	return tmp
}

// problem: O(N²) as it usually has low number of stacks it is kind of OK
func verify(stacks []*EqualStacksStack) bool {
	for i := 0; i < len(stacks); i++ {
		for j := 0; j < len(stacks); j++ {
			if stacks[i].height != stacks[j].height {
				return false
			}
		}
	}
	return true
}

func popFromHeighest(stacks []*EqualStacksStack) {
	heighest := stacks[0]

	for i := 1; i < len(stacks); i++ {
		if heighest.height < stacks[i].height {
			heighest = stacks[i]
		}
	}
	heighest.Pop()
}

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	var stacks []*EqualStacksStack

	stack1 := NewEqualStacksStack(h1)
	stack2 := NewEqualStacksStack(h2)
	stack3 := NewEqualStacksStack(h3)
	stacks = append(stacks, stack1, stack2, stack3)

	for {
		isEqual := verify(stacks)
		if !isEqual {
			popFromHeighest(stacks)
		} else {
			break
		}
	}

	return stack1.height
}
