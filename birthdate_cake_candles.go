package main

/**
 * https://www.hackerrank.com/challenges/birthday-cake-candles/problem?isFullScreen=true
 */
func BirthdateCakeCandles(candles []int32) int32 {
	var counter int32 = 0
	var tallest int32 = 0

	for _, candle := range candles {
		if candle > tallest {
			counter = 1
			tallest = candle
		} else if candle == tallest {
			counter += 1
		}
	}

	return counter
}
