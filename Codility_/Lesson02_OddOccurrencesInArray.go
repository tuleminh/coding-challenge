package main

func OddOccurrencesInArray(A []int) int {
	m := make(map[int]bool)
	for _, v := range A {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}
	unpairedElem := -1
	for k, _ := range m {
		unpairedElem = k
		break
	}
	return unpairedElem
}
