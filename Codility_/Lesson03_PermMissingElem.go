package main

func PermMissingElem(A []int) int {
	var (
		n       = len(A)
		counter = make([]int, n+1)
	)

	for i := range A {
		counter[A[i]-1]++
	}

	for i := range counter {
		if counter[i] == 0 {
			return i + 1
		}
	}

	return 0
}

// https://app.codility.com/demo/results/trainingMTS4EQ-D6N/
