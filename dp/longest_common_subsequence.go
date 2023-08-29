package dp

func longestCommonSubsequence(arr1 []int32, arr2 []int32) []int32 {
	var result, temp []int32
	memo := make([][]int32, len(arr1)+1)

	for i := 0; i <= len(arr1); i++ {
		memo[i] = make([]int32, len(arr2)+1)
	}

	for i := 0; i <= len(arr1); i++ {
		for j := 0; j <= len(arr2); j++ {
			if i == 0 || j == 0 {
				memo[i][j] = 0
				continue
			}

			if arr1[i-1] == arr2[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}

	i := len(arr1)
	j := len(arr2)
	for i > 0 && j > 0 {
		if arr1[i-1] == arr2[j-1] {
			temp = append(temp, arr1[i-1])
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
