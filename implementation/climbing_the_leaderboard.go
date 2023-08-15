package implementation

type Stack struct {
	items []int32
}

func (s *Stack) push(value int32) { s.items = append(s.items, value) }
func (s *Stack) isEmpty() bool    { return len(s.items) == 0 }
func (s *Stack) peek() int32      { return s.items[len(s.items)-1] }

func (s *Stack) pop() int32 {
	result := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return result
}

func climbingLeaderboard(ranked []int32, players []int32) []int32 {
	result := make([]int32, len(players))
	resultIndex := len(players) - 1
	currentPosition := int32(1)

	stack := &Stack{}
	for _, item := range players {
		stack.push(item)
	}

	for i := 0; i < len(ranked); i++ {
		if i > 0 && ranked[i-1] > ranked[i] {
			currentPosition++
		}

		for !stack.isEmpty() && stack.peek() >= ranked[i] {
			stack.pop()
			result[resultIndex] = currentPosition
			resultIndex--
		}

	}

	for i := resultIndex; i >= 0; i-- {
		result[i] = currentPosition + 1
	}

	return result
}
