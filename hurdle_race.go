package main

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

func HurdleRace(k int32, height []int32) int32 {
	max := int32(0)

	for _, value := range height {
		if max < value {
			max = value
		}
	}

	if k >= max {
		return 0
	}

	return max - k
}
