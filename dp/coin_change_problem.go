package dp

func getWays(n int64, coins []int64) int64 {
	ways := make([]int64, n+1)
	ways[0] = 1

	for _, coin := range coins {
		var i int64
		for i = coin; i <= n; i++ {
			ways[i] += ways[i-coin]
		}
	}

	return ways[n]
}
