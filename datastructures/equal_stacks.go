package datastructures

// https://www.hackerrank.com/challenges/equal-stacks/problem?isFullScreen=false

type Stack struct {
	items  []int32
	height int32
}

func NewStack(items []int32) *Stack {
	height := int32(0)
	for _, val := range items {
		height += val
	}

	return &Stack{
		items:  items,
		height: height,
	}
}

func (s *Stack) Pop() {

}

// problema: O(N²)
func Verify(stacks []*Stack) bool {
	for i := 0; i < len(stacks); i++ {
		for j := 0; j < len(stacks); j++ {
			if stacks[i].height != stacks[j].height {
				return false
			}
		}
	}
	return true
}

func PopFromHighest(stacks []*Stack) {

}

func EqualStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	var stacks []*Stack

	stack1 := NewStack(h1)
	stack2 := NewStack(h2)
	stack3 := NewStack(h3)

	stacks = append(stacks, stack1, stack2, stack3)

	return 0
}
