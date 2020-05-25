package main

func CyclicRotation(A []int, K int) []int {
	ALen := len(A)

	if ALen == K {
		return A
	}

	shiftIdxFn := func(i, K, N int) int {
		return (i + K) % N
	}

	shiftedA := make([]int, ALen)
	for i := range A {
		idx := shiftIdxFn(i, K, ALen)
		shiftedA[idx] = A[i]
	}

	return shiftedA
}

// https://app.codility.com/demo/results/trainingNKBBNR-DMF/
