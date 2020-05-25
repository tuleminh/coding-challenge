package main

import (
	"math"
)

func BinaryGap(N int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	bitWidthFn := func(num int) uint {
		return uint(math.Log2(float64(num))) + 1
	}

	var (
		bitWidth        = bitWidthFn(N)
		count, maxCount = 0, 0
	)

	for mask := uint(1) << (bitWidth - 1); mask > 0; mask >>= 1 {
		if uint(N)&mask == 0 {
			count++
		} else {
			maxCount = maxFn(count, maxCount)
			count = 0
		}
	}

	return maxCount
}

// https://app.codility.com/demo/results/trainingUFERS7-X92/
