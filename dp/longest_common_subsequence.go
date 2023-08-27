package dp

func lcs(x []int32, y []int32, i int32, j int32, memo [][]int32) int32 {
	if i == -1 || j == -1 {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if x[i] == y[j] {
		memo[i][j] = 1 + lcs(x, y, i-1, j-1, memo)
		return memo[i][j]
	}

	memo[i][j] = max(
		lcs(x, y, i-1, j, memo),
		lcs(x, y, i, j-1, memo),
	)

	return memo[i][j]
}

func longestCommonSubsequence(x []int32, y []int32) int32 {
	memo := make([][]int32, len(x))

	for i := 0; i < len(x); i++ {
		memo[i] = make([]int32, len(y))
	}

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			memo[i][j] = -1
		}
	}

	return lcs(x, y, int32(len(x)-1), int32(len(y)-1), memo)
}
