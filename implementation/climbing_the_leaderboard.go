package implementation

func searchIndex(ranked []int32, player int32, current int) int32 {
	// is the first one
	if ranked[0] < player {
		return -1
	}

	// is the last one
	if ranked[len(ranked)-1] > player {
		return int32(len(ranked) - 1)
	}

	// binary search
	mid := len(ranked) / 2
	for {
		if ranked[mid] == player {
			return int32(mid)
		}

		if ranked[mid] > player && ranked[mid+1] <= player { // found the middle guy
			return int32(mid)
		}

		if ranked[mid] < player {
			mid = mid / 2
		} else {
			mid = mid + mid/2
		}
	}
}

func climbingLeaderboard(ranked []int32, players []int32) []int32 {
	var result []int32
	positions := make([]int32, len(ranked))

	counter := int32(1)
	positions[0] = counter

	for i := 1; i < len(ranked); i++ {
		if ranked[i] == ranked[i-1] {
			positions[i] = counter
		} else {
			counter++
			positions[i] = counter
		}
	}

	for i := 0; i < len(players); i++ {
		previous := searchIndex(ranked, players[i], len(ranked)/2)

		if previous == -1 {
			result = append(result, 1)
		} else if ranked[previous] == players[i] {
			result = append(result, positions[previous])
		} else {
			result = append(result, positions[previous]+1)
		}
	}

	return result
}
