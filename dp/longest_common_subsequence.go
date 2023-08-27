package dp

func longestCommonSubsequence(x []int32, y []int32) []int32 {
	var result, temp []int32
	memo := make([][]int32, len(x)+1)

	for i := 0; i <= len(x); i++ {
		memo[i] = make([]int32, len(y)+1)
	}

	for i := 0; i <= len(x); i++ {
		for j := 0; j <= len(y); j++ {
			if i == 0 || j == 0 {
				memo[i][j] = 0
				continue
			}

			if x[i-1] == y[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}

	i := len(x)
	j := len(y)
	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			temp = append(temp, x[i-1])
			i--
			j--
		} else if memo[i-1][j] > memo[i][j-1] {
			i--
		} else {
			j--
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}

	return result
}
