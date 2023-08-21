package dp

import (
	"math/big"
)

// https://www.hackerrank.com/challenges/fibonacci-modified/problem?isFullScreen=false

func FibonacciModified(t1 int8, t2 int8, n int8) big.Int {
	table := make(map[int8]*big.Int)
	table[1] = big.NewInt(int64(t1))
	table[2] = big.NewInt(int64(t2))

	var i int8
	for i = 3; i <= n; i++ {
		num1 := table[i-2]
		num2 := table[i-1]

		table[i] = big.NewInt(0).Add(num1, big.NewInt(0).Exp(num2, big.NewInt(2), nil))
	}

	return *table[n]
}
