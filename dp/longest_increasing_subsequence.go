package dp

func longestIncreasingSubsequence(arr []int32) int32 {
	var maxSequence int32 = 0
	memo := make([]int32, len(arr))

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				memo[i] = max(memo[i], memo[j]+1)
			}
		}

		if memo[i] > maxSequence {
			maxSequence = memo[i]
		}
	}

	return maxSequence + 1
}
